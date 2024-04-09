// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenSource} from "./TeleporterTokenSource.sol";
import {INativeTokenBridge} from "./interfaces/INativeTokenBridge.sol";
import {INativeSendAndCallReceiver} from "./interfaces/INativeSendAndCallReceiver.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    SingleHopCallMessage
} from "./interfaces/ITeleporterTokenBridge.sol";
import {IWrappedNativeToken} from "./interfaces/IWrappedNativeToken.sol";
import {GasUtils} from "./utils/GasUtils.sol";
import {SafeWrappedNativeTokenDeposit} from "./SafeWrappedNativeTokenDeposit.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title NativeTokenSource
 * @notice This contract is an {INativeTokenBridge} that sends native tokens to another chain's
 * {ITeleporterTokenBridge} instance, and gets represented by the tokens of that destination
 * token bridge instance.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
contract NativeTokenSource is INativeTokenBridge, TeleporterTokenSource {
    using GasUtils for *;
    using SafeWrappedNativeTokenDeposit for IWrappedNativeToken;

    /**
     * @notice The wrapped native token contract that represents the native tokens on this chain.
     */
    IWrappedNativeToken public immutable token;

    /**
     * @notice Initializes this source token bridge instance
     * @dev Teleporter fees are paid by a {IWrappedNativeToken} instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address feeTokenAddress
    ) TeleporterTokenSource(teleporterRegistryAddress, teleporterManager, feeTokenAddress) {
        token = IWrappedNativeToken(feeTokenAddress);
    }

    /**
     * @notice Receives native tokens transferred to this contract.
     * @dev This function is called when the token bridge is withdrawing native tokens to
     * transfer to the recipient. The caller must be the wrapped native token contract.
     */
    receive() external payable {
        require(msg.sender == feeTokenAddress, "NativeTokenSource: invalid receive payable sender");
    }

    /**
     * @dev See {INativeTokenBridge-send}
     */
    function send(SendTokensInput calldata input) external payable nonReentrant {
        _send(input, msg.value, false);
    }

    function sendAndCall(SendAndCallInput calldata input) external payable {
        _sendAndCall(input, msg.value, false);
    }

    /**
     * @dev See {TeleportTokenSource-_deposit}
     * Deposits the native tokens sent to this contract
     */
    function _deposit(uint256 amount) internal virtual override returns (uint256) {
        return token.safeDeposit(amount);
    }

    /**
     * @dev See {TeleportTokenSource-_withdraw}
     * Withdraws the wrapped tokens for native tokens,
     * and sends them to the recipient.
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        token.withdraw(amount);
        payable(recipient).transfer(amount);
    }

    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual override {
        // Withdraw the native token from the wrapped native token contract.
        token.withdraw(amount);

        // Encode the call to {INativeSendAndCallReceiver-receiveTokens}
        bytes memory payload =
            abi.encodeCall(INativeSendAndCallReceiver.receiveTokens, (message.recipientPayload));

        // Call the destination contract with the given payload, gas amount, and value.
        bool success = GasUtils._callWithExactGasAndValue(
            message.recipientGasLimit, amount, message.recipientContract, payload
        );

        // If the call failed, send the funds to the fallback recipient.
        if (!success) {
            payable(message.fallbackRecipient).transfer(amount);
        }
    }
}
