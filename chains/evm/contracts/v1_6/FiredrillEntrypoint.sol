// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FiredrillToken} from  "../v1_0/FiredrillToken.sol";
import {FiredrillRevertMessageReceiver} from "../v1_0/FiredrillRevertMessageReceiver.sol";
import {FiredrillOnRamp} from "./FiredrillOnRamp.sol";
import {FiredrillOffRamp} from "./FiredrillOffRamp.sol";
import {FiredrillCompound} from "./FiredrillCompound.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";

contract FiredrillEntrypoint is FiredrillCompound {
    uint64 private s_send_last;

    constructor(uint64 chainSelector) FiredrillCompound(chainSelector) {}

    function prepare_Register() public onlyOwner {
        i_offRamp.emitSourceChainConfigSet(); // register OffRamp
    }

    function drill_PendingCommit_PendingQueue_TxSpike(uint8 from, uint8 to) public onlyOwner {
        require(from <= to, "nothing to send");
        require(from > s_send_last, "message already sent");
        for (uint64 i = from; i <= to; i++) {
            i_onRamp.emitCCIPMessageSent(msg.sender, i);
        }
        s_send_last = to;
    }

    function drill_PendingExecution(uint8 from, uint8 to) public onlyOwner {
        require(from <= to, "nothing to send");
        require(to <= s_send_last, "not yet sent");
        i_offRamp.emitCommitReportAccepted(from, to);
    }

    function drill_PriceRegistries() public onlyOwner {
        emitUsdPerTokenUpdated();
    }

    function typeAndVersion() external pure returns (string memory) {
        return "RouterFeeQuoter 1.6.0";
    }
}