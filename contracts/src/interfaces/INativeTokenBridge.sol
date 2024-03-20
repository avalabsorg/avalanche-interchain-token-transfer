// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterTokenBridge, SendTokensInput} from "./ITeleporterTokenBridge.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Interface for a Teleporter token bridge that sends native tokens to another chain.
 */
interface INativeTokenBridge is ITeleporterTokenBridge {
    /**
     * @notice Receives native tokens transferred to this contract.
     * @dev This function is called when the token bridge is withdrawing native tokens to
     * transfer to the recipient. The caller must be the wrapped native token contract.
     */
    receive() external payable;

    /**
     * @notice Sends native tokens transferred to this contract to the destination token bridge instance.
     * @param input specifies information for delivery of the tokens
     */
    function send(SendTokensInput calldata input) external payable;
}
