module firedrill::compound {
    use std::signer;
    use std::event::{Self, EventHandle};
    use std::account::{Self};
    use std::object::{Self, Object};
    use std::timestamp;
    use firedrill::token;
    use firedrill::ownable::{Self, OwnableState};
    use firedrill::state_object;

    friend firedrill::entrypoint;

    const COMPOUND_SEED: vector<u8> = b"FiredrillCompound";
    const CHAIN_SELECTOR: u64 = 10;

    struct StaticConfig has copy, drop, store {
        max_fee_juels_per_msg: u128,
        link_token: address,
        token_price_staleness_threshold: u32
    }

    struct OffRamp has copy, drop, store {
        source_chain_selector: u64,
        off_ramp: address
    }

    #[resource_group_member(group = aptos_framework::object::ObjectGroup)]
    struct FeeQuoterState has key {
        ownable_state: OwnableState,
        chain_selector: u64,
        token_address: address,
        onramp_address: address,
        offramp_address: address,
        receiver_address: address,
        usd_per_token_updated_events: EventHandle<UsdPerTokenUpdated>,
        usd_per_unit_gas_updated_events: EventHandle<UsdPerUnitGasUpdated>
    }

    #[event]
    struct UsdPerTokenUpdated has drop, store {
        token: address,
        value: u256,
        timestamp: u64
    }

    #[event]
    struct UsdPerUnitGasUpdated has store, drop {
        dest_chain_selector: u64,
        usd_per_unit_gas: u256,
        timestamp: u64
    }

    fun init_module(_deployer: &signer) {
        let object_signer = &state_object::object_signer();
        let ownable_state = ownable::new(object_signer, @firedrill);

        move_to(
            object_signer,
            FeeQuoterState {
                ownable_state,
                chain_selector: CHAIN_SELECTOR,
                token_address: token::token_address(),
                onramp_address: @firedrill,
                offramp_address: @firedrill,
                receiver_address: @firedrill,
                usd_per_token_updated_events: account::new_event_handle(object_signer),
                usd_per_unit_gas_updated_events: account::new_event_handle(object_signer)
            }
        );
    }

    /// Update receiver address
    public entry fun set_receiver(caller: &signer, receiver: address) acquires FeeQuoterState {
        assert_only_owner(caller);
        borrow_fee_quoter_state_mut().receiver_address = receiver;
    }

    public(friend) fun emit_usd_per_token_updated() acquires FeeQuoterState {
        let state = borrow_fee_quoter_state_mut();

        event::emit_event(
            &mut state.usd_per_token_updated_events,
            UsdPerTokenUpdated {
                token: state.token_address,
                value: 1,
                timestamp: timestamp::now_seconds()
            }
        );
        event::emit(
            UsdPerTokenUpdated {
                token: state.token_address,
                value: 1,
                timestamp: timestamp::now_seconds()
            }
        )
    }

    public(friend) fun emit_usd_per_unit_gas_updated() acquires FeeQuoterState {
        let state = borrow_fee_quoter_state_mut();
        event::emit_event(
            &mut state.usd_per_unit_gas_updated_events,
            UsdPerUnitGasUpdated {
                dest_chain_selector: state.chain_selector,
                usd_per_unit_gas: 1,
                timestamp: timestamp::now_seconds()
            }
        );
        event::emit(
            UsdPerUnitGasUpdated {
                dest_chain_selector: state.chain_selector,
                usd_per_unit_gas: 1,
                timestamp: timestamp::now_seconds()
            }
        )
    }

    // ========== Router ==========

    #[view]
    public fun get_on_ramp(_: u64): address acquires FeeQuoterState {
        borrow_fee_quoter_state().onramp_address
    }

    #[view]
    public fun get_off_ramps(): vector<OffRamp> acquires FeeQuoterState {
        let state = borrow_fee_quoter_state();
        vector[
            OffRamp {
                source_chain_selector: state.chain_selector,
                off_ramp: state.offramp_address
            }
        ]
    }

    // ========== FeeQuoter ==========
    #[view]
    public fun get_static_config(): StaticConfig {
        StaticConfig {
            max_fee_juels_per_msg: 1,
            link_token: token::token_address(),
            token_price_staleness_threshold: 0
        }
    }

    // ========== ARMProxy ==========
    #[view]
    public fun get_arm(): address {
        @firedrill
    }

    #[view]
    public fun chain_selector(): u64 acquires FeeQuoterState {
        borrow_fee_quoter_state().chain_selector
    }

    #[view]
    public fun token_address(): address acquires FeeQuoterState {
        borrow_fee_quoter_state().token_address
    }

    #[view]
    public fun onramp_address(): address acquires FeeQuoterState {
        borrow_fee_quoter_state().onramp_address
    }

    #[view]
    public fun offramp_address(): address acquires FeeQuoterState {
        borrow_fee_quoter_state().offramp_address
    }

    #[view]
    public fun receiver_address(): address acquires FeeQuoterState {
        borrow_fee_quoter_state().receiver_address
    }

    #[view]
    public fun get_fee_quoter_state_object(): Object<FeeQuoterState> {
        object::address_to_object<FeeQuoterState>(get_state_address())
    }

    #[view]
    public fun get_state_address(): address {
        state_object::object_address()
    }

    // ======================== Ownable ========================

    public entry fun transfer_ownership(caller: &signer, to: address) acquires FeeQuoterState {
        ownable::transfer_ownership(
            caller, &mut borrow_fee_quoter_state_mut().ownable_state, to
        )
    }

    public entry fun accept_ownership(caller: &signer) acquires FeeQuoterState {
        ownable::accept_ownership(
            caller, &mut borrow_fee_quoter_state_mut().ownable_state
        )
    }

    public entry fun execute_ownership_transfer(
        caller: &signer, to: address
    ) acquires FeeQuoterState {
        ownable::execute_ownership_transfer(
            caller, &mut borrow_fee_quoter_state_mut().ownable_state, to
        )
    }

    inline fun borrow_fee_quoter_state_mut(): &mut FeeQuoterState {
        borrow_global_mut<FeeQuoterState>(get_state_address())
    }

    inline fun borrow_fee_quoter_state(): &FeeQuoterState {
        borrow_global<FeeQuoterState>(get_state_address())
    }

    public fun assert_only_owner(owner: &signer) acquires FeeQuoterState {
        ownable::assert_only_owner(
            signer::address_of(owner), &borrow_fee_quoter_state().ownable_state
        );
    }
}
