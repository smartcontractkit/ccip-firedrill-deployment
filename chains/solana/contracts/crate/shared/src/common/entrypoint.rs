use anchor_lang::prelude::*;

#[account]
#[derive(Debug)]
pub struct FiredrillEntrypoint {
    pub owner: Pubkey,
    pub chain_selector: u64,
    pub active: bool,
    pub token: Pubkey,
    pub on_ramp: Pubkey,
    pub send_last: u8,
}
