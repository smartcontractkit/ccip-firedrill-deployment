module firedrill::offramp {
    use std::signer;
    use std::string::{Self, String};
    use std::vector;
    use std::event::{Self, EventHandle};
    use std::object::{Self, Object};
    use std::account;
    use std::bcs;

    use firedrill::compound::{Self};
    use firedrill::ownable::{Self, OwnableState};
    use firedrill::state_object;

    friend firedrill::entrypoint;

    const EXECUTION_STATE_UNTOUCHED: u8 = 0;
    const EXECUTION_STATE_SUCCESS: u8 = 2;

    #[resource_group_member(group = aptos_framework::object::ObjectGroup)]
    struct OffRampState has key {
        ownable_state: OwnableState,

        /// Events
        static_config_set_events: EventHandle<StaticConfigSet>,
        dynamic_config_set_events: EventHandle<DynamicConfigSet>,
        source_chain_config_set_events: EventHandle<SourceChainConfigSet>,
        skipped_already_executed_events: EventHandle<SkippedAlreadyExecuted>,
        execution_state_changed_events: EventHandle<ExecutionStateChanged>,
        commit_report_accepted_events: EventHandle<CommitReportAccepted>,
        skipped_report_execution_events: EventHandle<SkippedReportExecution>
    }

    struct SourceChainConfig has store, drop, copy {
        router: address,
        is_enabled: bool,
        min_seq_nr: u64,
        is_rmn_verification_disabled: bool,
        on_ramp: vector<u8>
    }

    struct RampMessageHeader has drop {
        message_id: vector<u8>,
        source_chain_selector: u64,
        dest_chain_selector: u64,
        sequence_number: u64,
        nonce: u64
    }

    struct Any2AptosRampMessage has drop {
        header: RampMessageHeader,
        sender: vector<u8>,
        data: vector<u8>,
        receiver: address,
        gas_limit: u256,
        token_amounts: vector<Any2AptosTokenTransfer>
    }

    struct Any2AptosTokenTransfer has drop {
        source_pool_address: vector<u8>,
        dest_token_address: address,
        dest_gas_amount: u32,
        extra_data: vector<u8>,
        amount: u256 // This is the amount to transfer, as set on the source chain.
    }

    struct ExecutionReport has drop {
        source_chain_selector: u64,
        message: Any2AptosRampMessage,
        offchain_token_data: vector<vector<u8>>,
        proofs: vector<vector<u8>> // Proofs used to construct the merkle root
    }

    // Matches the EVM struct
    struct CommitReport has store, drop, copy {
        price_updates: PriceUpdates, // Price updates for the fee_quoter
        blessed_merkle_roots: vector<MerkleRoot>, // Merkle roots that have been blessed by RMN
        unblessed_merkle_roots: vector<MerkleRoot>, // Merkle roots that don't require RMN blessing
        rmn_signatures: vector<vector<u8>> // The signatures for the blessed merkle roots
    }

    struct PriceUpdates has store, drop, copy {
        token_price_updates: vector<TokenPriceUpdate>,
        gas_price_updates: vector<GasPriceUpdate>
    }

    struct TokenPriceUpdate has store, drop, copy {
        source_token: address, // This is the local token
        usd_per_token: u256
    }

    struct GasPriceUpdate has store, drop, copy {
        dest_chain_selector: u64,
        usd_per_unit_gas: u256
    }

    struct MerkleRoot has store, drop, copy {
        source_chain_selector: u64,
        on_ramp_address: vector<u8>,
        min_seq_nr: u64,
        max_seq_nr: u64,
        merkle_root: vector<u8>
    }

    struct StaticConfig has store, drop, copy {
        chain_selector: u64,
        rmn_remote: address,
        token_admin_registry: address,
        nonce_manager: address
    }

    struct DynamicConfig has store, drop, copy {
        // On EVM, the feeQuoter is a dynamic address but due to the Aptos implementation using a static
        // upgradable FeeQuoter, this value is actually static. For compatibility reasons, we keep it as a dynamic config.
        fee_quoter: address,
        permissionless_execution_threshold_seconds: u32 // The delay before manual exec is enabled
    }

    #[event]
    struct StaticConfigSet has store, drop {
        static_config: StaticConfig
    }

    #[event]
    struct DynamicConfigSet has store, drop {
        dynamic_config: DynamicConfig
    }

    #[event]
    struct SourceChainConfigSet has store, drop {
        source_chain_selector: u64,
        source_chain_config: SourceChainConfig
    }

    #[event]
    struct SkippedAlreadyExecuted has store, drop {
        source_chain_selector: u64,
        sequence_number: u64
    }

    #[event]
    struct AlreadyAttempted has store, drop {
        source_chain_selector: u64,
        sequence_number: u64
    }

    #[event]
    struct ExecutionStateChanged has store, drop {
        source_chain_selector: u64,
        sequence_number: u64,
        message_id: vector<u8>,
        message_hash: vector<u8>,
        state: u8
    }

    #[event]
    struct CommitReportAccepted has store, drop {
        blessed_merkle_roots: vector<MerkleRoot>,
        unblessed_merkle_roots: vector<MerkleRoot>,
        price_updates: PriceUpdates
    }

    #[event]
    struct SkippedReportExecution has store, drop {
        source_chain_selector: u64
    }

    const E_NOT_CODE_OBJECT: u64 = 1;

    fun init_module(deployer: &signer) {
        assert!(signer::address_of(deployer) == @firedrill, E_NOT_CODE_OBJECT);

        let object_signer = &state_object::object_signer();
        move_to(
            object_signer,
            OffRampState {
                ownable_state: ownable::new(object_signer, @firedrill),
                static_config_set_events: account::new_event_handle(object_signer),
                dynamic_config_set_events: account::new_event_handle(object_signer),
                source_chain_config_set_events: account::new_event_handle(object_signer),
                skipped_already_executed_events: account::new_event_handle(object_signer),
                execution_state_changed_events: account::new_event_handle(object_signer),
                commit_report_accepted_events: account::new_event_handle(object_signer),
                skipped_report_execution_events: account::new_event_handle(object_signer)
            }
        );
    }

    public(friend) fun emit_commit_report_accepted(
        min_seq_nr: u64, max_seq_nr: u64
    ) acquires OffRampState {
        let chain_selector = compound::chain_selector();
        let on_ramp_address = bcs::to_bytes(&compound::onramp_address());

        let merkle_root = vector[];
        vector::append(&mut merkle_root, on_ramp_address);
        vector::append(&mut merkle_root, bcs::to_bytes(&min_seq_nr));
        vector::append(&mut merkle_root, bcs::to_bytes(&max_seq_nr));

        let blessed_root = MerkleRoot {
            source_chain_selector: chain_selector,
            on_ramp_address,
            min_seq_nr,
            max_seq_nr,
            merkle_root
        };

        let price_updates = PriceUpdates {
            token_price_updates: vector[],
            gas_price_updates: vector[]
        };

        event::emit_event(
            &mut borrow_mut().commit_report_accepted_events,
            CommitReportAccepted {
                blessed_merkle_roots: vector[blessed_root],
                unblessed_merkle_roots: vector[],
                price_updates
            }
        );
        event::emit(
            CommitReportAccepted {
                blessed_merkle_roots: vector[blessed_root],
                unblessed_merkle_roots: vector[],
                price_updates
            }
        );
    }

    public(friend) fun emit_static_config_set() acquires OffRampState {
        let static_config = get_static_config();
        event::emit_event(
            &mut borrow_mut().static_config_set_events,
            StaticConfigSet { static_config }
        );
        event::emit(StaticConfigSet { static_config });
    }

    public(friend) fun emit_dynamic_config_set() acquires OffRampState {
        let dynamic_config = get_dynamic_config();
        event::emit_event(
            &mut borrow_mut().dynamic_config_set_events,
            DynamicConfigSet { dynamic_config }
        );
        event::emit(DynamicConfigSet { dynamic_config });
    }

    public(friend) fun emit_source_chain_config_set() acquires OffRampState {
        let source_chain_config = SourceChainConfig {
            router: @firedrill,
            is_enabled: true,
            min_seq_nr: 0,
            is_rmn_verification_disabled: false,
            on_ramp: bcs::to_bytes(&compound::onramp_address())
        };
        event::emit_event(
            &mut borrow_mut().source_chain_config_set_events,
            SourceChainConfigSet {
                source_chain_selector: compound::chain_selector(),
                source_chain_config
            }
        );
        event::emit(
            SourceChainConfigSet {
                source_chain_selector: compound::chain_selector(),
                source_chain_config
            }
        );
    }

    public(friend) fun emit_skipped_report_execution() acquires OffRampState {
        event::emit_event(
            &mut borrow_mut().skipped_report_execution_events,
            SkippedReportExecution { source_chain_selector: compound::chain_selector() }
        );
        event::emit(
            SkippedReportExecution { source_chain_selector: compound::chain_selector() }
        );
    }

    public(friend) fun emit_skipped_already_executed() acquires OffRampState {
        let source_chain_selector = compound::chain_selector();
        let sequence_number = 0;

        event::emit_event(
            &mut borrow_mut().skipped_already_executed_events,
            SkippedAlreadyExecuted { source_chain_selector, sequence_number }
        );
        event::emit(SkippedAlreadyExecuted { source_chain_selector, sequence_number });
    }

    public(friend) fun emit_execution_state_changed(
        sender: address, index: u64
    ) acquires OffRampState {
        let message_id = vector[];
        message_id.append(bcs::to_bytes(&sender));
        message_id.append(bcs::to_bytes(&index));

        let message_hash = vector[];
        message_hash.append(bcs::to_bytes(&sender));
        message_hash.append(bcs::to_bytes(&index));

        event::emit_event(
            &mut borrow_mut().execution_state_changed_events,
            ExecutionStateChanged {
                source_chain_selector: compound::chain_selector(),
                sequence_number: index,
                message_id,
                message_hash,
                state: EXECUTION_STATE_SUCCESS
            }
        );
        event::emit(
            ExecutionStateChanged {
                source_chain_selector: compound::chain_selector(),
                sequence_number: index,
                message_id,
                message_hash,
                state: EXECUTION_STATE_SUCCESS
            }
        );
    }

    #[view]
    public fun get_static_config(): StaticConfig {
        StaticConfig {
            chain_selector: compound::chain_selector(),
            rmn_remote: @firedrill,
            token_admin_registry: @firedrill,
            nonce_manager: @firedrill
        }
    }

    #[view]
    public fun get_dynamic_config(): DynamicConfig {
        DynamicConfig {
            fee_quoter: @firedrill,
            permissionless_execution_threshold_seconds: 10
        }
    }

    #[view]
    public fun get_source_chain_config(_source_chain_selector: u64): SourceChainConfig {
        SourceChainConfig {
            router: @firedrill,
            is_enabled: true,
            min_seq_nr: 0,
            is_rmn_verification_disabled: false,
            on_ramp: bcs::to_bytes(&compound::onramp_address())
        }
    }

    #[view]
    public fun type_and_version(): String {
        string::utf8(b"OffRamp 1.6.0")
    }

    #[view]
    public fun get_state_object(): Object<OffRampState> {
        object::address_to_object<OffRampState>(get_state_address())
    }

    #[view]
    public fun get_state_address(): address {
        state_object::object_address()
    }

    // ======================== Ownable ========================

    public entry fun transfer_ownership(caller: &signer, to: address) acquires OffRampState {
        ownable::transfer_ownership(caller, &mut borrow_mut().ownable_state, to)
    }

    public entry fun accept_ownership(caller: &signer) acquires OffRampState {
        ownable::accept_ownership(caller, &mut borrow_mut().ownable_state)
    }

    public entry fun execute_ownership_transfer(
        caller: &signer, to: address
    ) acquires OffRampState {
        ownable::execute_ownership_transfer(caller, &mut borrow_mut().ownable_state, to)
    }

    inline fun borrow_mut(): &mut OffRampState {
        borrow_global_mut<OffRampState>(get_state_address())
    }

    inline fun borrow(): &OffRampState {
        borrow_global<OffRampState>(get_state_address())
    }
}
