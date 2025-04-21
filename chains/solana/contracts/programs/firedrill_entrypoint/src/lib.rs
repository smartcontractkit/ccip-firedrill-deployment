use anchor_lang::prelude::*;
use shared::seed;

declare_id!("351TRJhiMuGnf9YJjhyVA1HHGDAUbwq1U7ppQQZNhsqj");

#[program]
pub mod firedrill_entrypoint {
    use super::*;

    pub fn initialize(
        ctx: Context<Initialize>,
        chain_selector: u64,
        token: Pubkey,
        off_ramp: Pubkey,
        fee_quoter: Pubkey,
        compound: Pubkey,
        receiver: Pubkey,
    ) -> Result<()> {
        let state = &mut ctx.accounts.entrypoint;
        state.owner = ctx.accounts.authority.key();
        state.chain_selector = chain_selector;
        state.token = token;
        state.off_ramp = off_ramp;
        state.fee_quoter = fee_quoter;
        state.compound = compound;
        state.receiver = receiver;
        state.send_last = 0;
        Ok(())
    }
}

#[account]
#[derive(Debug)]
pub struct FiredrillEntrypoint {
    pub owner: Pubkey,
    pub chain_selector: u64,
    pub token: Pubkey,
    pub off_ramp: Pubkey,
    pub fee_quoter: Pubkey,
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