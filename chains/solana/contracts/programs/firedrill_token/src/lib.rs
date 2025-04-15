use anchor_lang::prelude::*;
use shared::ids::token::ID;

#[program]
pub mod firedrill_token {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        let token = &mut ctx.accounts.token;
        token.entrypoint = ctx.accounts.entrypoint.key();
        Ok(())
    }
}

#[account]
pub struct FiredrillToken {
    pub entrypoint: Pubkey,
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(init, payer = payer, space = 8 + 32)]
    pub token: Account<'info, FiredrillToken>,

    /// CHECK: We're just saving the pubkey
    pub entrypoint: AccountInfo<'info>,

    #[account(mut)]
    pub payer: Signer<'info>,

    pub system_program: Program<'info, System>,
}
