// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {HasStatus} from "../common/HasStatus.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";
import {MaybeRevertMessageReceiver} from "@chainlink/ccip/test/helpers/receivers/MaybeRevertMessageReceiver.sol";

contract FiredrillRevertMessageReceiver is MaybeRevertMessageReceiver, HasStatus, Ownable2Step, ITypeAndVersion {
    HasStatus private i_ctrl;

    constructor(HasStatus ctrl) MaybeRevertMessageReceiver(true) Ownable(msg.sender) {
        i_ctrl = ctrl;
    }

    function isActive() public view returns (bool) {
        return i_ctrl.isActive();
    }

    function typeAndVersion() external pure returns (string memory){
        return "FiredrillRevertMessageReceiver 1.0.0";
    }
}
