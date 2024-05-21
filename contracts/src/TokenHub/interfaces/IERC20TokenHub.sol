// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20Bridge} from "./IERC20Bridge.sol";
import {ITokenHub} from "./ITokenHub.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Interface for a "hub" ERC20 bridge contract that locks its specified ERC20
 * token on its chain to be bridged to supported spoke bridge contracts on other chains.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
interface IERC20TokenHub is IERC20Bridge, ITokenHub {
    /**
     * @notice Adds collateral to the hub bridge contract for the specified spoke instance. If more value is provided
     * than the amount of collateral needed, the excess amount is returned to the caller.
     * @param spokeBlockchainID The blockchain ID of the spoke instance to add collateral for.
     * @param spokeBridgeAddress The address of the spoke instance to add collateral for on the {spokeBlockchainID}.
     * @param amount Amount of tokens to add as collateral.
     */
    function addCollateral(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount
    ) external;
}
