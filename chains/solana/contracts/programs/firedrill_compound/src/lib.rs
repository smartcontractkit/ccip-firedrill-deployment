use anchor_lang::prelude::*;
use shared::seed;

pub mod messages;
use crate::messages::*;

use ethnum::U256;

declare_id!("9foztXMAs3Zdgx7ZkxDksb2S6piKE6R5G9iM5oAuVwFP");

// FiredrillCompound is Router, OnRamp, RMNRemote, FeeQuoter and TokenAdminRegistry
#[program]
pub mod firedrill_compound {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, token: Pubkey, chain_selector: u64) -> Result<()> {
        let compound = &mut ctx.accounts.compound;
        compound.chain_selector = chain_selector;
        compound.owner = ctx.accounts.authority.key();
        compound.token = token;
        Ok(())
    }

    pub fn emit_ccip_message_sent(
        ctx: Context<EmitMessage>,
        sender: Pubkey,
        index: u64,
    ) -> Result<()> {
        let compound = &ctx.accounts.compound;
        let message = SVM2AnyRampMessage {
            header: RampMessageHeader {
                message_id: [0u8; 32],
                source_chain_selector: compound.chain_selector,
                dest_chain_selector: compound.chain_selector,
                sequence_number: index,
                nonce: 1,
            },
            sender,
            data: vec![],
            receiver: sender.to_bytes().to_vec(),
            extra_args: vec![],
            fee_token: compound.token,
            token_amounts: vec![],
            fee_token_amount: U256::from(0u64).into(),
            fee_value_juels: U256::from(0u64).into(),
        };

        emit!(CCIPMessageSent {
            dest_chain_selector: compound.chain_selector,
            sequence_number: index,
            message,
        });

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
    pub chain_selector: u64,
    pub owner: Pubkey,
    pub token: Pubkey,
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::COMPOUND], bump, payer = authority, space = 8 + 8 + 32 + 32)]
    pub compound: Account<'info, FiredrillCompound>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct EmitMessage<'info> {
    #[account(mut, has_one = owner)]
    pub compound: Account<'info, FiredrillCompound>,
    pub owner: Signer<'info>,
}

#[derive(Accounts)]
pub struct EmitUsdPerToken<'info> {
    #[account(mut, has_one = owner)]
    pub compound: Account<'info, FiredrillCompound>,
    pub owner: Signer<'info>,
}

#[event]
pub struct CCIPMessageSent {
    pub dest_chain_selector: u64,
    pub sequence_number: u64,
    pub message: SVM2AnyRampMessage,
}

#[event]
pub struct UsdPerTokenUpdated {
    pub token: Pubkey,
    pub value: [u8; 28],
    pub timestamp: i64,
}
