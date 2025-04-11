// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {HasStatus} from "../common/HasStatus.sol";
import {FiredrillEntrypoint} from "./FiredrillEntrypoint.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";

contract FiredrillCompound is Ownable2Step, HasStatus, ITypeAndVersion {
    event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp);
    event Released(address indexed sender, address indexed recipient, uint256 amount);
    event TaggedRootBlessed(uint32 indexed configVersion, TaggedRoot taggedRoot, uint16 accumulatedWeight);
    event Cursed(uint32 indexed configVersion, bytes16 subject, uint64 blockTimestamp);
    event VotedToCurse(
      uint32 indexed configVersion,
      address indexed voter,
      bytes16 subject,
      bytes16 curseId,
      uint8 weight,
      uint64 blockTimestamp,
      bytes28 cursesHash,
      uint16 accumulatedWeight
    );

    struct TokenBucket {
        uint128 tokens; // ──────╮ Current number of tokens that are in the bucket.
        uint32 lastUpdated; //   │ Timestamp in seconds of the last token refill, good for 100+ years.
        bool isEnabled; // ──────╯ Indication whether the rate limiting is enabled or not
        uint128 capacity; // ────╮ Maximum number of tokens that can be in the bucket.
        uint128 rate; // ────────╯ Number of tokens per second that the bucket is refilled.
    }

    struct TaggedRoot {
        address commitStore;
        bytes32 root;
    }
    struct OffRamp {
        uint64 sourceChainSelector;
        address offRamp;
    }

    FiredrillEntrypoint internal immutable i_ctrl;

    constructor(FiredrillEntrypoint ctrl) Ownable(msg.sender) {
        i_ctrl = ctrl;
    }

    function isActive() public view returns (bool) {
        return i_ctrl.isActive();
    }

    function emitUsdPerTokenUpdated() public {
        emit UsdPerTokenUpdated({
            token: i_ctrl.token(),
            value: 1, 
            timestamp: block.timestamp
        });
    }

    function emitReleased() public {
        emit Released({
            sender:  address(this), 
            recipient: msg.sender, 
            amount: 100
        });
    }

    function emitTagRootBlessed(uint64 minSeqNr, uint64 maxSeqNr) public {
        TaggedRoot memory taggedRoot = TaggedRoot({
            commitStore: i_ctrl.offRamp(),
            root: keccak256(abi.encode(i_ctrl.offRamp(), minSeqNr, maxSeqNr))
        });
        emit TaggedRootBlessed(0, taggedRoot, 0);
    }

    function emitCursed() public {
        emit Cursed({
            configVersion: 0,
            subject: 0,
            blockTimestamp: uint64(block.timestamp)
        });
    }

    function emitVotedToCurse() public {
        emit VotedToCurse({
            configVersion: 0,
            voter: msg.sender,
            subject: 0,
            curseId: 0,
            weight: 3,
            blockTimestamp: uint64(block.timestamp),
            cursesHash: 0,
            accumulatedWeight: 0
        });
    }

    function getOnRamp(uint64) external view returns (address) {
        return i_ctrl.onRamp();
    }

    function getOffRamps() external view returns (OffRamp[] memory) {
        OffRamp[] memory offRamps = new OffRamp[](1);
        offRamps[0] = OffRamp({
            sourceChainSelector: i_ctrl.chainSelector(),
            offRamp: i_ctrl.offRamp()
        });
        return offRamps;
    }

    function getStalenessThreshold() external pure returns (uint128) {
        return 0;
    }

    function getCurrentInboundRateLimiterState(uint64) public view returns (TokenBucket memory) {
        return TokenBucket({
            tokens: 0,
            lastUpdated: uint32(block.timestamp),
            isEnabled: true,
            capacity: 50,
            rate: 0
        });
    }

    function getCurrentOutboundRateLimiterState(uint64) public view returns (TokenBucket memory) {
        return TokenBucket({
            tokens: 0,
            lastUpdated: uint32(block.timestamp),
            isEnabled: true,
            capacity: 50,
            rate: 0
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

    function getARM() public view returns (address) {
        return i_ctrl.compound();
    }

    function typeAndVersion() external pure returns (string memory) {
        return "Router 1.2.0";
    }
}