module ccip_onramp::onramp;

use std::bcs;
use sui::event;

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

public fun emit_dest_chain_config_set(router: address) {
    let sui_selector = 1234567890;
    event::emit(DestChainConfigSet {
        dest_chain_selector: sui_selector,
        sequence_number: 0,
        router,
        allowlist_enabled: false,
    });
}

public fun emit_allowlist_senders_added(dest_chain_selector: u64) {
    event::emit(AllowlistSendersAdded { dest_chain_selector, senders: vector[] });
}

public fun emit_allowlist_senders_removed(dest_chain_selector: u64) {
    event::emit(AllowlistSendersRemoved { dest_chain_selector, senders: vector[] });
}

public fun emit_ccip_message_sent(
    index: u64,
    source_chain_selector: u64,
    dest_chain_selector: u64,
    fee_token: address,
    ctx: &TxContext,
) {
    let mut message_id = vector[];
    message_id.append(bcs::to_bytes(&ctx.sender()));
    message_id.append(bcs::to_bytes(&index));

    let message = Sui2AnyRampMessage {
        header: RampMessageHeader {
            message_id,
            source_chain_selector,
            dest_chain_selector,
            sequence_number: index,
            nonce: 1,
        },
        sender: ctx.sender(),
        data: b"123",
        receiver: bcs::to_bytes(&ctx.sender()),
        extra_args: b"123",
        fee_token,
        fee_token_amount: 0,
        fee_value_juels: 0,
        token_amounts: vector[],
    };

    event::emit(CCIPMessageSent {
        dest_chain_selector,
        sequence_number: index,
        message,
    });
}
