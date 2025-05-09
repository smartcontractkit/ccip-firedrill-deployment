use anchor_lang::prelude::*;
use ethnum::U256;

use shared::seed;

pub mod messages;
use crate::messages::*;

pub mod state;
use crate::state::*;

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
        receiver: Pubkey,
    ) -> Result<()> {
        // Initialize entrypoint state with core routing information
        let state = &mut ctx.accounts.entrypoint;
        state.owner = ctx.accounts.authority.key();
        state.chain_selector = chain_selector;
        state.token = token;
        state.off_ramp = off_ramp;
        state.fee_quoter = fee_quoter;
        state.receiver = receiver;
        state.send_last = 0;

        // Initialize config for the program (merged from compound)
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

        // Initialize destination chain configuration (merged from compound)
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
        let entrypoint = &ctx.accounts.entrypoint;
        let message = SVM2AnyRampMessage {
            header: RampMessageHeader {
                message_id: [0u8; 32],
                source_chain_selector: entrypoint.chain_selector,
                dest_chain_selector: entrypoint.chain_selector,
                sequence_number: index,
                nonce: 1,
            },
            sender,
            data: vec![],
            receiver: sender.to_bytes().to_vec(),
            extra_args: vec![],
            fee_token: entrypoint.token,
            token_amounts: vec![],
            fee_token_amount: U256::from(0u64).into(),
            fee_value_juels: U256::from(0u64).into(),
        };

        emit!(CCIPMessageSent {
            dest_chain_selector: entrypoint.chain_selector,
            sequence_number: index,
            message,
        });

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
    pub receiver: Pubkey,
    pub send_last: u8,
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
#[instruction(chain_selector: u64, token: Pubkey, off_ramp: Pubkey, fee_quoter: Pubkey, receiver: Pubkey)]
pub struct Initialize<'info> {
    #[account(init, seeds = [seed::ENTRYPOINT], bump, payer = authority, space = 8 + 32 + 8 + 32 + 32 + 32 + 32 + 1)]
    pub entrypoint: Account<'info, FiredrillEntrypoint>,

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
    pub entrypoint: Account<'info, FiredrillEntrypoint>,
    pub owner: Signer<'info>,
}

#[event]
pub struct CCIPMessageSent {
    pub dest_chain_selector: u64,
    pub sequence_number: u64,
    pub message: SVM2AnyRampMessage,
}