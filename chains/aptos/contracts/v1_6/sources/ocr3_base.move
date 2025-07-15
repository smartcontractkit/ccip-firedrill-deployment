module firedrill::ocr3_base {
    use std::account;
    use std::event::{Self, EventHandle};

    friend firedrill::offramp;

    const MAX_NUM_ORACLES: u64 = 256;

    const OCR_PLUGIN_TYPE_COMMIT: u8 = 0;
    const OCR_PLUGIN_TYPE_EXECUTION: u8 = 1;

    struct OCR3BaseState has store {
        config_set_events: EventHandle<ConfigSet>,
        transmitted_events: EventHandle<Transmitted>
    }

    #[event]
    struct ConfigSet has store, drop {
        ocr_plugin_type: u8,
        config_digest: vector<u8>,
        signers: vector<vector<u8>>,
        transmitters: vector<address>,
        big_f: u8
    }

    #[event]
    struct Transmitted has store, drop {
        ocr_plugin_type: u8,
        config_digest: vector<u8>,
        sequence_number: u64
    }

    public fun new(event_account: &signer): OCR3BaseState {
        OCR3BaseState {
            config_set_events: account::new_event_handle(event_account),
            transmitted_events: account::new_event_handle(event_account)
        }
    }

    public(friend) fun emit_config_set(state: &mut OCR3BaseState) {
        event::emit_event(
            &mut state.config_set_events,
            ConfigSet {
                ocr_plugin_type: 0,
                config_digest: vector[],
                signers: vector[],
                transmitters: vector[],
                big_f: 0
            }
        );
        event::emit(
            ConfigSet {
                ocr_plugin_type: 0,
                config_digest: vector[],
                signers: vector[],
                transmitters: vector[],
                big_f: 0
            }
        );
    }
}
