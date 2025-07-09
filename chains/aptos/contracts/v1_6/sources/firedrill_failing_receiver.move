module firedrill::firedrill_failing_receiver {
    use std::account;
    use std::event;
    use std::object::Object;
    use std::option::{Option};
    use std::signer;
    use std::string::{Self, String};

    use ccip::receiver_registry;

    #[event]
    struct ReceivedMessage has store, drop {
        data: vector<u8>
    }

    struct CCIPReceiverState has key {
        received_message_events: event::EventHandle<ReceivedMessage>
    }

    const E_NOT_CODE_OBJECT: u64 = 1;
    const E_FAILING_RECEIVER: u64 = 2;

    #[view]
    public fun type_and_version(): String {
        string::utf8(b"FiredrillFailingReceiver 1.6.0")
    }

    fun init_module(publisher: &signer) {
        assert!(signer::address_of(publisher) == @firedrill, E_NOT_CODE_OBJECT);

        account::create_account_if_does_not_exist(@firedrill);

        let handle = account::new_event_handle(publisher);

        move_to(publisher, CCIPReceiverState { received_message_events: handle });

        receiver_registry::register_receiver(
            publisher, b"firedrill_failing_receiver", FailingReceiverProof {}
        );
    }

    struct FailingReceiverProof has drop {}

    public fun ccip_receive<T: key>(_metadata: Object<T>): Option<u128> {
        abort E_FAILING_RECEIVER
    }

    inline fun borrow_state_mut(): &mut CCIPReceiverState {
        borrow_global_mut<CCIPReceiverState>(@firedrill)
    }

    #[test_only]
    public fun new_received_message_event(data: vector<u8>): ReceivedMessage {
        ReceivedMessage { data }
    }

    #[test_only]
    public fun test_init_module(publisher: &signer) {
        init_module(publisher);
    }

    #[test_only]
    public fun new_failing_receiver_proof(): FailingReceiverProof {
        FailingReceiverProof {}
    }

    #[test_only]
    public fun get_received_message_events(): vector<ReceivedMessage> acquires CCIPReceiverState {
        let state = borrow_global<CCIPReceiverState>(@firedrill);
        event::emitted_events_by_handle<ReceivedMessage>(&state.received_message_events)
    }
}
