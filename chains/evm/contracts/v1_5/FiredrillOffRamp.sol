// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../common/Internal.sol";
import {FiredrillCompound} from "./FiredrillCompound.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";

contract FiredrillOffRamp is Ownable2Step, ITypeAndVersion {
    event TokensConsumed(uint256 tokens);
    event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
    event ReportAccepted(CommitReport report);
    event ExecutionStateChanged(uint64 indexed sequenceNumber, bytes32 indexed messageId, MessageExecutionState state, bytes returnData);

    enum MessageExecutionState {
        UNTOUCHED,
        IN_PROGRESS,
        SUCCESS,
        FAILURE
    }

    struct TokenBucket {
        uint128 tokens; // ──────╮ Current number of tokens that are in the bucket.
        uint32 lastUpdated; //   │ Timestamp in seconds of the last token refill, good for 100+ years.
        bool isEnabled; // ──────╯ Indication whether the rate limiting is enabled or not
        uint128 capacity; // ────╮ Maximum number of tokens that can be in the bucket.
        uint128 rate; // ────────╯ Number of tokens per second that the bucket is refilled.
    }

    struct StaticConfig {
        address commitStore; // ────────╮  CommitStore address on the destination chain
        uint64 chainSelector; // ───────╯  Destination chainSelector
        uint64 sourceChainSelector; // ─╮  Source chainSelector
        address onRamp; // ─────────────╯  OnRamp address on the source chain
        address prevOffRamp; //            Address of previous-version OffRamp
        address rmnProxy; //               RMN proxy address
        address tokenAdminRegistry; //     Token admin registry address
    }

    /// @notice Dynamic offRamp config
    /// @dev since OffRampConfig is part of OffRampConfigChanged event, if changing it, we should update the ABI on Atlas
    struct DynamicConfig {
        uint32 permissionLessExecutionThresholdSeconds; // ─╮ Waiting time before manual execution is enabled
        uint32 maxDataBytes; //                             │ Maximum payload data size in bytes
        uint16 maxNumberOfTokensPerMsg; //                  │ Maximum number of ERC20 token transfers that can be included per message
        address router; // ─────────────────────────────────╯ Router address
        address priceRegistry; //                             Price registry address
    }

    struct Interval {
        uint64 min; // ───╮ Minimum sequence number, inclusive
        uint64 max; // ───╯ Maximum sequence number, inclusive
    }
    struct CommitReport {
        Internal.PriceUpdates priceUpdates;
        Interval interval;
        bytes32 merkleRoot;
    }

    FiredrillCompound private immutable i_ctrl;

    constructor(FiredrillCompound ctrl) Ownable(msg.sender) {
        i_ctrl = ctrl;
    }

    function emitConfigSet() public {
        require(msg.sender == address(i_ctrl), "only control");
        emit ConfigSet(
            StaticConfig({
                commitStore: i_ctrl.offRamp(),
                chainSelector: i_ctrl.chainSelector(),
                sourceChainSelector: i_ctrl.chainSelector(),
                onRamp: i_ctrl.onRamp(),
                prevOffRamp: address(0),
                rmnProxy: address(i_ctrl),
                tokenAdminRegistry: address(i_ctrl)
            }),
            DynamicConfig({
                permissionLessExecutionThresholdSeconds: 10,
                maxDataBytes: 10,
                maxNumberOfTokensPerMsg: 10,
                router: address(i_ctrl),
                priceRegistry: address(i_ctrl)
            })
        );
    }

    function emitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) public {
        Internal.PriceUpdates memory priceUpdates;
        Interval memory interval = Interval({
            min: minSeqNr,
            max: maxSeqNr
        });
        emit ReportAccepted(
            CommitReport({
                priceUpdates: priceUpdates,
                interval: interval,
                merkleRoot: keccak256(abi.encode(i_ctrl.offRamp(), minSeqNr, maxSeqNr))
            })
        );
    }

    function emitExecutionStateChanged(address sender, uint64 index) public {
        require(msg.sender == address(i_ctrl), "only control");
        emit ExecutionStateChanged({
            sequenceNumber: 1,
            messageId: keccak256(abi.encodePacked(sender, index)),
            state: MessageExecutionState.IN_PROGRESS,
            returnData: bytes("123")
        });
    }

    function emitTokenConsumed() public {
        require(msg.sender == address(i_ctrl), "only control");
        emit TokensConsumed(100);
    }

    function getStaticConfig() external view returns (StaticConfig memory) {
        return StaticConfig({
            commitStore: i_ctrl.offRamp(),
            chainSelector: i_ctrl.chainSelector(),
            sourceChainSelector: i_ctrl.chainSelector(),
            onRamp: i_ctrl.onRamp(),
            prevOffRamp: address(0),
            rmnProxy: address(i_ctrl),
            tokenAdminRegistry: address(i_ctrl)
        });
    }

    function getDynamicConfig() external view returns (DynamicConfig memory) {
        return DynamicConfig({
            permissionLessExecutionThresholdSeconds: 10,
            maxDataBytes: 10,
            maxNumberOfTokensPerMsg: 10,
            router: address(i_ctrl),
            priceRegistry: address(i_ctrl)
        });
    }

    function currentRateLimiterState() public view returns (TokenBucket memory) {
        return TokenBucket({
            tokens: 0,
            lastUpdated: uint32(block.timestamp),
            isEnabled: true,
            capacity: 50,
            rate: 0
        });
    }

    function getToken() public view returns (address) {
        return i_ctrl.token();
    }

    function typeAndVersion() external pure returns (string memory) {
        return "OffRamp 1.5.0";
    }
}