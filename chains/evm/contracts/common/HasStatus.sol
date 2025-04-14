// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

interface HasStatus {
    function isActive() external view returns (bool);
}
