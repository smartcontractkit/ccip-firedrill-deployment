use anchor_lang::prelude::*;

use crate::context::{MerkleRoot, PriceUpdates};
use crate::state::{MessageExecutionState, SourceChainConfig};

#[event]
pub struct CommitReportAccepted {
    pub merkle_root: Option<MerkleRoot>,
    pub price_updates: PriceUpdates,
}

#[event]
pub struct ExecutionStateChanged {
    pub source_chain_selector: u64,
    pub sequence_number: u64,
    pub message_id: [u8; 32],
    pub message_hash: [u8; 32],
    pub state: MessageExecutionState,
}

#[event]
pub struct SourceChainAdded {
    pub source_chain_selector: u64,
    pub source_chain_config: SourceChainConfig,
}

#[event]
pub struct ConfigSet {
    pub svm_chain_selector: u64,
    pub enable_manual_execution_after: i64,
}
