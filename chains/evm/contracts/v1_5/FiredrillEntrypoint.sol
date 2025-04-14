// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {HasStatus} from "../common/HasStatus.sol";
import {FiredrillToken} from  "../v1_0/FiredrillToken.sol";
import {FiredrillRevertMessageReceiver} from "../v1_0/FiredrillRevertMessageReceiver.sol";
import {FiredrillOnRamp} from "./FiredrillOnRamp.sol";
import {FiredrillOffRamp} from "./FiredrillOffRamp.sol";
import {FiredrillCompound} from "./FiredrillCompound.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";

contract FiredrillEntrypoint is Ownable2Step, HasStatus, ITypeAndVersion {
    uint64 internal immutable i_chainSelector;
    FiredrillToken internal immutable i_token;
    FiredrillOnRamp internal immutable i_onRamp;
    FiredrillOffRamp internal immutable i_offRamp;
    FiredrillCompound internal immutable i_compound;
    FiredrillRevertMessageReceiver internal immutable i_receiver;

    // s_active flag signalizes if firedrill process is in active state
    // deactivation will be reflected in rdd-monster and alerts will stop triggering
    bool private s_active;

    uint64 private s_send_last;

    constructor(uint64 chainSelector) Ownable(msg.sender) {
        i_chainSelector = chainSelector;
        i_token = new FiredrillToken(this);
        i_onRamp = new FiredrillOnRamp(this);
        i_offRamp = new FiredrillOffRamp(this);
        i_compound = new FiredrillCompound(this);
        i_receiver = new FiredrillRevertMessageReceiver(this);
        s_active = true;
    }

    function deactivate() public onlyOwner {
        s_active = false;
    }

    function isActive() external view returns (bool) {
        return s_active;
    }

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
        i_compound.emitTagRootBlessed(from, to);
    }

    function drill_InvalidMessageState(uint8 from, uint8 to) public onlyOwner {
        require(from <= to, "nothing to send");
        require(to <= s_send_last, "not yet sent");
        for (uint64 i = from; i <= to; i++) {
            i_offRamp.emitExecutionStateChanged(msg.sender, i);
        }
    }

    function drill_PriceRegistries() public onlyOwner {
        i_compound.emitUsdPerTokenUpdated();
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
        i_compound.emitReleased();
    }

    function drill_RmnCurse() public onlyOwner {
        i_compound.emitCursed();
    }

    function drill_RmnVotedToCurse() public onlyOwner {
        i_compound.emitVotedToCurse();
    }

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

    function compound() public view returns (address) {
        return address(i_compound);
    }

    function receiver() public view returns (address) {
        return address(i_receiver);
    }

    function typeAndVersion() external pure returns (string memory) {
        return "FiredrillEntrypoint 1.6.0";
    }
}
