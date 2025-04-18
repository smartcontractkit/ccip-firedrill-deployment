use anchor_lang::prelude::*;
use anchor_spl::token_interface::{Mint, TokenInterface};
use shared::seed;

declare_id!("8pNvK7oc2K9oq62Y6aBuihqEgN2q9bUH3vnmrBg1XdA");

#[program]
pub mod firedrill_token {
    use super::*;

    pub fn initialize(_ctx: Context<Initialize>) -> Result<()> {
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(
        init,
        owner = anchor_spl::token::ID,
        payer = authority,
        mint::decimals = 6,
        mint::authority = mint.key(),
        mint::freeze_authority = mint.key(),
        seeds = [seed::TOKEN],
        bump
    )]
    pub mint: InterfaceAccount<'info, Mint>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub token_program: Interface<'info, TokenInterface>,
    pub system_program: Program<'info, System>,
    pub rent: Sysvar<'info, Rent>,
}
