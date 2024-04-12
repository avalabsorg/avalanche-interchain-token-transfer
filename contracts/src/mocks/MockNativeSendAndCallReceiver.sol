// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {INativeSendAndCallReceiver} from "../interfaces/INativeSendAndCallReceiver.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice TODO
 */
contract MockNativeSendAndCallReceiver is INativeSendAndCallReceiver {
    /**
     * @dev Emitted when receiveTokens is called.
     */
    event TokensReceived(uint256 amount, bytes payload);

    /**
     * @dev See {INativeSendAndCallReceiver-receiveTokens}
     */
    function receiveTokens(bytes calldata payload) external payable {
        emit TokensReceived(msg.value, payload);

        require(payload.length != 0, "MockNativeSendAndCallReceiver: empty payload");
        // No implementation required to accept native tokens
    }
}
