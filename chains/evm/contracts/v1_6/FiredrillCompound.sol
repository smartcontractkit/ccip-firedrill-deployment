// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";
import {FiredrillToken} from "../v1_0/FiredrillToken.sol";
import {FiredrillRevertMessageReceiver} from "../v1_0/FiredrillRevertMessageReceiver.sol";
import {FiredrillOnRamp} from "./FiredrillOnRamp.sol";
import {FiredrillOffRamp} from "./FiredrillOffRamp.sol";

abstract contract FiredrillCompound is Ownable2Step, ITypeAndVersion {
    event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp);

    struct OffRamp {
        uint64 sourceChainSelector;
        address offRamp;
    }

    // FeeQuoter static config
    struct StaticConfig {
        uint96 maxFeeJuelsPerMsg; // ─╮ Maximum fee that can be charged for a message.
        address linkToken; // ────────╯ LINK token address.
        uint32 tokenPriceStalenessThreshold;
    }

    uint64 internal immutable i_chainSelector;
    FiredrillToken internal immutable i_token;
    FiredrillOnRamp internal immutable i_onRamp;
    FiredrillOffRamp internal immutable i_offRamp;
    FiredrillRevertMessageReceiver internal immutable i_receiver;

    constructor(uint64 chainSelector_) Ownable(msg.sender) {
        i_chainSelector = chainSelector_;
        i_token = new FiredrillToken();
        i_onRamp = new FiredrillOnRamp(this);
        i_offRamp = new FiredrillOffRamp(this);
        i_receiver = new FiredrillRevertMessageReceiver();
    }

    function emitUsdPerTokenUpdated() public {
        emit UsdPerTokenUpdated({
            token: address(i_token),
            value: 1,
            timestamp: block.timestamp
        });
    }

    // ========== Router ==========
    function getOnRamp(uint64) external view returns (address) {
        return address(i_onRamp);
    }

    function getOffRamps() external view returns (OffRamp[] memory) {
        OffRamp[] memory offRamps = new OffRamp[](1);
        offRamps[0] = OffRamp({
            sourceChainSelector: i_chainSelector,
            offRamp: address(i_offRamp)
        });
        return offRamps;
    }
    // ==================================

    // ========== FeeQuoter ==========
    function getStaticConfig() external view returns (StaticConfig memory) {
        return StaticConfig({
            maxFeeJuelsPerMsg: 1,
            linkToken: address(i_token),
            tokenPriceStalenessThreshold: 0
        });
    }
    // ==================================

    // ========== ARMProxy ==========
    function getARM() public view returns (address) {
        return address(this);
    }
    // ==================================

    function chainSelector() public view returns (uint64) {
        return i_chainSelector;
    }

    function token() public view returns (address) {
        return address(i_token);
    }

    function onRamp() public view returns (address) {
        return address(i_onRamp);
    }

    function offRamp() public view returns (address) {
        return address(i_offRamp);
    }

    function receiver() public view returns (address) {
        return address(i_receiver);
    }
}
