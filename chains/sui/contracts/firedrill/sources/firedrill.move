module firedrill::firedrill;

use ccip::fee_quoter;
use ccip_offramp::offramp;
use ccip_onramp::onramp;
use std::string;

const ENothingToSend: u64 = 1;
const EMessageAlreadySent: u64 = 2;
const EMessageNotSent: u64 = 3;

public struct FiredrillState has key {
    id: UID,
    s_send_last: u64,
}

fun init(ctx: &mut TxContext) {
    let state = FiredrillState {
        id: object::new(ctx),
        s_send_last: 0,
    };

    transfer::share_object(state);
}

/// ================================ OFFRAMP FUNCTIONS ================================

public fun drill_pending_execution(
    state: &FiredrillState,
    from: u64,
    to: u64,
    onramp_address: address,
    source_chain_selector: u64,
) {
    assert!(from <= to, ENothingToSend);
    assert!((to as u64) <= state.s_send_last, EMessageNotSent);
    offramp::emit_commit_report_accepted(from, to, onramp_address, source_chain_selector);
}

public fun drill_offramp_execute(
    source_chain_selector: u64,
    sequence_number: u64,
    ctx: &TxContext,
) {
    offramp::emit_skipped_already_executed(source_chain_selector, sequence_number);
    offramp::emit_skipped_report_execution(source_chain_selector);
    offramp::emit_execution_state_changed(source_chain_selector, sequence_number, ctx);
}

public fun drill_offramp_initialize() {
    offramp::emit_static_config_set();
    offramp::emit_dynamic_config_set();
    offramp::emit_source_chain_config_set();
}

public fun prepare_register() {
    offramp::emit_source_chain_config_set(); // register OffRamp
    offramp::emit_ocr3_base_config_set(); // register OCR3Base
}

/// ================================ ONRAMP FUNCTIONS ================================

public fun drill_onramp_initialize(router: address) {
    onramp::emit_dest_chain_config_set(router);
}

public fun drill_allowlist_senders_added_removed(dest_chain_selector: u64) {
    onramp::emit_allowlist_senders_added(dest_chain_selector);
    onramp::emit_allowlist_senders_removed(dest_chain_selector);
}

public fun drill_pending_commit_pending_queue_tx_spike(
    state: &mut FiredrillState,
    from: u8,
    to: u8,
    source_chain_selector: u64,
    dest_chain_selector: u64,
    fee_token: address,
    ctx: &TxContext,
) {
    assert!(from <= to, ENothingToSend);
    assert!((from as u64) > state.s_send_last, EMessageAlreadySent);

    let mut i = from;
    while (i <= to) {
        onramp::emit_ccip_message_sent(
            (i as u64),
            source_chain_selector,
            dest_chain_selector,
            fee_token,
            ctx,
        );
        i = i + 1;
    };

    state.s_send_last = (to as u64);
}

/// ================================ FEE QUOTER FUNCTIONS ================================

public fun drill_price_registries(dest_chain_selector: u64, token: address, usd_per_token: u256) {
    fee_quoter::emit_usd_per_token_updated(token, usd_per_token);
    fee_quoter::emit_usd_per_unit_gas_updated(dest_chain_selector);
}

/// ================================ VIEW FUNCTIONS ================================

public fun get_send_last(state: &FiredrillState): u64 {
    state.s_send_last
}

public fun can_execute_range(state: &FiredrillState, from: u8, to: u8): bool {
    from <= to && (to as u64) <= state.s_send_last
}

public fun can_send_range(state: &FiredrillState, from: u8, to: u8): bool {
    from <= to && (from as u64) > state.s_send_last
}

public fun type_and_version(): string::String {
    string::utf8(b"FiredrillEntrypoint 1.0.0")
}
