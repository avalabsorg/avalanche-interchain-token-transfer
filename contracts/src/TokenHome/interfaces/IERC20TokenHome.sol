// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20TokenTransferer} from "../../interfaces/IERC20TokenTransferer.sol";
import {ITokenHome} from "./ITokenHome.sol";

/**
 * @notice Interface for a ERC20 token "home" contract that locks its specified ERC20
 * token on its chain to be transferred to supported remote token transfer contracts on other chains.
 *
 * @custom:security-contact https://github.com/ava-labs/avalanche-interchain-token-transfer/blob/main/SECURITY.md
 */
interface IERC20TokenHome is IERC20TokenTransferer, ITokenHome {
    /**
     * @notice Adds collateral to the home token transfer contract for the specified TokenRemote instance. If more value is provided
     * than the amount of collateral needed, the excess amount is returned to the caller.
     * @param remoteBlockchainID The blockchain ID of the TokenRemote instance to add collateral for.
     * @param remoteTokenTransferAddress The address of the TokenRemote instance to add collateral for on the {remoteBlockchainID}.
     * @param amount Amount of tokens to add as collateral.
     */
    function addCollateral(
        bytes32 remoteBlockchainID,
        address remoteTokenTransferAddress,
        uint256 amount
    ) external;
}
