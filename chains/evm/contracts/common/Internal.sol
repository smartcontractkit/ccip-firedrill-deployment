// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

library Internal {
    struct TokenPriceUpdate {
        address sourceToken; // Source token
        uint224 usdPerToken; // 1e18 USD per 1e18 of the smallest token denomination.
    }
    struct GasPriceUpdate {
        uint64 destChainSelector; // Destination chain selector
        uint224 usdPerUnitGas; // 1e18 USD per smallest unit (e.g. wei) of destination chain gas
    }
    struct PriceUpdates {
        TokenPriceUpdate[] tokenPriceUpdates;
        GasPriceUpdate[] gasPriceUpdates;
    }
}