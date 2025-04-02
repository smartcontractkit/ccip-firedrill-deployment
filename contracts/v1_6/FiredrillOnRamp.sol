// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {HasStatus} from "../common/HasStatus.sol";
import {FiredrillToken} from "../v1_0/FiredrillToken.sol";
import {FiredrillEntrypoint} from "./FiredrillEntrypoint.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {EnumerableSet} from "@openzeppelin/utils/structs/EnumerableSet.sol";

contract FiredrillOnRamp is Ownable2Step, HasStatus {
    using EnumerableSet for EnumerableSet.AddressSet;

    event CCIPMessageSent(uint64 indexed destChainSelector, uint64 indexed sequenceNumber, EVM2AnyRampMessage message);

    enum MessageExecutionState {
      UNTOUCHED,
      IN_PROGRESS,
      SUCCESS,
      FAILURE
    }

    struct StaticConfig {
        uint64 chainSelector; // ────╮ Source chain selector.
        address rmnRemote; // ────╯ RMN remote address.
        address nonceManager; //       Nonce manager address.
        address tokenAdminRegistry; // Token admin registry address.
    }

    struct DynamicConfig {
        address feeQuoter; // FeeQuoter address.
        bool reentrancyGuardEntered; // Reentrancy protection.
        address messageInterceptor; // Optional message interceptor to validate messages. Zero address = no interceptor.
        address feeAggregator; // Fee aggregator address.
        address allowlistAdmin; // authorized admin to add or remove allowed senders.
    }

    struct DestChainConfig {
        // The last used sequence number. This is zero in the case where no messages have yet been sent.
        // 0 is not a valid sequence number for any real transaction as this value will be incremented before use.
        uint64 sequenceNumber; // ──╮ The last used sequence number.
        bool allowlistEnabled; //   │ True if the allowlist is enabled.
        address router; // ─────────╯ Local router address  that is allowed to send messages to the destination chain.
        EnumerableSet.AddressSet allowedSendersList; // The list of addresses allowed to send messages.
    }

    struct RampMessageHeader {
        bytes32 messageId; // Unique identifier for the message, generated with the source chain's encoding scheme (i.e. not necessarily abi.encoded).
        uint64 sourceChainSelector; // ─╮ the chain selector of the source chain, note: not chainId.
        uint64 destChainSelector; //    │ the chain selector of the destination chain, note: not chainId.
        uint64 sequenceNumber; //       │ sequence number, not unique across lanes.
        uint64 nonce; // ───────────────╯ nonce for this lane for this sender, not unique across senders/lanes.
    }
    struct EVM2AnyTokenTransfer {
        // The source pool EVM address. This value is trusted as it was obtained through the onRamp. It can be relied
        // upon by the destination pool to validate the source pool.
        address sourcePoolAddress;
        // The EVM address of the destination token.
        // This value is UNTRUSTED as any pool owner can return whatever value they want.
        bytes destTokenAddress;
        // Optional pool data to be transferred to the destination chain. Be default this is capped at
        // CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
        // has to be set for the specific token.
        bytes extraData;
        uint256 amount; // Amount of tokens.
        // Destination chain data used to execute the token transfer on the destination chain. For an EVM destination, it
        // consists of the amount of gas available for the releaseOrMint and transfer calls made by the offRamp.
        bytes destExecData;
    }
    struct EVM2AnyRampMessage {
        RampMessageHeader header; // Message header.
        address sender; // sender address on the source chain.
        bytes data; // arbitrary data payload supplied by the message sender.
        bytes receiver; // receiver address on the destination chain.
        bytes extraArgs; // destination-chain specific extra args, such as the gasLimit for EVM chains.
        address feeToken; // fee token.
        uint256 feeTokenAmount; // fee token amount.
        uint256 feeValueJuels; // fee amount in Juels.
        EVM2AnyTokenTransfer[] tokenAmounts; // array of tokens and amounts to transfer.
    }

    FiredrillEntrypoint private immutable i_ctrl;

    constructor(FiredrillEntrypoint ctrl) Ownable(msg.sender) {
        i_ctrl = ctrl;
    }

    function isActive() public view returns (bool) {
        return i_ctrl.isActive();
    }

    function getStaticConfig() external view returns (StaticConfig memory) {
        return StaticConfig({
            chainSelector: i_ctrl.chainSelector(),
            rmnRemote: address(i_ctrl.compound()),
            nonceManager: address(i_ctrl.compound()),
            tokenAdminRegistry: address(i_ctrl.compound())
        });
    }

    function getDynamicConfig() external view returns (DynamicConfig memory) {
        return DynamicConfig({
            feeQuoter: address(i_ctrl.compound()),
            reentrancyGuardEntered: false,
            messageInterceptor: address(0),
            feeAggregator: address(i_ctrl.compound()),
            allowlistAdmin: address(0)
        });
    }

    function getDestChainConfig(uint64 destChainSelector) external view returns (uint64 sequenceNumber, bool allowlistEnabled, address router) {
        return (0, false, address(i_ctrl.compound()));
    }

    function emitCCIPMessageSent(address sender, uint64 index) external {
        require(msg.sender == address(i_ctrl), "only control");
        EVM2AnyRampMessage memory msg = EVM2AnyRampMessage({
            header: RampMessageHeader({
                messageId: keccak256(abi.encodePacked(sender, index)),
                sourceChainSelector: i_ctrl.chainSelector(),
                destChainSelector: i_ctrl.chainSelector(),
                sequenceNumber: index,
                nonce: 1
            }),
            sender: sender,
            data: bytes("123"),
            receiver: abi.encode(sender),
            extraArgs: bytes("123"),
            feeToken: i_ctrl.token(),
            feeTokenAmount: 0,
            feeValueJuels: 0,
            tokenAmounts: new EVM2AnyTokenTransfer[](0)
        });
        emit CCIPMessageSent(i_ctrl.chainSelector(), index, msg);
    }

    function typeAndVersion() external pure returns (string memory) {
        return "OnRamp 1.6.0";
    }
}
