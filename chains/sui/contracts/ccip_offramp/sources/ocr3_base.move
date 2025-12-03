module ccip_offramp::ocr3_base;

use sui::event;

public struct ConfigSet has copy, drop {
    ocr_plugin_type: u8,
    config_digest: vector<u8>,
    signers: vector<vector<u8>>,
    transmitters: vector<address>,
    big_f: u8,
}

public fun emit_config_set() {
    event::emit(ConfigSet {
        ocr_plugin_type: 1,
        config_digest: vector[],
        signers: vector[],
        transmitters: vector[],
        big_f: 2,
    });
}
