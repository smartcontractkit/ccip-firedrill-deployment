use anchor_lang::prelude::*;
use anchor_spl::token_interface::Mint;
use shared::seed;

pub mod state;
use crate::state::*;

declare_id!("8y39U4xMcKnPP1NyWdYCEFTuCgZ18STYgeCxsjnSVcHm");

#[program]
pub mod firedrill_feequoter {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, chain_selector: u64, compound: Pubkey) -> Result<()> {
        let fee_quoter = &mut ctx.accounts.fee_quoter;
        fee_quoter.chain_selector = chain_selector;
        fee_quoter.owner = ctx.accounts.authority.key();
        fee_quoter.compound = compound;
        fee_quoter.token = ctx.accounts.token.key();

        ctx.accounts.config.set_inner(Config {
            version: 1,
            owner: ctx.accounts.authority.key(),
            proposed_owner: Pubkey::default(),
            max_fee_juels_per_msg: 100,
            link_token_mint: ctx.accounts.token.key(),
            link_token_local_decimals: ctx.accounts.token.decimals,
            onramp: compound,
            default_code_version: CodeVersion::V1,
        });

        ctx.accounts.dest_chain.set_inner(DestChain {
            version: 1,
            chain_selector,
            state: DestChainState {
                usd_per_unit_gas: TimestampedPackedU224 {
                    value: [0u8; 28],
                    timestamp: 0,
                },
            },
            config: DestChainConfig {
                is_enabled: true,
                lane_code_version: CodeVersion::V1,
                max_number_of_tokens_per_msg: 10,
                max_data_bytes: 1024,
                max_per_msg_gas_limit: 5_000_000,
                dest_gas_overhead: 20_000,
                dest_gas_per_payload_byte_base: 16,
                dest_gas_per_payload_byte_high: 64,
                dest_gas_per_payload_byte_threshold: 512,
                dest_data_availability_overhead_gas: 0,
                dest_gas_per_data_availability_byte: 0,
                dest_data_availability_multiplier_bps: 0,
                default_token_fee_usdcents: 100,
                default_token_dest_gas_overhead: 50_000,
                default_tx_gas_limit: 1_000_000,
                gas_multiplier_wei_per_eth: 1,
                network_fee_usdcents: 50,
                gas_price_staleness_threshold: 0,
                enforce_out_of_order: false,
                chain_family_selector: [0u8; 4],
            },
        });

        Ok(())
    }

    pub fn emit_usd_per_token_updated(ctx: Context<EmitUsdPerToken>) -> Result<()> {
        let fee_quoter = &ctx.accounts.fee_quoter;

        let token = fee_quoter.token;
        let mut value = [0u8; 28];
        let val = 1u128.to_le_bytes(); // 16 bytes
        value[..16].copy_from_slice(&val);
        let timestamp = Clock::get()?.unix_timestamp;

        emit!(UsdPerTokenUpdated {
            token,
            value,
            timestamp,
        });

        Ok(())
    }
}

#[account]
pub struct FiredrillFeeQuoter {
    pub chain_selector: u64,
    pub owner: Pubkey,
    pub compound: Pubkey,
    pub token: Pubkey,
}

#[derive(Accounts)]
#[instruction(chain_selector: u64)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::FEE_QUOTER], bump, payer = authority, space = 8 + 8 + 32 + 32 + 32)]
    pub fee_quoter: Account<'info, FiredrillFeeQuoter>,

    #[account(
        init,
        seeds = [seed::CONFIG],
        bump,
        payer = authority,
        space = 8 + Config::INIT_SPACE,
    )]
    pub config: Account<'info, Config>,

    #[account(
        init,
        seeds = [seed::DEST_CHAIN, chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = 8 + DestChain::INIT_SPACE,
    )]
    pub dest_chain: Account<'info, DestChain>,

    pub token: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Config {
    pub version: u8,

    pub owner: Pubkey,
    pub proposed_owner: Pubkey,

    // Static config fields
    //
    // Maximum fee that can be charged for a message.
    pub max_fee_juels_per_msg: u128,
    pub link_token_mint: Pubkey,
    // local LINK mint may not use 18 decimals as EVM, so we store local decimals
    // to calculate the transfer fees in juels later.
    pub link_token_local_decimals: u8,
    // TODO The following field is unused until the day we integrate with feeds to fetch fresh values
    // pub token_price_staleness_threshold: u32,
    pub onramp: Pubkey,

    pub default_code_version: CodeVersion,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct DestChain {
    // Config for SVM2Any
    pub version: u8,
    pub chain_selector: u64,     // Chain selector used for the seed
    pub state: DestChainState,   // values that are updated automatically
    pub config: DestChainConfig, // values configured by an admin
}

#[derive(Accounts)]
pub struct EmitUsdPerToken<'info> {
    #[account(mut, has_one = owner)]
    pub fee_quoter: Account<'info, FiredrillFeeQuoter>,
    pub owner: Signer<'info>,
}

#[event]
pub struct UsdPerTokenUpdated {
    pub token: Pubkey,
    pub value: [u8; 28],
    pub timestamp: i64,
}