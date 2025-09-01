module firedrill::entrypoint {
    use std::signer;
    use std::string::{Self, String};
    use std::object::{Self, Object};
    use firedrill::fee_quoter::{Self};
    use firedrill::onramp::{Self};
    use firedrill::offramp::{Self};
    use firedrill::state_object;

    #[resource_group_member(group = aptos_framework::object::ObjectGroup)]
    struct EntrypointState has key {
        s_send_last: u64
    }

    const E_NOT_OWNER: u64 = 1;
    const E_NOT_CODE_OBJECT: u64 = 2;
    const E_NOTHING_TO_SEND: u64 = 3;
    const E_MESSAGE_ALREADY_SENT: u64 = 4;
    const E_MESSAGE_NOT_SENT: u64 = 5;

    fun init_module(deployer: &signer) {
        assert!(signer::address_of(deployer) == @firedrill, E_NOT_CODE_OBJECT);

        let object_signer = &state_object::object_signer();
        move_to(object_signer, EntrypointState { s_send_last: 0 });
    }

    public entry fun prepare_register(admin: &signer) {
        // fee_quoter::assert_only_owner(admin);

        offramp::emit_source_chain_config_set(); // register OffRamp
        offramp::emit_ocr3_base_config_set(); // register OCR3Base
    }

    public entry fun drill_pending_commit_pending_queue_tx_spike(
        admin: &signer, from: u8, to: u8
    ) acquires EntrypointState {
        // fee_quoter::assert_only_owner(admin);

        let state = borrow_mut();
        assert!(from <= to, E_NOTHING_TO_SEND);
        assert!((from as u64) > state.s_send_last, E_MESSAGE_ALREADY_SENT);

        for (i in from..(to + 1)) {
            onramp::emit_ccip_message_sent(signer::address_of(admin), (i as u64));
        };

        state.s_send_last = (to as u64);
    }

    public entry fun drill_pending_execution(
        admin: &signer, from: u8, to: u8
    ) acquires EntrypointState {
        // fee_quoter::assert_only_owner(admin);

        assert!(from <= to, E_NOTHING_TO_SEND);
        assert!((to as u64) <= borrow().s_send_last, E_MESSAGE_NOT_SENT);

        offramp::emit_commit_report_accepted((from as u64), (to as u64));
    }

    public entry fun drill_price_registries(admin: &signer) {
        // fee_quoter::assert_only_owner(admin);

        fee_quoter::emit_usd_per_token_updated();
        fee_quoter::emit_usd_per_unit_gas_updated();
    }

    public entry fun drill_offramp_initialize(admin: &signer) {
        // fee_quoter::assert_only_owner(admin);

        offramp::emit_static_config_set();
        offramp::emit_dynamic_config_set();
        offramp::emit_source_chain_config_set();
    }

    public entry fun drill_onramp_initialize(admin: &signer) {
        // fee_quoter::assert_only_owner(admin);

        onramp::emit_dest_chain_config_set();
    }

    public entry fun drill_allowlist_senders_added_removed(admin: &signer) {
        // fee_quoter::assert_only_owner(admin);

        onramp::emit_allowlist_senders_added();
        onramp::emit_allowlist_senders_removed();
    }

    public entry fun drill_offramp_execute(admin: &signer) {
        // fee_quoter::assert_only_owner(admin);

        offramp::emit_skipped_already_executed();
        offramp::emit_skipped_report_execution();
        offramp::emit_execution_state_changed(signer::address_of(admin), 0);
    }

    #[view]
    public fun get_send_last(): u64 acquires EntrypointState {
        borrow().s_send_last
    }

    #[view]
    public fun can_execute_range(from: u8, to: u8): bool acquires EntrypointState {
        from <= to && (to as u64) <= borrow().s_send_last
    }

    #[view]
    public fun can_send_range(from: u8, to: u8): bool acquires EntrypointState {
        from <= to && (from as u64) > borrow().s_send_last
    }

    #[view]
    public fun type_and_version(): String {
        string::utf8(b"RouterFeeQuoter 1.6.0")
    }

    #[view]
    public fun entrypoint_state_object(): Object<EntrypointState> {
        object::address_to_object<EntrypointState>(get_state_address())
    }

    #[view]
    public fun get_state_address(): address {
        state_object::object_address()
    }

    inline fun borrow_mut(): &mut EntrypointState {
        borrow_global_mut<EntrypointState>(get_state_address())
    }

    inline fun borrow(): &EntrypointState {
        borrow_global<EntrypointState>(get_state_address())
    }

    inline fun assert_owner(owner: &signer) {
        assert!(
            object::owns(entrypoint_state_object(), signer::address_of(owner)),
            E_NOT_OWNER
        );
    }
}
