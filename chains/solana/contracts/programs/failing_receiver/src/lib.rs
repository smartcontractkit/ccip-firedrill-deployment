use anchor_lang::prelude::*;

declare_id!("Dri11e51CFemqdySVVhmJbRzKeyVBJjch49RF2zH6Nt6");

#[program]
pub mod failing_receiver {
    use super::*;

    pub fn ccip_receive(_ctx: Context<CcipReceive>, _message: Any2SVMMessage) -> Result<()> {
        return err!(CustomError::AlwaysErrors);
    }
}

#[derive(Accounts, Debug)]
pub struct CcipReceive {}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize)]
pub struct Any2SVMMessage {
    pub message_id: [u8; 32],
    pub source_chain_selector: u64,
    pub sender: Vec<u8>,
    pub data: Vec<u8>,
    pub token_amounts: Vec<SVMTokenAmount>,
}

#[derive(Debug, Clone, AnchorSerialize, AnchorDeserialize, Default)]
pub struct SVMTokenAmount {
    pub token: Pubkey,
    pub amount: u64, // solana local token amount
}

#[error_code]
pub enum CustomError {
    #[msg("Always errors.")]
    AlwaysErrors,
}