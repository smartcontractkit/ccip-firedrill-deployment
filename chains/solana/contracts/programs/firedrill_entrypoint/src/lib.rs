use anchor_lang::prelude::*;
use shared::ids::entrypoint::ID;
use shared::common::seed;

use firedrill_compound::program::FiredrillCompound;
use firedrill_offramp::cpi::accounts::EmitCommitReport as OffRampCommitReport;
use firedrill_offramp::program::FiredrillOfframp;
use firedrill_onramp::cpi::accounts::EmitMessage as OnRampEmitMessage;
use firedrill_onramp::program::FiredrillOnramp;

#[program]
pub mod firedrill_entrypoint {
    use super::*;

    pub fn initialize(
        ctx: Context<Initialize>,
        chain_selector: u64,
        token: Pubkey,
        on_ramp: Pubkey,
        off_ramp: Pubkey,
        compound: Pubkey,
        receiver: Pubkey,
    ) -> Result<()> {
        let state = &mut ctx.accounts.entrypoint;
        state.owner = ctx.accounts.authority.key();
        state.chain_selector = chain_selector;
        state.token = token;
        state.on_ramp = on_ramp;
        state.off_ramp = off_ramp;
        state.compound = compound;
        state.receiver = receiver;
        state.send_last = 0;
        Ok(())
    }

    pub fn prepare_register(ctx: Context<PrepareRegister>) -> Result<()> {
        firedrill_offramp::cpi::emit_source_chain_added(CpiContext::new(
            ctx.accounts.offramp_program.to_account_info(),
            firedrill_offramp::cpi::accounts::EmitConfig {
                offramp: ctx.accounts.offramp.to_account_info(),
                owner: ctx.accounts.owner.to_account_info(),
            },
        ))
    }

    pub fn drill_pending_commit_queue_tx_spike(
        ctx: Context<OnlyOwner>,
        from: u8,
        to: u8,
    ) -> Result<()> {
        let ep = &mut ctx.accounts.entrypoint;

        require!(from <= to, CustomError::NothingToSend);
        require!(from > ep.send_last, CustomError::AlreadySent);

        for i in from..=to {
            firedrill_onramp::cpi::emit_ccip_message_sent(
                CpiContext::new(
                    ctx.accounts.onramp_program.to_account_info(),
                    OnRampEmitMessage {
                        onramp: ctx.accounts.onramp.to_account_info(),
                        owner: ctx.accounts.owner.to_account_info(),
                        caller_program: ep.clone().to_account_info(),
                    },
                ),
                ctx.program_id.key(),
                i as u64,
            )?;
        }

        ep.send_last = to;
        Ok(())
    }

    pub fn drill_pending_execution(ctx: Context<OnlyOwner>, from: u8, to: u8) -> Result<()> {
        require!(from <= to, CustomError::NothingToSend);
        require!(
            to <= ctx.accounts.entrypoint.send_last,
            CustomError::NotYetSent
        );

        firedrill_offramp::cpi::emit_commit_report_accepted(
            CpiContext::new(
                ctx.accounts.offramp_program.to_account_info(),
                OffRampCommitReport {
                    offramp: ctx.accounts.offramp.to_account_info(),
                    owner: ctx.accounts.owner.to_account_info(),
                },
            ),
            from.into(),
            to.into(),
        )?;

        Ok(())
    }

    pub fn drill_price_registries(ctx: Context<OnlyOwner>) -> Result<()> {
        firedrill_compound::cpi::emit_usd_per_token_updated(CpiContext::new(
            ctx.accounts.compound_program.to_account_info(),
            firedrill_compound::cpi::accounts::EmitUsdPerToken {
                compound: ctx.accounts.compound.to_account_info(),
            },
        ))?;

        Ok(())
    }
}

#[account]
#[derive(Debug)]
pub struct FiredrillEntrypoint {
    pub owner: Pubkey,
    pub chain_selector: u64,
    pub token: Pubkey,
    pub on_ramp: Pubkey,
    pub off_ramp: Pubkey,
    pub compound: Pubkey,
    pub receiver: Pubkey,
    pub send_last: u8,
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::ENTRYPOINT], bump, payer = authority, space = 8 + 32 + 8 + 32 + 32 + 32 + 32 + 32 + 1)]
    pub entrypoint: Account<'info, FiredrillEntrypoint>,
    #[account(mut)]
    pub authority: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct PrepareRegister<'info> {
    #[account(mut, has_one = owner)]
    pub entrypoint: Account<'info, FiredrillEntrypoint>,

    /// CHECK: account is passed to the OffRamp CPI, verified inside that program
    #[account(mut)]
    pub offramp: AccountInfo<'info>,

    /// CHECK: only passed to CPI, validated in called program
    pub owner: Signer<'info>,

    /// CHECK: CPI target program
    pub offramp_program: Program<'info, FiredrillOfframp>,
}

#[derive(Accounts)]
pub struct OnlyOwner<'info> {
    #[account(mut, has_one = owner)]
    pub entrypoint: Account<'info, FiredrillEntrypoint>,
    pub owner: Signer<'info>,

    /// CHECK: CPI
    #[account(mut)]
    pub onramp: AccountInfo<'info>,

    /// CHECK: CPI
    #[account(mut)]
    pub offramp: AccountInfo<'info>,

    /// CHECK: CPI
    #[account(mut)]
    pub compound: AccountInfo<'info>,

    pub onramp_program: Program<'info, FiredrillOnramp>,
    pub offramp_program: Program<'info, FiredrillOfframp>,
    pub compound_program: Program<'info, FiredrillCompound>,
}

#[error_code]
pub enum CustomError {
    #[msg("Nothing to send.")]
    NothingToSend,
    #[msg("Message already sent.")]
    AlreadySent,
    #[msg("Message not yet sent.")]
    NotYetSent,
}
