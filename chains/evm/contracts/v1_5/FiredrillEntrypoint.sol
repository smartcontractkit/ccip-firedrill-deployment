// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FiredrillCompound} from "./FiredrillCompound.sol";

contract FiredrillEntrypoint is FiredrillCompound {
    uint64 private s_send_last;

    constructor(uint64 chainSelector) FiredrillCompound(chainSelector) {}

    function prepare_Register() public onlyOwner {
        i_offRamp.emitConfigSet(); // register OffRamp
    }

    function drill_PendingCommit_PendingQueue_TxSpike(uint8 from, uint8 to) public onlyOwner {
        require(from <= to, "nothing to send");
        require(from > s_send_last, "message already sent");
        for (uint64 i = from; i <= to; i++) {
            i_onRamp.emitCCIPSendRequested(msg.sender, i);
        }
        s_send_last = to;
    }

    function drill_PendingBless(uint8 from, uint8 to) public onlyOwner {
        require(from <= to, "nothing to send");
        require(to <= s_send_last, "not yet sent");
        i_offRamp.emitReportAccepted(from, to);
    }

    function drill_PendingExecution(uint8 from, uint8 to) public onlyOwner {
        require(from <= to, "nothing to send");
        require(to <= s_send_last, "not yet sent");
        emitTagRootBlessed(from, to);
    }

    function drill_InvalidMessageState(uint8 from, uint8 to) public onlyOwner {
        require(from <= to, "nothing to send");
        require(to <= s_send_last, "not yet sent");
        for (uint64 i = from; i <= to; i++) {
            i_offRamp.emitExecutionStateChanged(msg.sender, i);
        }
    }

    function drill_PriceRegistries() public onlyOwner {
        emitUsdPerTokenUpdated();
    }

    function drill_PrivilegedFunctionCalled() public onlyOwner {
        i_offRamp.emitConfigSet();
    }

    /**
     * onramp/offramp: [Mainnet] Less than 20% of aggregate ratelimit capacity left for prolonged time
     * onramp/offramp: [Mainnet] 20% of aggregate ratelimit capacity hit
     */
    function drill_TokenPoolRateLimit_Offramp() public onlyOwner {
        i_offRamp.emitTokenConsumed();
    }

    /**
    * tokenpool: [Mainnet] Less than 20% of tokenpool ratelimit capacity left for prolonged time
    * tokenpool: [Mainnet] 20% of tokenpool ratelimit capacity hit
    * tokenpool: [Mainnet] For a given token, total released&minted value greater than locked&burned value
    * tokenpool: [Mainnet] Token pool flux exceeded the rate limit capacity
    * tokenpool: [Mainnet] LOW URGENCY - Token Pool low liquidity
    * tokenpool: [Mainnet] HIGH URGENCY -  Token Pool low liquidity
    * tokenpool: [Mainnet] Token value transfer greater than 80% of capacity
    */
    function drill_TokenPoolRateLimit_Inbound() public onlyOwner {
        emitReleased();
    }

    function drill_RmnCurse() public onlyOwner {
        emitCursed();
    }

    function drill_RmnVotedToCurse() public onlyOwner {
        emitVotedToCurse();
    }

    function typeAndVersion() external pure override returns (string memory) {
        return "RouterPriceRegistry 1.5.0";
    }
}