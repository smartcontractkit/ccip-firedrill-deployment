module ccip::state_object;

use std::ascii;
use std::string::{Self, String};
use std::type_name;
use sui::address;
use sui::derived_object;
use sui::event;
use sui::object;
use sui::transfer;
use sui::vec_map::{Self, VecMap};

public fun type_and_version(): String {
    string::utf8(b"StateObject 1.6.0")
}

public struct CCIPObject has key {
    id: UID,
}

public struct CCIPObjectRef has key, store {
    id: UID,
    package_ids: vector<address>,
}

public struct CCIPObjectRefPointer has key, store {
    id: UID,
    ccip_object_id: address,
}

public struct STATE_OBJECT has drop {}

fun init(otw: STATE_OBJECT, ctx: &mut TxContext) {
    let mut ccip_object = CCIPObject { id: object::new(ctx) };

    let mut ref = CCIPObjectRef {
        id: derived_object::claim(&mut ccip_object.id, b"CCIPObjectRef"),
        package_ids: vector[],
    };

    let pointer = CCIPObjectRefPointer {
        id: object::new(ctx),
        ccip_object_id: object::id_address(&ccip_object),
    };

    let tn = type_name::with_original_ids<STATE_OBJECT>();
    let package_bytes = ascii::into_bytes(tn.address_string());
    let package_id = address::from_ascii_bytes(&package_bytes);
    ref.package_ids.push_back(package_id);

    transfer::share_object(ref);
    transfer::share_object(ccip_object);

    transfer::transfer(pointer, package_id);
}

public fun add_package_id(ref: &mut CCIPObjectRef, package_id: address) {
    ref.package_ids.push_back(package_id);
}

public fun remove_package_id(ref: &mut CCIPObjectRef, package_id: address) {
    let (found, idx) = ref.package_ids.index_of(&package_id);
    assert!(found, 1);
    ref.package_ids.swap_remove(idx);
}
