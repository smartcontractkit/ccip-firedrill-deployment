module ccip::fee_quoter;

use sui::clock;
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

public fun emit_usd_per_token_updated(clock: &clock::Clock, token: address, usd_per_token: u256) {
    event::emit(UsdPerTokenUpdated {
        token,
        usd_per_token,
        timestamp: clock.timestamp_ms(),
    })
}

public fun emit_usd_per_unit_gas_updated(clock: &clock::Clock, usd_per_unit_gas: u256) {
    event::emit(UsdPerUnitGasUpdated {
        dest_chain_selector: 9762610643973837292,
        usd_per_unit_gas,
        timestamp: clock.timestamp_ms(),
    })
}
