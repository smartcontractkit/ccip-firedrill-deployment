module ccip_offramp::offramp;

use ccip::state_object;
use ccip_offramp::ocr3_base;
use std::ascii;
use std::bcs;
use std::string::{Self, String};
use std::type_name;
use sui::address;
use sui::clock;
use sui::derived_object;
use sui::event;
use sui::hash;
use sui::package::{Self, UpgradeCap};
use sui::table::{Self, Table};
use sui::vec_map::{Self, VecMap};

const EXECUTION_STATE_UNTOUCHED: u8 = 0;
const EXECUTION_STATE_IN_PROGRESS: u8 = 1;
const EXECUTION_STATE_SUCCESS: u8 = 2;
const EXECUTION_STATE_FAILURE: u8 = 3;

const ENothingToSend: u64 = 1;

public struct OffRampObject has key {
    id: UID,
}

public struct OffRampStatePointer has key, store {
    id: UID,
    off_ramp_object_id: address,
}

public struct StaticConfigSet has copy, drop {
    chain_selector: u64,
}

public struct StaticConfig has copy, drop, store {
    chain_selector: u64,
    rmn_remote: address,
    token_admin_registry: address,
    nonce_manager: address,
}

public struct DynamicConfig has copy, drop, store {
    fee_quoter: address,
    permissionless_execution_threshold_seconds: u32, // The delay before manual exec is enabled
}

public struct SourceChainConfig has copy, drop, store {
    router: address,
    is_enabled: bool,
    min_seq_nr: u64,
    is_rmn_verification_disabled: bool,
    on_ramp: vector<u8>,
}

public struct DynamicConfigSet has copy, drop {
    dynamic_config: DynamicConfig,
}

public struct SourceChainConfigSet has copy, drop {
    source_chain_selector: u64,
    source_chain_config: SourceChainConfig,
}

public struct PriceUpdates has copy, drop, store {
    token_price_updates: vector<TokenPriceUpdate>,
    gas_price_updates: vector<GasPriceUpdate>,
}

public struct TokenPriceUpdate has copy, drop, store {
    source_token: address,
    usd_per_token: u256,
}

public struct GasPriceUpdate has copy, drop, store {
    dest_chain_selector: u64,
    usd_per_unit_gas: u256,
}

public struct MerkleRoot has copy, drop, store {
    source_chain_selector: u64,
    on_ramp_address: vector<u8>,
    min_seq_nr: u64,
    max_seq_nr: u64,
    merkle_root: vector<u8>,
}

public struct SkippedAlreadyExecuted has copy, drop {
    source_chain_selector: u64,
    sequence_number: u64,
}

public struct ExecutionStateChanged has copy, drop {
    source_chain_selector: u64,
    sequence_number: u64,
    message_id: vector<u8>,
    message_hash: vector<u8>,
    state: u8,
}

public struct CommitReportAccepted has copy, drop {
    blessed_merkle_roots: vector<MerkleRoot>,
    unblessed_merkle_roots: vector<MerkleRoot>,
    price_updates: PriceUpdates,
}

public struct SkippedReportExecution has copy, drop {
    source_chain_selector: u64,
}

public struct OffRampState has key, store {
    id: UID,
    package_ids: vector<address>,
}

public struct OFFRAMP has drop {}

fun init(otw: OFFRAMP, ctx: &mut TxContext) {
    let mut off_ramp_object = OffRampObject { id: object::new(ctx) };

    let pointer = OffRampStatePointer {
        id: object::new(ctx),
        off_ramp_object_id: object::id_address(&off_ramp_object),
    };

    let tn = type_name::with_original_ids<OFFRAMP>();
    let package_bytes = ascii::into_bytes(tn.address_string());
    let package_id = address::from_ascii_bytes(&package_bytes);

    let state = OffRampState {
        id: derived_object::claim(&mut off_ramp_object.id, b"OffRampState"),
        package_ids: vector[package_id],
    };

    transfer::share_object(state);
    transfer::share_object(off_ramp_object);

    transfer::transfer(pointer, package_id);
}

public fun emit_commit_report_accepted(
    min_seq_nr: u64,
    max_seq_nr: u64,
    onramp_address: address,
    source_chain_selector: u64,
) {
    let mut merkle_root = vector[];
    vector::append(&mut merkle_root, bcs::to_bytes(&onramp_address));
    vector::append(&mut merkle_root, bcs::to_bytes(&min_seq_nr));
    vector::append(&mut merkle_root, bcs::to_bytes(&max_seq_nr));

    let merkle_root = MerkleRoot {
        source_chain_selector,
        on_ramp_address: bcs::to_bytes(&onramp_address),
        min_seq_nr,
        max_seq_nr,
        merkle_root,
    };

    event::emit(CommitReportAccepted {
        blessed_merkle_roots: vector[],
        unblessed_merkle_roots: vector[merkle_root],
        price_updates: PriceUpdates { token_price_updates: vector[], gas_price_updates: vector[] },
    });
}

public fun emit_skipped_report_execution(source_chain_selector: u64) {
    event::emit(SkippedReportExecution { source_chain_selector });
}

public fun emit_skipped_already_executed(source_chain_selector: u64, sequence_number: u64) {
    event::emit(SkippedAlreadyExecuted { source_chain_selector, sequence_number });
}

public fun emit_execution_state_changed(source_chain_selector: u64, index: u64, ctx: &TxContext) {
    let mut message_id = vector[];
    message_id.append(bcs::to_bytes(&ctx.sender()));
    message_id.append(bcs::to_bytes(&index));

    let mut message_hash = vector[];
    message_hash.append(bcs::to_bytes(&ctx.sender()));
    message_hash.append(bcs::to_bytes(&index));

    event::emit(ExecutionStateChanged {
        source_chain_selector,
        sequence_number: index,
        message_id,
        message_hash,
        state: EXECUTION_STATE_SUCCESS,
    });
}

public fun emit_execution_state_changed_untouched(
    source_chain_selector: u64,
    index: u64,
    ctx: &TxContext,
) {
    let mut message_id = vector[];
    message_id.append(bcs::to_bytes(&ctx.sender()));
    message_id.append(bcs::to_bytes(&index));

    let mut message_hash = vector[];
    message_hash.append(bcs::to_bytes(&ctx.sender()));
    message_hash.append(bcs::to_bytes(&index));

    event::emit(ExecutionStateChanged {
        source_chain_selector,
        sequence_number: index,
        message_id,
        message_hash,
        state: EXECUTION_STATE_UNTOUCHED,
    });
}

public fun emit_execution_state_changed_in_progress(
    source_chain_selector: u64,
    index: u64,
    ctx: &TxContext,
) {
    let mut message_id = vector[];
    message_id.append(bcs::to_bytes(&ctx.sender()));
    message_id.append(bcs::to_bytes(&index));

    let mut message_hash = vector[];
    message_hash.append(bcs::to_bytes(&ctx.sender()));
    message_hash.append(bcs::to_bytes(&index));

    event::emit(ExecutionStateChanged {
        source_chain_selector,
        sequence_number: index,
        message_id,
        message_hash,
        state: EXECUTION_STATE_IN_PROGRESS,
    });
}

public fun emit_execution_state_changed_failure(
    source_chain_selector: u64,
    index: u64,
    ctx: &TxContext,
) {
    let mut message_id = vector[];
    message_id.append(bcs::to_bytes(&ctx.sender()));
    message_id.append(bcs::to_bytes(&index));

    let mut message_hash = vector[];
    message_hash.append(bcs::to_bytes(&ctx.sender()));
    message_hash.append(bcs::to_bytes(&index));

    event::emit(ExecutionStateChanged {
        source_chain_selector,
        sequence_number: index,
        message_id,
        message_hash,
        state: EXECUTION_STATE_FAILURE,
    });
}

public fun emit_static_config_set() {
    let sui_selector = 17529533435026248318;
    event::emit(StaticConfigSet { chain_selector: sui_selector });
}

public fun emit_dynamic_config_set(ref: &state_object::CCIPObjectRef, state: &OffRampState) {
    let dynamic_config = get_dynamic_config(ref, state);
    event::emit(DynamicConfigSet { dynamic_config });
}

public fun emit_source_chain_config_set() {
    let source_chain_config = SourceChainConfig {
        router: @ccip,
        is_enabled: true,
        min_seq_nr: 0,
        is_rmn_verification_disabled: false,
        on_ramp: bcs::to_bytes(&@onramp),
    };
    let sui_selector = 17529533435026248318;
    event::emit(SourceChainConfigSet { source_chain_selector: sui_selector, source_chain_config });
}

public fun get_static_config(
    ref: &state_object::CCIPObjectRef,
    state: &OffRampState,
): StaticConfig {
    let sui_selector = 17529533435026248318;
    StaticConfig {
        chain_selector: sui_selector,
        rmn_remote: @ccip,
        token_admin_registry: @ccip,
        nonce_manager: @ccip,
    }
}

public fun get_dynamic_config(
    ref: &state_object::CCIPObjectRef,
    state: &OffRampState,
): DynamicConfig {
    DynamicConfig {
        fee_quoter: @ccip,
        permissionless_execution_threshold_seconds: 10 as u32,
    }
}

public fun get_static_config_fields(
    ref: &state_object::CCIPObjectRef,
    cfg: StaticConfig,
): (u64, address, address, address) {
    (cfg.chain_selector, cfg.rmn_remote, cfg.token_admin_registry, cfg.nonce_manager)
}

public fun get_source_chain_config(
    ref: &state_object::CCIPObjectRef,
    state: &OffRampState,
    source_chain_selector: u64,
): SourceChainConfig {
    SourceChainConfig {
        router: @ccip,
        is_enabled: true,
        min_seq_nr: 0,
        is_rmn_verification_disabled: false,
        on_ramp: bcs::to_bytes(&@onramp),
    }
}

public fun get_dynamic_config_fields(_: &state_object::CCIPObjectRef, cfg: DynamicConfig): (address, u32) {
    (cfg.fee_quoter, cfg.permissionless_execution_threshold_seconds)
}

public fun type_and_version(): String {
    string::utf8(b"OffRamp 1.6.0")
}

public fun emit_ocr3_base_config_set() {
    ocr3_base::emit_config_set();
}

public fun get_ccip_package_id(): address {
    @ccip
}

public fun get_all_source_chain_configs(
    ref: &state_object::CCIPObjectRef,
    state: &OffRampState,
): (vector<u64>, vector<SourceChainConfig>) {
    let sui_selector = 17529533435026248318;
    let source_chain_selectors = vector[sui_selector];
    let source_chain_config = SourceChainConfig {
        router: @ccip,
        is_enabled: true,
        min_seq_nr: 2,
        is_rmn_verification_disabled: false,
        on_ramp: bcs::to_bytes(&@onramp),
    };

    (source_chain_selectors, vector[source_chain_config])
}

public fun add_package_id(state: &mut OffRampState, package_id: address) {
    state.package_ids.push_back(package_id);
}

public fun remove_package_id(state: &mut OffRampState, package_id: address) {
    let (found, idx) = state.package_ids.index_of(&package_id);
    assert!(found, 1);
    state.package_ids.swap_remove(idx);
}

public fun prepare_register() {
    emit_source_chain_config_set();
    emit_ocr3_base_config_set();
}

public fun drill_pending_execution(state: &OffRampState, from: u8, to: u8) {
    assert!(from <= to, ENothingToSend);

    emit_commit_report_accepted((from as u64), (to as u64), @onramp, 17529533435026248318);
}

public fun drill_offramp_initialize(ref: &state_object::CCIPObjectRef, state: &OffRampState) {
    emit_static_config_set();
    emit_dynamic_config_set(ref, state);
    emit_source_chain_config_set();
}

public fun drill_offramp_execute(ctx: &TxContext) {
    emit_skipped_already_executed(17529533435026248318, 1);
    emit_skipped_report_execution(17529533435026248318);
    emit_execution_state_changed(17529533435026248318, 1, ctx);
}

public fun get_source_chain_config_fields(
    _: &state_object::CCIPObjectRef,
    source_chain_config: SourceChainConfig,
): (address, bool, u64, bool, vector<u8>) {
    (
        source_chain_config.router,
        source_chain_config.is_enabled,
        source_chain_config.min_seq_nr,
        source_chain_config.is_rmn_verification_disabled,
        source_chain_config.on_ramp,
    )
}
