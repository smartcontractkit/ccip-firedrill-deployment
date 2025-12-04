module ccip_router::router;

use std::ascii;
use std::string::{Self, String};
use std::type_name;
use sui::address;
use sui::derived_object;
use sui::event;
use sui::object;
use sui::transfer;
use sui::vec_map::{Self, VecMap};

public struct ROUTER has drop {}

const EOnrampNotFound: u64 = 1;
const EParamsLengthMismatch: u64 = 2;
const EInvalidOnrampAddress: u64 = 3;

public struct RouterObject has key {
    id: UID,
}

public struct OnRampSet has copy, drop {
    dest_chain_selector: u64,
    on_ramp_package_id: address,
}

public struct RouterState has key {
    id: UID,
    on_ramp_package_ids: VecMap<u64, address>, // dest_chain_selector -> on_ramp_package_id
}

public struct RouterStatePointer has key, store {
    id: UID,
    router_object_id: address,
}

public fun type_and_version(): String {
    string::utf8(b"Router 1.6.0")
}

fun init(otw: ROUTER, ctx: &mut TxContext) {
    let mut router_object = RouterObject { id: object::new(ctx) };

    let router = RouterState {
        id: derived_object::claim(&mut router_object.id, b"RouterState"),
        on_ramp_package_ids: vec_map::empty(),
    };

    let router_state_pointer = RouterStatePointer {
        id: object::new(ctx),
        router_object_id: object::id_address(&router_object),
    };

    let tn = type_name::with_original_ids<ROUTER>();
    let package_bytes = ascii::into_bytes(tn.address_string());
    let package_id = address::from_ascii_bytes(&package_bytes);

    transfer::share_object(router);
    transfer::share_object(router_object);

    transfer::transfer(router_state_pointer, package_id);
}

public fun is_chain_supported(router: &RouterState, dest_chain_selector: u64): bool {
    router.on_ramp_package_ids.contains(&dest_chain_selector)
}

// Returns the on ramp package id for the given destination chain selector.
public fun get_on_ramp(router: &RouterState, dest_chain_selector: u64): address {
    assert!(router.on_ramp_package_ids.contains(&dest_chain_selector), EOnrampNotFound);

    *router.on_ramp_package_ids.get(&dest_chain_selector)
}

public fun get_dest_chains(router: &RouterState): vector<u64> {
    router.on_ramp_package_ids.keys()
}

public fun set_on_ramps(
    router: &mut RouterState,
    dest_chain_selectors: vector<u64>,
    on_ramp_package_ids: vector<address>,
) {
    assert!(dest_chain_selectors.length() == on_ramp_package_ids.length(), EParamsLengthMismatch);

    let mut i = 0;
    let selector_len = dest_chain_selectors.length();
    while (i < selector_len) {
        let dest_chain_selector = dest_chain_selectors[i];
        let on_ramp_package_id = on_ramp_package_ids[i];
        assert!(on_ramp_package_id != @0x0, EInvalidOnrampAddress);

        if (router.on_ramp_package_ids.contains(&dest_chain_selector)) {
            router.on_ramp_package_ids.remove(&dest_chain_selector);
        };
        router.on_ramp_package_ids.insert(dest_chain_selector, on_ramp_package_id);
        event::emit(OnRampSet { dest_chain_selector, on_ramp_package_id });
        i = i + 1;
    };
}
