// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../common/Internal.sol";
import {FiredrillCompound} from "./FiredrillCompound.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";

contract FiredrillOffRamp is Ownable2Step {
    event SourceChainConfigSet(uint64 indexed sourceChainSelector, SourceChainConfig sourceConfig);
    event CommitReportAccepted(
        MerkleRoot[] blessedMerkleRoots,
        MerkleRoot[] unblessedMerkleRoots,
        Internal.PriceUpdates priceUpdates
    );
    
    enum MessageExecutionState {
      UNTOUCHED,
      IN_PROGRESS,
      SUCCESS,
      FAILURE
    }

    struct StaticConfig {
        uint64 chainSelector; // ───────╮ Destination chainSelector
        uint16 gasForCallExactCheck; // | Gas for call exact check
        address rmnRemote; // ───────╯ RMN Verification Contract
        address tokenAdminRegistry; // Token admin registry address
        address nonceManager; // Nonce manager address
    }

    struct DynamicConfig {
        address feeQuoter; // ──────────────────────────────╮ FeeQuoter address on the local chain.
        uint32 permissionLessExecutionThresholdSeconds; // ─╯ Waiting time before manual execution is enabled.
        address messageInterceptor; // Optional, validates incoming messages (zero address = no interceptor).
    }

    struct SourceChainConfig {
        address router; // ─────────────────╮ Local router to use for messages coming from this source chain.
        bool isEnabled; //                  │ Flag whether the source chain is enabled or not.
        uint64 minSeqNr; //                 │ The min sequence number expected for future messages.
        bool isRMNVerificationDisabled; // ─╯ Flag whether the RMN verification is disabled or not.
        bytes onRamp; // OnRamp address on the source chain.
    }

    struct MerkleRoot {
        uint64 sourceChainSelector; // Remote source chain selector that the Merkle Root is scoped to
        bytes onRampAddress; //        Generic onRamp address, to support arbitrary sources; for EVM, use abi.encode
        uint64 minSeqNr; // ─────────╮ Minimum sequence number, inclusive
        uint64 maxSeqNr; // ─────────╯ Maximum sequence number, inclusive
        bytes32 merkleRoot; //         Merkle root covering the interval & source chain messages
    }

    FiredrillCompound private immutable i_ctrl;

    constructor(FiredrillCompound ctrl) Ownable(msg.sender) {
        i_ctrl = ctrl;
    }

    function emitSourceChainConfigSet() public {
        require(msg.sender == address(i_ctrl), "only control");
        emit SourceChainConfigSet({
            sourceChainSelector: i_ctrl.chainSelector(),
            sourceConfig: SourceChainConfig({
                router: address(i_ctrl),
                isEnabled: true,
                minSeqNr: 0,
                isRMNVerificationDisabled: false,
                onRamp: abi.encodePacked(i_ctrl.onRamp())
            })
        });
    }    
    
    function getStaticConfig() external view returns (StaticConfig memory) {
        return StaticConfig({
            chainSelector: i_ctrl.chainSelector(),
            gasForCallExactCheck: 0,
            rmnRemote: address(i_ctrl),
            tokenAdminRegistry: address(i_ctrl),
            nonceManager: address(i_ctrl)
        });
    }

    function getDynamicConfig() external view returns (DynamicConfig memory) {
        return DynamicConfig({
            feeQuoter: address(i_ctrl),
            permissionLessExecutionThresholdSeconds: 10,
            messageInterceptor: address(0)
        });
    }

    function getSourceChainConfig(uint64) external view returns (SourceChainConfig memory) {
        return SourceChainConfig({
            router: address(i_ctrl),
            isEnabled: true,
            minSeqNr: 0,
            isRMNVerificationDisabled: false,
            onRamp: abi.encodePacked(i_ctrl.onRamp())
        });
    }

    function emitCommitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) external {
        Internal.PriceUpdates memory priceUpdates;
        MerkleRoot[] memory blessedMerkleRoots = new MerkleRoot[](1);
        blessedMerkleRoots[0] = MerkleRoot({
            sourceChainSelector: i_ctrl.chainSelector(),
            onRampAddress: abi.encode(i_ctrl.onRamp()),
            minSeqNr: minSeqNr,
            maxSeqNr: maxSeqNr,
            merkleRoot: keccak256(abi.encode(i_ctrl.onRamp(), minSeqNr, maxSeqNr))
        });
        emit CommitReportAccepted({
            blessedMerkleRoots: blessedMerkleRoots,
            unblessedMerkleRoots: new MerkleRoot[](0),
            priceUpdates: priceUpdates
        });
    }

    function typeAndVersion() external pure returns (string memory) {
        return "OffRamp 1.6.0";
    }
}
