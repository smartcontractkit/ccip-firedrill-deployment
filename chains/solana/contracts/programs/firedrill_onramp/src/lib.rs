use anchor_lang::prelude::*;
use shared::ids::entrypoint::ID as ENTRYPOINT_ID;
use shared::ids::onramp::ID;
use shared::common::seed;

pub mod messages;
use crate::messages::*;

use ethnum::U256;

#[program]
pub mod firedrill_onramp {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, chain_selector: u64, token: Pubkey) -> Result<()> {
        let onramp = &mut ctx.accounts.onramp;
        onramp.owner = ctx.accounts.authority.key();
        onramp.chain_selector = chain_selector;
        onramp.token = token;
        Ok(())
    }

    pub fn emit_ccip_message_sent(
        ctx: Context<EmitMessage>,
        sender: Pubkey,
        index: u64,
    ) -> Result<()> {
        require!(
            ctx.accounts.caller_program.key() == ENTRYPOINT_ID,
            CustomError::NotEntrypointCaller
        );

        let onramp = &ctx.accounts.onramp;
        let message = SVM2AnyRampMessage {
            header: RampMessageHeader {
                message_id: [0u8; 32],
                source_chain_selector: onramp.chain_selector,
                dest_chain_selector: onramp.chain_selector,
                sequence_number: index,
                nonce: 1,
            },
            sender,
            data: vec![],
            receiver: sender.to_bytes().to_vec(),
            extra_args: vec![],
            fee_token: onramp.token,
            token_amounts: vec![],
            fee_token_amount: U256::from(0u64).into(),
            fee_value_juels: U256::from(0u64).into(),
        };

        emit!(CCIPMessageSent {
            dest_chain_selector: onramp.chain_selector,
            sequence_number: index,
            message,
        });

        Ok(())
    }
}

#[account]
pub struct FiredrillOnRamp {
    pub owner: Pubkey,
    pub chain_selector: u64,
    pub token: Pubkey,
}

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::ONRAMP], bump, payer = authority, space = 8 + 32 + 8 + 32)]
    pub onramp: Account<'info, FiredrillOnRamp>,

    #[account(mut)]
    pub authority: Signer<'info>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct EmitMessage<'info> {
    #[account(mut)]
    pub onramp: Account<'info, FiredrillOnRamp>,
    /// CHECK: Must be FiredrillEntrypoint
    pub caller_program: AccountInfo<'info>,

    pub owner: Signer<'info>,
}

#[event]
pub struct CCIPMessageSent {
    pub dest_chain_selector: u64,
    pub sequence_number: u64,
    pub message: SVM2AnyRampMessage,
}

#[error_code]
pub enum CustomError {
    #[msg("Only entrypoint allowed to call.")]
    NotEntrypointCaller,
}
