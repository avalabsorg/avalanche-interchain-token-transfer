// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20Upgradeable} from
    "@openzeppelin/contracts-upgradeable@4.9.6/token/ERC20/ERC20Upgradeable.sol";

/**
 * @title IWrappedNativeToken
 * @notice Interface for a wrapped native token
 * @dev Implements the {IERC20} interface, and adds deposit and withdraw functions.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
interface IWrappedNativeToken is IERC20Upgradeable {
    /**
     * @notice Emitted when native tokens are deposited.
     * @param sender Address that deposited the native tokens
     * @param amount Amount of native tokens deposited
     */
    event Deposit(address indexed sender, uint256 amount);

    /**
     * @notice Emitted when wrapped tokens are withdrawn for native tokens.
     * @param sender Address that withdrew the native tokens.
     * @param amount Amount of native tokens withdrawn
     */
    event Withdrawal(address indexed sender, uint256 amount);

    /**
     * @notice Deposits native tokens to receive wrapped tokens.
     */
    function deposit() external payable;

    /**
     * @notice Withdraws native tokens for wrapped tokens.
     * @param amount Amount of native tokens to withdraw
     */
    function withdraw(uint256 amount) external;
}
