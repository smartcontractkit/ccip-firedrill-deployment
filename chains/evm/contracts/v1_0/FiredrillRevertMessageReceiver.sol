// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";
import {MaybeRevertMessageReceiver} from "@chainlink/ccip/test/helpers/receivers/MaybeRevertMessageReceiver.sol";

contract FiredrillRevertMessageReceiver is MaybeRevertMessageReceiver, Ownable2Step, ITypeAndVersion {

    constructor() MaybeRevertMessageReceiver(true) Ownable(msg.sender) {}

    function typeAndVersion() external pure returns (string memory){
        return "FiredrillRevertMessageReceiver 1.0.0";
    }
}
