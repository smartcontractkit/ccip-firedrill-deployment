module ccip::fee_quoter;

use sui::event;

public struct UsdPerTokenUpdated has copy, drop {
    token: address,
    usd_per_token: u256,
    timestamp: u64,
}

public struct UsdPerUnitGasUpdated has copy, drop {
    dest_chain_selector: u64,
    usd_per_unit_gas: u256,
    timestamp: u64,
}

public fun emit_usd_per_token_updated(token: address, usd_per_token: u256) {
    event::emit(UsdPerTokenUpdated {
        token,
        usd_per_token,
        timestamp: 1234567890,
    })
}

public fun emit_usd_per_unit_gas_updated(dest_chain_selector: u64) {
    event::emit(UsdPerUnitGasUpdated {
        dest_chain_selector,
        usd_per_unit_gas: 191919191919,
        timestamp: 1234567890,
    })
}
