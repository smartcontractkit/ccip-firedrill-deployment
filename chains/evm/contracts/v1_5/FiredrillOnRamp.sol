// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FiredrillToken} from "../v1_0/FiredrillToken.sol";
import {FiredrillCompound} from "./FiredrillCompound.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";

contract FiredrillOnRamp is Ownable2Step, ITypeAndVersion {
    event CCIPSendRequested(EVM2EVMMessage message);

    struct EVM2EVMMessage {
        uint64 sourceChainSelector; // ────────╮ the chain selector of the source chain, note: not chainId
        address sender; // ────────────────────╯ sender address on the source chain
        address receiver; // ──────────────────╮ receiver address on the destination chain
        uint64 sequenceNumber; // ─────────────╯ sequence number, not unique across lanes
        uint256 gasLimit; //                     user supplied maximum gas amount available for dest chain execution
        bool strict; // ───────────────────────╮ DEPRECATED
        uint64 nonce; //                       │ nonce for this lane for this sender, not unique across senders/lanes
        address feeToken; // ──────────────────╯ fee token
        uint256 feeTokenAmount; //               fee token amount
        bytes data; //                           arbitrary data payload supplied by the message sender
        EVMTokenAmount[] tokenAmounts; // array of tokens and amounts to transfer
        bytes[] sourceTokenData; //              array of token data, one per token
        bytes32 messageId; //                    a hash of the message data
    }

    struct EVMTokenAmount {
        address token; // token address on the local chain.
        uint256 amount; // Amount of tokens.
    }

    struct StaticConfig {
        address linkToken; // ────────╮ Link token address
        uint64 chainSelector; // ─────╯ Source chainSelector
        uint64 destChainSelector; // ─╮ Destination chainSelector
        uint64 defaultTxGasLimit; //  │ Default gas limit for a tx
        uint96 maxNopFeesJuels; // ───╯ Max nop fee balance onramp can have
        address prevOnRamp; //          Address of previous-version OnRamp
        address rmnProxy; //            Address of RMN proxy
        address tokenAdminRegistry; //  Address of the token admin registry
    }

    struct DynamicConfig {
        address router; // ──────────────────────────╮ Router address
        uint16 maxNumberOfTokensPerMsg; //           │ Maximum number of distinct ERC20 token transferred per message
        uint32 destGasOverhead; //                   │ Gas charged on top of the gasLimit to cover destination chain costs
        uint16 destGasPerPayloadByte; //             │ Destination chain gas charged for passing each byte of `data` payload to receiver
        uint32 destDataAvailabilityOverheadGas; // ──╯ Extra data availability gas charged on top of the message, e.g. for OCR
        uint16 destGasPerDataAvailabilityByte; // ───╮ Amount of gas to charge per byte of message data that needs availability
        uint16 destDataAvailabilityMultiplierBps; // │ Multiplier for data availability gas, multiples of bps, or 0.0001
        address priceRegistry; //                    │ Price registry address
        uint32 maxDataBytes; //                      │ Maximum payload data size in bytes
        uint32 maxPerMsgGasLimit; // ────────────────╯ Maximum gas limit for messages targeting EVMs
        //                                           │
        // The following three properties are defaults, they can be overridden by setting the TokenTransferFeeConfig for a token
        uint16 defaultTokenFeeUSDCents; // ──────────╮ Default token fee charged per token transfer
        uint32 defaultTokenDestGasOverhead; //       │ Default gas charged to execute the token transfer on the destination chain
        //                                           │ Default data availability bytes that are returned from the source pool and sent
        uint32 defaultTokenDestBytesOverhead; //     | to the destination pool. Must be >= Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES
        bool enforceOutOfOrder; // ──────────────────╯ Whether to enforce the allowOutOfOrderExecution extraArg value to be true.
    }

    FiredrillCompound private immutable i_ctrl;

    constructor(FiredrillCompound ctrl) Ownable(msg.sender) {
        i_ctrl = ctrl;
    }

    function getStaticConfig() external view returns (StaticConfig memory) {
        return StaticConfig({
            linkToken: i_ctrl.token(),
            chainSelector: i_ctrl.chainSelector(),
            destChainSelector: i_ctrl.chainSelector(),
            defaultTxGasLimit: 0,
            maxNopFeesJuels: 0,
            prevOnRamp: address(0),
            rmnProxy: address(i_ctrl),
            tokenAdminRegistry: address(i_ctrl)
        });
    }

    function getDynamicConfig() external view returns (DynamicConfig memory dynamicConfig) {
        return DynamicConfig({
            router: address(i_ctrl),
            maxNumberOfTokensPerMsg: 0,
            destGasOverhead: 0,
            destGasPerPayloadByte: 0,
            destDataAvailabilityOverheadGas: 0,
            destGasPerDataAvailabilityByte: 0,
            destDataAvailabilityMultiplierBps: 0,
            priceRegistry: address(i_ctrl),
            maxDataBytes: 0,
            maxPerMsgGasLimit: 0,
            defaultTokenFeeUSDCents: 0,
            defaultTokenDestGasOverhead: 0,
            defaultTokenDestBytesOverhead: 0,
            enforceOutOfOrder: false
        });
    }

    function emitCCIPSendRequested(address sender, uint64 index) public {
        require(msg.sender == address(i_ctrl), "only control");
        emit CCIPSendRequested(EVM2EVMMessage({
            sourceChainSelector: i_ctrl.chainSelector(),
            sender: sender,
            receiver: sender,
            sequenceNumber: index,
            gasLimit: 1000,
            strict: false,
            nonce: 1,
            feeToken: i_ctrl.token(),
            feeTokenAmount: 0,
            data: bytes("123"),
            tokenAmounts: new EVMTokenAmount[](0),
            sourceTokenData: new bytes[](0),
            messageId: keccak256(abi.encodePacked(sender, index))
        }));
    }

    function typeAndVersion() external pure returns (string memory) {
        return "OnRamp 1.5.0";
    }
}