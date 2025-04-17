use anchor_lang::prelude::*;
use shared::seed;

declare_id!("F5jMhWDPMrCbnCaw6mf91NtsRV4ccEuwGsrJUpBckQoD");

#[program]
pub mod firedrill_token {}

#[account]
pub struct FiredrillToken {}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::TOKEN], bump, payer = authority, space = 8)]
    pub token: Account<'info, FiredrillToken>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}
