use anchor_lang::prelude::*;
use solana_program::keccak::hash;

use shared::ids::offramp::ID;
use shared::common::seed;

mod state;
use crate::state::*;

mod event;
use crate::event::*;

mod context;
use crate::context::*;

#[program]
pub mod firedrill_offramp {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, chain_selector: u64, token: Pubkey, on_ramp: Pubkey) -> Result<()> {
        let offramp = &mut ctx.accounts.offramp;
        offramp.owner = ctx.accounts.authority.key();
        offramp.chain_selector = chain_selector;
        offramp.token = token;
        offramp.on_ramp = on_ramp;
        Ok(())
    }

    /// Initializes the CCIP Offramp Config account.
    ///
    /// The initialization of the Offramp is responsibility of Admin, nothing more than calling these
    /// initialization methods should be done first.
    ///
    /// # Arguments
    ///
    /// * `ctx` - The context containing the accounts required for initialization of the config.
    /// * `svm_chain_selector` - The chain selector for SVM.
    /// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
    pub fn initialize_config(
        ctx: Context<InitializeConfig>,
        svm_chain_selector: u64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_init()?;
        require!(config.version == 0, CcipOfframpError::InvalidVersion); // assert uninitialized state - AccountLoader doesn't work with constraint
        config.version = 1;
        config.default_code_version = CodeVersion::V1.into();
        config.svm_chain_selector = svm_chain_selector;
        config.enable_manual_execution_after = 1;
        config.owner = ctx.accounts.authority.key();
        config.ocr3 = [
            Ocr3Config::new(OcrPluginType::Commit),
            Ocr3Config::new(OcrPluginType::Execution),
        ];

        emit!(ConfigSet {
            svm_chain_selector,
            enable_manual_execution_after: 1,
        });

        Ok(())
    }

    pub fn emit_source_chain_added(ctx: Context<EmitConfig>) -> Result<()> {
        let offramp = &ctx.accounts.offramp;
        let source_chain_selector = offramp.chain_selector;
        let on_ramp_address = OnRampAddress::from(offramp.key().to_bytes());
        let source_chain_config = SourceChainConfig {
            is_enabled: true,
            is_rmn_verification_disabled: false,
            lane_code_version: CodeVersion::V1,
            on_ramp: on_ramp_address,
        };

        emit!(SourceChainAdded {
            source_chain_selector,
            source_chain_config,
        });

        Ok(())
    }

    pub fn emit_commit_report_accepted(
        ctx: Context<EmitCommitReport>,
        min_seq_nr: u64,
        max_seq_nr: u64,
    ) -> Result<()> {
        let offramp = &ctx.accounts.offramp;
        let source_chain_selector = offramp.chain_selector;
        let on_ramp_address = offramp.on_ramp.to_bytes().to_vec();

        // Compute hash for the merkle root
        let mut hash_input = Vec::new();
        hash_input.extend_from_slice(&offramp.on_ramp.to_bytes());
        hash_input.extend_from_slice(&min_seq_nr.to_le_bytes());
        hash_input.extend_from_slice(&max_seq_nr.to_le_bytes());
        let computed_hash = hash(&hash_input).0;

        // Construct the MerkleRoot using the CCIP router type.
        let merkle = MerkleRoot {
            source_chain_selector,
            on_ramp_address,
            min_seq_nr,
            max_seq_nr,
            merkle_root: computed_hash,
        };

        let price_updates = PriceUpdates {
            gas_price_updates: vec![],
            token_price_updates: vec![],
        };

        emit!(CommitReportAccepted {
            merkle_root: Some(merkle),
            price_updates,
        });

        Ok(())
    }
}

#[account]
pub struct FiredrillOffRamp {
    pub owner: Pubkey,
    pub chain_selector: u64,
    pub token: Pubkey,
    pub on_ramp: Pubkey,
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::OFFRAMP], bump, payer = authority, space = 8 + 32 + 8 + 32 + 32)]
    pub offramp: Account<'info, FiredrillOffRamp>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct EmitConfig<'info> {
    #[account(mut, has_one = owner)]
    pub offramp: Account<'info, FiredrillOffRamp>,

    pub owner: Signer<'info>,
}

#[derive(Accounts)]
pub struct EmitCommitReport<'info> {
    /// The OffRamp account—for ownership verification.
    #[account(mut, has_one = owner)]
    pub offramp: Account<'info, FiredrillOffRamp>,
    pub owner: Signer<'info>,
}

#[error_code]
pub enum CcipOfframpError {
    #[msg("Invalid OnRamp address")]
    InvalidOnrampAddress,

    #[msg("Invalid code version")]
    InvalidCodeVersion,

    #[msg("Invalid version")]
    InvalidVersion,

    #[msg("Invalid plugin type")]
    InvalidPluginType,

    #[msg("Unauthorized")]
    Unauthorized,
}
