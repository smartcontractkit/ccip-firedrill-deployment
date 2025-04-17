use anchor_lang::prelude::*;
use shared::ids::compound::ID;
use shared::common::seed;

#[program]
pub mod firedrill_compound {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, token: Pubkey) -> Result<()> {
        let compound = &mut ctx.accounts.compound;
        compound.owner = ctx.accounts.authority.key();
        compound.token = token;
        Ok(())
    }

    pub fn emit_usd_per_token_updated(ctx: Context<EmitUsdPerToken>) -> Result<()> {
        let compound = &ctx.accounts.compound;

        let token = compound.token;
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
 pub struct FiredrillCompound {
     pub owner: Pubkey,
     pub token: Pubkey,
 }

 #[derive(Accounts)]
 pub struct Initialize<'info> {
     #[account(init, seeds = [seed::COMPOUND], bump, payer = authority, space = 8 + 32 + 32)]
     pub compound: Account<'info, FiredrillCompound>,

     #[account(mut)]
     pub authority: Signer<'info>,

     pub system_program: Program<'info, System>,
 }

#[derive(Accounts)]
pub struct EmitUsdPerToken<'info> {
    #[account(mut)]
    pub compound: Account<'info, FiredrillCompound>,
}

#[event]
pub struct UsdPerTokenUpdated {
    pub token: Pubkey,
    pub value: [u8; 28],
    pub timestamp: i64,
}
