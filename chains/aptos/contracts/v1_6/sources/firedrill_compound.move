module firedrill::firedrill_compound {
    use std::signer;
    use std::event::{Self, EventHandle};
    use std::account::{Self};
    use std::object::{Self, Object, ExtendRef};
    use std::timestamp;
    use firedrill::firedrill_token;
    use firedrill::ownable::{Self, OwnableState};

    friend firedrill::firedrill_entrypoint;

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
    struct CompoundState has key {
        extend_ref: ExtendRef,
        ownable_state: OwnableState,
        chain_selector: u64,
        token_address: address,
        onramp_address: address,
        offramp_address: address,
        receiver_address: address,
        usd_per_token_updated_events: EventHandle<UsdPerTokenUpdated>
    }

    #[event]
    struct UsdPerTokenUpdated has drop, store {
        token: address,
        value: u256,
        timestamp: u64
    }

    fun init_module(deployer: &signer) {
        let constructor_ref = object::create_named_object(deployer, COMPOUND_SEED);
        let extend_ref = object::generate_extend_ref(&constructor_ref);
        let object_signer = &object::generate_signer(&constructor_ref);

        account::create_account_if_does_not_exist(signer::address_of(object_signer));

        let ownable_state = ownable::new(object_signer, @firedrill);

        move_to(
            object_signer,
            CompoundState {
                extend_ref,
                ownable_state,
                chain_selector: CHAIN_SELECTOR,
                token_address: firedrill_token::token_address(),
                onramp_address: @firedrill,
                offramp_address: @firedrill,
                receiver_address: @firedrill,
                usd_per_token_updated_events: account::new_event_handle(object_signer)
            }
        );
    }

    /// Update receiver address
    public entry fun set_receiver(caller: &signer, receiver: address) acquires CompoundState {
        assert_only_owner(caller);
        borrow_mut().receiver_address = receiver;
    }

    public(friend) fun emit_usd_per_token_updated() acquires CompoundState {
        let state = borrow_mut();

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

    // ========== Router ==========

    #[view]
    public fun get_on_ramp(_: u64): address acquires CompoundState {
        borrow().onramp_address
    }

    #[view]
    public fun get_off_ramps(): vector<OffRamp> acquires CompoundState {
        vector[
            OffRamp {
                source_chain_selector: borrow().chain_selector,
                off_ramp: borrow().offramp_address
            }
        ]
    }

    // ========== FeeQuoter ==========
    #[view]
    public fun get_static_config(): StaticConfig {
        StaticConfig {
            max_fee_juels_per_msg: 1,
            link_token: firedrill_token::token_address(),
            token_price_staleness_threshold: 0
        }
    }

    // ========== ARMProxy ==========
    #[view]
    public fun get_arm(): address {
        @firedrill
    }

    #[view]
    public fun chain_selector(): u64 acquires CompoundState {
        borrow().chain_selector
    }

    #[view]
    public fun token_address(): address acquires CompoundState {
        borrow().token_address
    }

    #[view]
    public fun onramp_address(): address acquires CompoundState {
        borrow().onramp_address
    }

    #[view]
    public fun offramp_address(): address acquires CompoundState {
        borrow().offramp_address
    }

    #[view]
    public fun receiver_address(): address acquires CompoundState {
        borrow().receiver_address
    }

    #[view]
    public fun compound_state_object(): Object<CompoundState> {
        object::address_to_object<CompoundState>(compound_state_address())
    }

    #[view]
    public fun compound_state_address(): address {
        object::create_object_address(&@firedrill, COMPOUND_SEED)
    }

    // ======================== Ownable ========================

    public entry fun transfer_ownership(caller: &signer, to: address) acquires CompoundState {
        ownable::transfer_ownership(caller, &mut borrow_mut().ownable_state, to)
    }

    public entry fun accept_ownership(caller: &signer) acquires CompoundState {
        ownable::accept_ownership(caller, &mut borrow_mut().ownable_state)
    }

    public entry fun execute_ownership_transfer(
        caller: &signer, to: address
    ) acquires CompoundState {
        ownable::execute_ownership_transfer(caller, &mut borrow_mut().ownable_state, to)
    }

    inline fun borrow_mut(): &mut CompoundState {
        borrow_global_mut<CompoundState>(compound_state_address())
    }

    inline fun borrow(): &CompoundState {
        borrow_global<CompoundState>(compound_state_address())
    }

    public fun assert_only_owner(owner: &signer) acquires CompoundState {
        ownable::assert_only_owner(signer::address_of(owner), &borrow().ownable_state);
    }
}
