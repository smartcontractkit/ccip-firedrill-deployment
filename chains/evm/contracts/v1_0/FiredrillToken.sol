// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {Ownable2Step} from "@openzeppelin/access/Ownable2Step.sol";
import {ITypeAndVersion} from "@chainlink/shared/interfaces/ITypeAndVersion.sol";

contract FiredrillToken is Ownable2Step, ITypeAndVersion {
    constructor() Ownable(msg.sender) {}

    function balanceOf(address) public pure returns (uint256) {
        return 0;
    }

    function decimals() public pure returns (uint8) {
        return 0;
    }

    function symbol() public pure returns (string memory) {
        return "LINK";
    }

    function typeAndVersion() external pure returns (string memory){
        return "FiredrillToken 1.0.0";
    }
}
