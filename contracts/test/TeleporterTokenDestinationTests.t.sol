// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {TeleporterTokenDestination, IWarpMessenger} from "../src/TeleporterTokenDestination.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITeleporterTokenBridge.sol";

abstract contract TeleporterTokenDestinationTest is TeleporterTokenBridgeTest {
    TeleporterTokenDestination public tokenDestination;

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_DESTINATION_BLOCKCHAIN_ID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion.selector
            )
        );
    }

    function testInvalidSendingBackToSourceBlockchain() public {
        uint256 amount = 3;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(this);
        _setUpExpectedDeposit(amount);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, amount);
    }

    function testSendingToSameInstance() public {
        uint256 amount = 3;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenDestination.blockchainID();
        input.destinationBridgeAddress = address(tokenDestination);
        _setUpExpectedDeposit(amount);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, amount);
    }

    function testSendToSameBlockchainDifferentDestination() public {
        uint256 amount = 2;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenDestination.blockchainID();
        // Set the desintation bridge address to an address different than the token destination contract.
        input.destinationBridgeAddress = address(0x55);

        _sendSuccess(amount, input.primaryFee);
    }

    function testSendAndCallInvalidContractRecipient() public {
        vm.expectRevert(_formatErrorMessage("zero recipient contract address"));
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.recipientContract = address(0);
        _sendAndCall(input, 1);
    }

    function testReceiveInvalidSourceChain() public {
        vm.expectRevert(_formatErrorMessage("invalid source chain"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, new bytes(0)
        );
    }

    function testReceiveInvalidTokenSourceAddress() public {
        vm.expectRevert(_formatErrorMessage("invalid token source address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, address(0), new bytes(0)
        );
    }

    function testReceiveInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, bytes("test")
        );
    }

    function testReceiveWithdrawSuccess() public {
        uint256 amount = 2;
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, amount);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testInvalidMessageType() public {
        vm.expectRevert(_formatErrorMessage("invalid message type"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeMultiHopSendMessage(
                1,
                DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                DEFAULT_DESTINATION_ADDRESS,
                DEFAULT_RECIPIENT_ADDRESS,
                0
            )
        );
    }

    function _requiredGasLimit() internal view virtual override returns (uint256) {
        return tokenDestination.SEND_TOKENS_REQUIRED_GAS();
    }

    function _createDefaultSendTokensInput()
        internal
        pure
        override
        returns (SendTokensInput memory)
    {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: TOKEN_SOURCE_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
    }

    function _createDefaultSendAndCallInput()
        internal
        pure
        override
        returns (SendAndCallInput memory)
    {
        return SendAndCallInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: TOKEN_SOURCE_ADDRESS,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: new bytes(16),
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
    }

    function _formatErrorMessage(string memory message)
        internal
        pure
        override
        returns (bytes memory)
    {
        return bytes(string.concat("TeleporterTokenDestination: ", message));
    }
}
