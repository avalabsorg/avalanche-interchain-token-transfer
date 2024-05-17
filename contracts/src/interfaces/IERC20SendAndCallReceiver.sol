// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Interface for a contracts that are called to receive bridge tokens.
 */
interface IERC20SendAndCallReceiver {
    /**
     * @notice Called to receive the amount of the given token
     * @param sourceBlockchainID blockchain ID that the transfer originated from
     * @param originBridgeAddress address of the bridge that initiated the Teleporter message
     * @param originSenderAddress address of the sender that sent the transfer. This value
     * should only be trusted if {originBridgeAddress} is verified and known.
     * @param token address of the token to be received
     * @param amount amount of the token to be received
     * @param payload arbitrary data provided by the caller
     */
    function receiveTokens(
        bytes32 sourceBlockchainID,
        address originBridgeAddress,
        address originSenderAddress,
        address token,
        uint256 amount,
        bytes calldata payload
    ) external;
}
