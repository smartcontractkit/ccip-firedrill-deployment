// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {HasStatus} from "../common/HasStatus.sol";
import {FiredrillToken} from "../v1_0/FiredrillToken.sol";
import {FiredrillRevertMessageReceiver} from "../v1_0/FiredrillRevertMessageReceiver.sol";
import {FiredrillOnRamp} from "./FiredrillOnRamp.sol";
import {FiredrillOffRamp} from "./FiredrillOffRamp.sol";
import {FiredrillEntrypoint} from "./FiredrillEntrypoint.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";

contract FiredrillCompound is Ownable2Step, HasStatus, ITypeAndVersion {
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

    FiredrillEntrypoint private immutable i_ctrl;

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

    // ========== Router ==========
    function getOnRamp(uint64) external view returns (address) {
        return address(i_ctrl.onRamp());
    }

    function getOffRamps() external view returns (OffRamp[] memory) {
        OffRamp[] memory offRamps = new OffRamp[](1);
        offRamps[0] = OffRamp({
            sourceChainSelector: i_ctrl.chainSelector(),
            offRamp: address(i_ctrl.offRamp())
        });
        return offRamps;
    }
    // ==================================

    // ========== FeeQuoter ==========
    function getStaticConfig() external view returns (StaticConfig memory) {
        return StaticConfig({
            maxFeeJuelsPerMsg: 1,
            linkToken: i_ctrl.token(),
            tokenPriceStalenessThreshold: 0
        });
    }
    // ==================================

    // ========== ARMProxy ==========
    function getARM() public view returns (address) {
        return i_ctrl.compound();
    }
    // ==================================

    function typeAndVersion() external pure returns (string memory) {
        return "Router 1.2.0 FeeQuoter 1.6.0";
    }
}
