module firedrill::onramp {
    use std::signer;
    use std::string::{Self, String};
    use std::event::{Self, EventHandle};
    use std::object::{Self, Object};
    use std::bcs;
    use std::account;

    use firedrill::fee_quoter::{Self};
    use firedrill::ownable::{Self, OwnableState};
    use firedrill::state_object;

    friend firedrill::entrypoint;

    #[resource_group_member(group = aptos_framework::object::ObjectGroup)]
    struct OnRampState has key {
        ownable_state: OwnableState,

        /// Events
        config_set_events: EventHandle<ConfigSet>,
        dest_chain_config_set_events: EventHandle<DestChainConfigSet>,
        ccip_message_sent_events: EventHandle<CCIPMessageSent>,
        allowlist_senders_added_events: EventHandle<AllowlistSendersAdded>,
        allowlist_senders_removed_events: EventHandle<AllowlistSendersRemoved>
    }

    struct DestChainConfig has store, drop {
        sequence_number: u64,
        allowlist_enabled: bool,
        router: address,
        allowed_senders: vector<address>
    }

    struct RampMessageHeader has store, drop, copy {
        message_id: vector<u8>,
        source_chain_selector: u64,
        dest_chain_selector: u64,
        sequence_number: u64,
        nonce: u64
    }

    struct Aptos2AnyRampMessage has store, drop, copy {
        header: RampMessageHeader,
        sender: address,
        data: vector<u8>,
        receiver: vector<u8>,
        extra_args: vector<u8>,
        fee_token: address,
        fee_token_amount: u64,
        fee_value_juels: u256,
        token_amounts: vector<Aptos2AnyTokenTransfer>
    }

    struct Aptos2AnyTokenTransfer has store, drop, copy {
        source_pool_address: address,
        dest_token_address: vector<u8>,
        extra_data: vector<u8>,
        amount: u64,
        dest_exec_data: vector<u8>
    }

    struct StaticConfig has drop, store {
        chain_selector: u64
    }

    struct DynamicConfig has store, drop, copy {
        fee_aggregator: address,
        allowlist_admin: address
    }

    #[event]
    struct ConfigSet has store, drop {
        static_config: StaticConfig,
        dynamic_config: DynamicConfig
    }

    #[event]
    struct DestChainConfigSet has store, drop {
        dest_chain_selector: u64,
        sequence_number: u64,
        router: address,
        allowlist_enabled: bool
    }

    #[event]
    struct CCIPMessageSent has store, drop {
        dest_chain_selector: u64,
        sequence_number: u64,
        message: Aptos2AnyRampMessage
    }

    #[event]
    struct AllowlistSendersAdded has store, drop {
        dest_chain_selector: u64,
        senders: vector<address>
    }

    #[event]
    struct AllowlistSendersRemoved has store, drop {
        dest_chain_selector: u64,
        senders: vector<address>
    }

    #[event]
    struct FeeTokenWithdrawn has store, drop {
        fee_aggregator: address,
        fee_token: address,
        amount: u64
    }

    const E_NOT_CODE_OBJECT: u64 = 1;

    fun init_module(deployer: &signer) {
        assert!(signer::address_of(deployer) == @firedrill, E_NOT_CODE_OBJECT);

        let object_signer = &state_object::object_signer();
        let state = OnRampState {
            ownable_state: ownable::new(object_signer, @firedrill),
            config_set_events: account::new_event_handle(object_signer),
            dest_chain_config_set_events: account::new_event_handle(object_signer),
            ccip_message_sent_events: account::new_event_handle(object_signer),
            allowlist_senders_added_events: account::new_event_handle(object_signer),
            allowlist_senders_removed_events: account::new_event_handle(object_signer)
        };
        move_to(object_signer, state);
    }

    public(friend) fun emit_ccip_message_sent(
        sender: address, index: u64
    ) acquires OnRampState {
        let message_id = vector[];
        message_id.append(bcs::to_bytes(&sender));
        message_id.append(bcs::to_bytes(&index));

        let message = Aptos2AnyRampMessage {
            header: RampMessageHeader {
                message_id,
                source_chain_selector: fee_quoter::chain_selector(),
                dest_chain_selector: fee_quoter::chain_selector(),
                sequence_number: index,
                nonce: 1
            },
            sender,
            data: b"123",
            receiver: bcs::to_bytes(&sender),
            extra_args: b"123",
            fee_token: fee_quoter::token_address(),
            fee_token_amount: 0,
            fee_value_juels: 0,
            token_amounts: vector[]
        };
        event::emit_event(
            &mut borrow_mut().ccip_message_sent_events,
            CCIPMessageSent {
                dest_chain_selector: fee_quoter::chain_selector(),
                sequence_number: index,
                message
            }
        );
        event::emit(
            CCIPMessageSent {
                dest_chain_selector: fee_quoter::chain_selector(),
                sequence_number: index,
                message
            }
        );
    }

    public(friend) fun emit_dest_chain_config_set() acquires OnRampState {
        event::emit_event(
            &mut borrow_mut().dest_chain_config_set_events,
            DestChainConfigSet {
                dest_chain_selector: fee_quoter::chain_selector(),
                sequence_number: 0,
                router: @firedrill,
                allowlist_enabled: false
            }
        );
        event::emit(
            DestChainConfigSet {
                dest_chain_selector: fee_quoter::chain_selector(),
                sequence_number: 0,
                router: @firedrill,
                allowlist_enabled: false
            }
        );
    }

    public(friend) fun emit_allowlist_senders_added() acquires OnRampState {
        event::emit_event(
            &mut borrow_mut().allowlist_senders_added_events,
            AllowlistSendersAdded {
                dest_chain_selector: fee_quoter::chain_selector(),
                senders: vector[]
            }
        );
        event::emit(
            AllowlistSendersAdded {
                dest_chain_selector: fee_quoter::chain_selector(),
                senders: vector[]
            }
        );
    }

    public(friend) fun emit_allowlist_senders_removed() acquires OnRampState {
        event::emit_event(
            &mut borrow_mut().allowlist_senders_removed_events,
            AllowlistSendersRemoved {
                dest_chain_selector: fee_quoter::chain_selector(),
                senders: vector[]
            }
        );
        event::emit(
            AllowlistSendersRemoved {
                dest_chain_selector: fee_quoter::chain_selector(),
                senders: vector[]
            }
        );
    }

    #[view]
    public fun get_static_config(): StaticConfig {
        StaticConfig { chain_selector: fee_quoter::chain_selector() }
    }

    #[view]
    public fun get_dynamic_config(): DynamicConfig {
        DynamicConfig { fee_aggregator: @firedrill, allowlist_admin: @0x0 }
    }

    #[view]
    public fun get_dest_chain_config(_dest_chain_selector: u64): (u64, bool, address) {
        (0, false, @firedrill)
    }

    #[view]
    public fun type_and_version(): String {
        string::utf8(b"OnRamp 1.6.0")
    }

    #[view]
    public fun onramp_state_object(): Object<OnRampState> {
        object::address_to_object<OnRampState>(get_state_address())
    }

    #[view]
    public fun get_state_address(): address {
        state_object::object_address()
    }

    // ======================== Ownable ========================

    public entry fun transfer_ownership(caller: &signer, to: address) acquires OnRampState {
        ownable::transfer_ownership(caller, &mut borrow_mut().ownable_state, to)
    }

    public entry fun accept_ownership(caller: &signer) acquires OnRampState {
        ownable::accept_ownership(caller, &mut borrow_mut().ownable_state)
    }

    public entry fun execute_ownership_transfer(
        caller: &signer, to: address
    ) acquires OnRampState {
        ownable::execute_ownership_transfer(caller, &mut borrow_mut().ownable_state, to)
    }

    inline fun borrow_mut(): &mut OnRampState {
        borrow_global_mut<OnRampState>(get_state_address())
    }

    inline fun borrow(): &OnRampState {
        borrow_global<OnRampState>(get_state_address())
    }
}
