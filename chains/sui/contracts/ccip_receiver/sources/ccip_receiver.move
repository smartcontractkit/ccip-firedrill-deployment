module ccip_receiver::ccip_receiver;

const EInvalidTimestamp: u64 = 1;

public struct CCIP_RECEIVER has drop {}

public struct CCIPReceiverState has key, store {
    id: UID,
}

fun init(_otw: CCIP_RECEIVER, ctx: &mut TxContext) {
    let ccip_receiver_state = CCIPReceiverState {
        id: object::new(ctx),
    };

    transfer::share_object(ccip_receiver_state);
}

public fun ccip_receive() {
    assert!(1 > 2, EInvalidTimestamp);
}
