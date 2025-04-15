use anchor_lang::prelude::*;
use solana_program::keccak::hash;

use shared::ids::offramp::ID;
use shared::FiredrillEntrypoint;

mod state;
use crate::state::*;

mod event;
use crate::event::*;

mod context;
use crate::context::*;

#[program]
pub mod firedrill_offramp {
    use super::*;

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
        enable_execution_after: i64,
    ) -> Result<()> {
        let mut config = ctx.accounts.config.load_init()?;
        require!(config.version == 0, CcipOfframpError::InvalidVersion); // assert uninitialized state - AccountLoader doesn't work with constraint
        config.version = 1;
        config.default_code_version = CodeVersion::V1.into();
        config.svm_chain_selector = svm_chain_selector;
        config.enable_manual_execution_after = enable_execution_after;
        config.owner = ctx.accounts.authority.key();
        config.ocr3 = [
            Ocr3Config::new(OcrPluginType::Commit),
            Ocr3Config::new(OcrPluginType::Execution),
        ];

        emit!(ConfigSet {
            svm_chain_selector,
            enable_manual_execution_after: enable_execution_after,
        });

        Ok(())
    }

    pub fn emit_source_chain_added(ctx: Context<EmitConfig>) -> Result<()> {
        let entrypoint = &ctx.accounts.entrypoint;
        let source_chain_selector = entrypoint.chain_selector;
        let on_ramp_address = OnRampAddress::from(entrypoint.key().to_bytes());
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
        let entrypoint = &ctx.accounts.entrypoint;
        let source_chain_selector = entrypoint.chain_selector;
        let on_ramp_address = entrypoint.on_ramp.to_bytes().to_vec();

        // Compute hash for the merkle root
        let mut hash_input = Vec::new();
        hash_input.extend_from_slice(&entrypoint.on_ramp.to_bytes());
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
}

#[derive(Accounts)]
pub struct EmitConfig<'info> {
    #[account(mut, has_one = owner)]
    pub offramp: Account<'info, FiredrillOffRamp>,

    /// CHECK: Just reading pubkey + data
    pub entrypoint: Account<'info, FiredrillEntrypoint>,

    pub owner: Signer<'info>,
}

#[derive(Accounts)]
pub struct EmitCommitReport<'info> {
    /// The OffRamp account—for ownership verification.
    #[account(mut, has_one = owner)]
    pub offramp: Account<'info, FiredrillOffRamp>,
    /// The control (entrypoint) account that holds chainSelector and onRamp.
    /// Use CHECK if the account does not need further verification.
    /// (Alternatively, if you have a full account type, you can specify it directly.)
    pub entrypoint: Account<'info, FiredrillEntrypoint>,
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
