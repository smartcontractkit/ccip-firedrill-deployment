use anchor_lang::prelude::*;
use anchor_spl::token::{Mint, Token, InitializeMint};
use shared::seed;

declare_id!("FVPaqWNMPn7DczMYAz5HnLXEoT4cg1HhYLd9HuGAPHrn");

#[program]
pub mod firedrill_token {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, decimals: u8) -> Result<()> {
        let cpi_ctx = CpiContext::new(
            ctx.accounts.token_program.to_account_info(),
            InitializeMint {
                mint: ctx.accounts.mint.to_account_info(),
                rent: ctx.accounts.rent.to_account_info(),
            },
        );
        anchor_spl::token::initialize_mint(
            cpi_ctx,
            decimals,
            ctx.accounts.payer.key,
            None,
        )?;
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(
        init,
        payer = payer,
        space = Mint::LEN,
        seeds = [seed::TOKEN],
        bump,
    )]
    pub mint: Account<'info, Mint>,

    #[account(mut)]
    pub payer: Signer<'info>,

    #[account(address = anchor_spl::token::ID)]
    pub token_program: Program<'info, Token>,

    pub system_program: Program<'info, System>,
    pub rent: Sysvar<'info, Rent>,
}
