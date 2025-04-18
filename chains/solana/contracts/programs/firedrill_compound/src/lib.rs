use anchor_lang::prelude::*;
use shared::seed;

pub mod messages;
use crate::messages::*;

pub mod state;
use crate::state::*;

use ethnum::U256;

declare_id!("mjNGiAfCXTWqcD5TSy1wJ65g3kmwepud9NYxj1B36oy");

// FiredrillCompound is Router, OnRamp, RMNRemote, FeeQuoter and TokenAdminRegistry
#[program]
pub mod firedrill_compound {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, chain_selector: u64, fee_quoter: Pubkey, token: Pubkey) -> Result<()> {
        let compound = &mut ctx.accounts.compound;
        compound.chain_selector = chain_selector;
        compound.owner = ctx.accounts.authority.key();
        compound.fee_quoter = fee_quoter;
        compound.token = token;

        ctx.accounts.config.set_inner(Config {
            version: 1,
            default_code_version: CodeVersion::V1,
            owner: ctx.accounts.authority.key(),
            proposed_owner: Pubkey::default(),
            svm_chain_selector: chain_selector,
            fee_quoter,
            rmn_remote: ID,
            link_token_mint: token,
            fee_aggregator: ID,
        });

        ctx.accounts.dest_chain.set_inner(DestChain {
            version: 1,
            chain_selector,
            state: DestChainState {
                sequence_number: 1,
                sequence_number_to_restore: 1,
                restore_on_action: RestoreOnAction::None,
            },
            config: DestChainConfig {
                lane_code_version: CodeVersion::V1,
                allowed_senders: vec![],
                allow_list_enabled: false,
            },
        });

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
}

#[account]
pub struct FiredrillCompound {
    pub chain_selector: u64,
    pub owner: Pubkey,
    pub fee_quoter: Pubkey,
    pub token: Pubkey,
}

#[account]
#[derive(InitSpace, Debug)]
pub struct DestChain {
    // Config for SVM2Any
    pub version: u8,
    pub chain_selector: u64,     // Chain selector used for the seed
    pub state: DestChainState,   // values that are updated automatically
    pub config: DestChainConfig, // values configured by an admin
}

#[account]
#[derive(InitSpace, Debug)]
pub struct Config {
    pub version: u8,

    pub default_code_version: CodeVersion,

    pub svm_chain_selector: u64,
    pub owner: Pubkey,
    pub proposed_owner: Pubkey,
    pub fee_quoter: Pubkey,
    pub rmn_remote: Pubkey,
    pub link_token_mint: Pubkey,
    pub fee_aggregator: Pubkey, // Allowed address to withdraw billed fees to (will use ATAs derived from it)
}

#[derive(Accounts)]
#[instruction(chain_selector: u64, fee_quoter: Pubkey, token: Pubkey)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::COMPOUND], bump, payer = authority, space = 8 + 8 + 32 + 32 + 32)]
    pub compound: Account<'info, FiredrillCompound>,

    #[account(
        init,
        seeds = [seed::CONFIG],
        bump,
        payer = authority,
        space = 8 + Config::INIT_SPACE,
    )]
    pub config: Account<'info, Config>,

    #[account(
        init,
        seeds = [seed::DEST_CHAIN_STATE, chain_selector.to_le_bytes().as_ref()],
        bump,
        payer = authority,
        space = 8 + DestChain::INIT_SPACE + 4, // 4 = empty Vec<Pubkey> len prefix
    )]
    pub dest_chain: Account<'info, DestChain>,

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

#[event]
pub struct CCIPMessageSent {
    pub dest_chain_selector: u64,
    pub sequence_number: u64,
    pub message: SVM2AnyRampMessage,
}
