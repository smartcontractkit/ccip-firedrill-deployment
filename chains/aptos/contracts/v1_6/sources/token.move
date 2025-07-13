module firedrill::token {
    use std::option;
    use std::string::{Self, String};
    use std::fungible_asset::{Self, Metadata, MintRef, BurnRef, TransferRef};
    use std::object::{Self, Object, ExtendRef};
    use std::primary_fungible_store;
    use std::signer;

    const TOKEN_NAME: vector<u8> = b"Firedrill Token";
    const TOKEN_SYMBOL: vector<u8> = b"LINK";
    const TOKEN_DECIMALS: u8 = 0;
    const TOKEN_ICON: vector<u8> = b"";
    const TOKEN_PROJECT: vector<u8> = b"";

    #[resource_group_member(group = aptos_framework::object::ObjectGroup)]
    struct TokenState has key {
        token_address: address,
        extend_ref: ExtendRef,
        mint_ref: MintRef,
        burn_ref: BurnRef,
        transfer_ref: TransferRef
    }

    const E_NOT_CODE_OBJECT: u64 = 1;

    fun init_module(deployer: &signer) {
        assert!(signer::address_of(deployer) == @firedrill, E_NOT_CODE_OBJECT);

        let constructor_ref = &object::create_named_object(deployer, TOKEN_SYMBOL);

        primary_fungible_store::create_primary_store_enabled_fungible_asset(
            constructor_ref,
            option::none(),
            string::utf8(TOKEN_NAME),
            string::utf8(TOKEN_SYMBOL),
            TOKEN_DECIMALS,
            string::utf8(TOKEN_ICON),
            string::utf8(TOKEN_PROJECT)
        );

        let metadata_object_signer = &object::generate_signer(constructor_ref);

        move_to(
            metadata_object_signer,
            TokenState {
                token_address: object::address_from_constructor_ref(constructor_ref),
                extend_ref: object::generate_extend_ref(constructor_ref),
                mint_ref: fungible_asset::generate_mint_ref(constructor_ref),
                burn_ref: fungible_asset::generate_burn_ref(constructor_ref),
                transfer_ref: fungible_asset::generate_transfer_ref(constructor_ref)
            }
        );
    }

    #[view]
    public fun balance_of(_account: address): u64 {
        0
    }

    #[view]
    public fun decimals(): u8 {
        TOKEN_DECIMALS
    }

    #[view]
    public fun symbol(): String {
        string::utf8(TOKEN_SYMBOL)
    }

    #[view]
    public fun name(): String {
        string::utf8(TOKEN_NAME)
    }

    #[view]
    public fun type_and_version(): String {
        string::utf8(b"FiredrillToken 1.0.0")
    }

    #[view]
    public fun token_address(): address {
        object::create_object_address(&@firedrill, TOKEN_SYMBOL)
    }

    #[view]
    public fun metadata(): Object<Metadata> {
        object::address_to_object(token_address())
    }
}
