use crate::program::FiredrillOfframp;
use crate::{CcipOfframpError, Config};
use anchor_lang::prelude::*;
use bytemuck::{Pod, Zeroable};
use shared::seed;

/// Static space allocated to any account: must always be added to space calculations.
pub const ANCHOR_DISCRIMINATOR: usize = 8;

#[derive(Accounts)]
pub struct InitializeConfig<'info> {
    #[account(
        init,
        seeds = [seed::CONFIG],
        bump,
        payer = authority,
        space = ANCHOR_DISCRIMINATOR + Config::INIT_SPACE,
    )]
    pub config: AccountLoader<'info, Config>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,

    #[account(constraint = program.programdata_address()? == Some(program_data.key()))]
    pub program: Program<'info, FiredrillOfframp>,

    // Initialization only allowed by program upgrade authority
    #[account(constraint = program_data.upgrade_authority_address == Some(authority.key()) @ CcipOfframpError::Unauthorized)]
    pub program_data: Account<'info, ProgramData>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct PriceUpdates {
    pub token_price_updates: Vec<TokenPriceUpdate>,
    pub gas_price_updates: Vec<GasPriceUpdate>,
}

#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct TokenPriceUpdate {
    pub source_token: Pubkey, // It is the mint, but called "token" for EVM compatibility.
    pub usd_per_token: [u8; 28], // EVM uses u224, 1e18 USD per 1e18 of the smallest token denomination.
}

/// Gas price for a given chain in USD; its value may contain tightly packed fields.
#[derive(Clone, AnchorSerialize, AnchorDeserialize)]
pub struct GasPriceUpdate {
    pub dest_chain_selector: u64,
    pub usd_per_unit_gas: [u8; 28], // EVM uses u224, 1e18 USD per smallest unit (e.g. wei) of destination chain gas
}

/// Struct to hold a merkle root and an interval for a source chain
#[derive(Default, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct MerkleRoot {
    pub source_chain_selector: u64, // Remote source chain selector that the Merkle Root is scoped to
    pub on_ramp_address: Vec<u8>,   // Generic onramp address, to support arbitrary sources
    pub min_seq_nr: u64,            // Minimum sequence number, inclusive
    pub max_seq_nr: u64,            // Maximum sequence number, inclusive
    pub merkle_root: [u8; 32],      // Merkle root covering the interval & source chain messages
}

/// It's not possible to store enums in zero_copy accounts, so we wrap the discriminant
/// in a struct to store in config.
#[derive(
    Copy, Clone, AnchorSerialize, AnchorDeserialize, PartialEq, Eq, InitSpace, Pod, Zeroable,
)]
#[repr(C)]
pub struct ConfigOcrPluginType {
    discriminant: u8,
}

impl From<OcrPluginType> for ConfigOcrPluginType {
    fn from(value: OcrPluginType) -> Self {
        Self {
            discriminant: value as u8,
        }
    }
}

impl TryFrom<ConfigOcrPluginType> for OcrPluginType {
    type Error = CcipOfframpError;

    fn try_from(value: ConfigOcrPluginType) -> std::result::Result<Self, Self::Error> {
        match value.discriminant {
            0 => Ok(Self::Commit),
            1 => Ok(Self::Execution),
            _ => Err(CcipOfframpError::InvalidPluginType),
        }
    }
}

#[derive(Copy, Clone, AnchorSerialize, AnchorDeserialize, PartialEq, Eq)]
pub enum OcrPluginType {
    Commit,
    Execution,
}
