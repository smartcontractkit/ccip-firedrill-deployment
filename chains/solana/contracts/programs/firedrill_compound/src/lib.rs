use anchor_lang::prelude::*;
use shared::ids::compound::ID;
use shared::FiredrillEntrypoint;

#[program]
pub mod firedrill_compound {
    use super::*;

    pub fn emit_usd_per_token_updated(ctx: Context<EmitUsdPerToken>) -> Result<()> {
        let entrypoint = &ctx.accounts.entrypoint;

        let token = entrypoint.token;
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

#[derive(Accounts)]
pub struct EmitUsdPerToken<'info> {
    /// CHECK: Only used for event emission context
    pub compound: AccountInfo<'info>,

    /// Must be passed in to read `.token`
    pub entrypoint: Account<'info, FiredrillEntrypoint>,
}

#[event]
pub struct UsdPerTokenUpdated {
    pub token: Pubkey,
    pub value: [u8; 28],
    pub timestamp: i64,
}
