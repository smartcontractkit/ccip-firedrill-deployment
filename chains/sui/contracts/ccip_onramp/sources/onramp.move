module ccip_onramp::onramp;

use std::ascii;
use std::bcs;
use std::string::{Self, String};
use std::type_name;
use sui::address;
use sui::derived_object;
use sui::event;
use sui::object;
use sui::transfer;

const ENothingToSend: u64 = 1;
const EMessageAlreadySent: u64 = 2;

public fun type_and_version(): String {
    string::utf8(b"OnRamp 1.6.0")
}

public struct ONRAMP has drop {}

public struct StaticConfig has drop, store {
    chain_selector: u64
}

public struct DynamicConfig has store, drop, copy {
    fee_aggregator: address,
    allowlist_admin: address
}

public struct OnRampState has key, store {
    id: UID,
    package_ids: vector<address>,
    s_send_last: u64,
}

public struct OnRampObject has key {
    id: UID,
}

public struct OnRampStatePointer has key, store {
    id: UID,
    on_ramp_object_id: address,
}

public struct RampMessageHeader has copy, drop, store {
    message_id: vector<u8>,
    source_chain_selector: u64,
    dest_chain_selector: u64,
    sequence_number: u64,
    nonce: u64,
}

public struct Sui2AnyRampMessage has copy, drop, store {
    header: RampMessageHeader,
    sender: address,
    data: vector<u8>,
    receiver: vector<u8>,
    extra_args: vector<u8>,
    fee_token: address,
    fee_token_amount: u64,
    fee_value_juels: u256,
    token_amounts: vector<Sui2AnyTokenTransfer>,
}

public struct Sui2AnyTokenTransfer has copy, drop, store {
    source_pool_address: address,
    // the token address on the destination chain
    dest_token_address: vector<u8>,
    extra_data: vector<u8>, // random bytes provided by token pool, e.g. encoded decimals
    amount: u64,
    dest_exec_data: vector<u8>, // destination gas amount
}

public struct CCIPMessageSent has copy, drop {
    dest_chain_selector: u64,
    sequence_number: u64,
    message: Sui2AnyRampMessage,
}

public struct DestChainConfigSet has copy, drop {
    dest_chain_selector: u64,
    sequence_number: u64,
    allowlist_enabled: bool,
    router: address,
}

public struct AllowlistSendersAdded has copy, drop {
    dest_chain_selector: u64,
    senders: vector<address>,
}

public struct AllowlistSendersRemoved has copy, drop {
    dest_chain_selector: u64,
    senders: vector<address>,
}

fun init(otw: ONRAMP, ctx: &mut TxContext) {
    let mut on_ramp_object = OnRampObject { id: object::new(ctx) };

    let pointer = OnRampStatePointer {
        id: object::new(ctx),
        on_ramp_object_id: object::id_address(&on_ramp_object),
    };

    let tn = type_name::with_original_ids<ONRAMP>();
    let package_bytes = ascii::into_bytes(tn.address_string());
    let package_id = address::from_ascii_bytes(&package_bytes);

    let state = OnRampState {
        id: derived_object::claim(&mut on_ramp_object.id, b"OnRampState"),
        package_ids: vector[package_id],
        s_send_last: 0,
    };

    transfer::share_object(state);
    transfer::share_object(on_ramp_object);

    transfer::transfer(pointer, package_id);
}

public fun emit_dest_chain_config_set(router: address) {
    let sui_selector = 9762610643973837292;
    event::emit(DestChainConfigSet {
        dest_chain_selector: sui_selector,
        sequence_number: 0,
        router: @ccip,
        allowlist_enabled: false,
    });
}

public fun get_static_config(_: &OnRampState): StaticConfig {
    StaticConfig { chain_selector: 9762610643973837292 }
}

public fun get_dynamic_config(_: &OnRampState): DynamicConfig {
    DynamicConfig { fee_aggregator: @ccip, allowlist_admin: @ccip }
}

public fun emit_allowlist_senders_added(dest_chain_selector: u64) {
    event::emit(AllowlistSendersAdded { dest_chain_selector, senders: vector[] });
}

public fun emit_allowlist_senders_removed(dest_chain_selector: u64) {
    event::emit(AllowlistSendersRemoved { dest_chain_selector, senders: vector[] });
}

public fun emit_ccip_message_sent(
    index: u64,
    fee_token: address,
    receiver: address,
    ctx: &TxContext,
) {
    let mut message_id = vector[];
    message_id.append(bcs::to_bytes(&ctx.sender()));
    message_id.append(bcs::to_bytes(&index));

    let message = Sui2AnyRampMessage {
        header: RampMessageHeader {
            message_id,
            source_chain_selector: 9762610643973837292,
            dest_chain_selector: 9762610643973837292,
            sequence_number: index,
            nonce: 1,
        },
        sender: ctx.sender(),
        data: b"123",
        receiver: bcs::to_bytes(&receiver),
        extra_args: b"123",
        fee_token,
        fee_token_amount: 0,
        fee_value_juels: 0,
        token_amounts: vector[],
    };

    event::emit(CCIPMessageSent {
        dest_chain_selector: 9762610643973837292,
        sequence_number: index,
        message,
    });
}

public fun get_ccip_package_id(): address {
    @ccip
}

public fun get_dest_chain_config(
    state: &OnRampState,
    dest_chain_selector: u64,
): (u64, bool, address) {
    (1, false, @ccip)
}

public fun drill_onramp_initialize(router: address) {
    emit_dest_chain_config_set(router);
}

public fun drill_allowlist_senders_added_removed() {
    emit_allowlist_senders_added(9762610643973837292);
    emit_allowlist_senders_removed(9762610643973837292);
}

public fun drill_pending_commit_pending_queue_tx_spike(
    state: &mut OnRampState,
    from: u8,
    to: u8,
    fee_token: address,
    ctx: &TxContext,
) {
    assert!(from <= to, ENothingToSend);
    assert!((from as u64) > state.s_send_last, EMessageAlreadySent);

    let mut i = from;
    while (i <= to) {
        emit_ccip_message_sent((i as u64), fee_token, ctx.sender(), ctx);
        i = i + 1;
    };

    state.s_send_last = (to as u64);
}

public fun drill_pending_commit_pending_queue_tx_spike_1(
    state: &mut OnRampState,
    from: u8,
    to: u8,
    fee_token: address,
    receiver: address,
    ctx: &TxContext,
) {
    assert!(from <= to, ENothingToSend);
    assert!((from as u64) > state.s_send_last, EMessageAlreadySent);

    let mut i = from;
    while (i <= to) {
        emit_ccip_message_sent((i as u64), fee_token, receiver, ctx);
        i = i + 1;
    };

    state.s_send_last = (to as u64);
}

public fun get_send_last(state: &OnRampState): u64 {
    state.s_send_last
}
