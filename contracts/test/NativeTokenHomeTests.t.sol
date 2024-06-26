// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TokenHomeTest} from "./TokenHomeTests.t.sol";
import {NativeTokenBridgeTest} from "./NativeTokenBridgeTests.t.sol";
import {NativeTokenHome} from "../src/TokenHome/NativeTokenHome.sol";
import {IWrappedNativeToken} from "../src/interfaces/IWrappedNativeToken.sol";
import {INativeSendAndCallReceiver} from "../src/interfaces/INativeSendAndCallReceiver.sol";
import {WrappedNativeToken} from "../src/WrappedNativeToken.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

contract NativeTokenHomeTest is NativeTokenBridgeTest, TokenHomeTest {
    using SafeERC20 for IERC20;

    NativeTokenHome public app;
    IWrappedNativeToken public mockWrappedToken;

    receive() external payable {
        require(msg.sender == address(app), "NativeTokenHomeTest: invalid receive payable sender");
    }

    function setUp() public override {
        TokenHomeTest.setUp();

        WrappedNativeToken token = new WrappedNativeToken();
        token.initialize("AVAX");
        mockWrappedToken = token;
        app = new NativeTokenHome();
        app.initialize(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockWrappedToken)
        );
        tokenHome = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        bridgedToken = IERC20(address(mockWrappedToken));
        tokenHomeDecimals = 18;
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        invalidInitialization(
            address(0),
            address(this),
            address(mockWrappedToken),
            "TeleporterUpgradeable: zero teleporter registry address"
        );
    }

    function testZeroTeleporterManagerAddress() public {
        invalidInitialization(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            address(mockWrappedToken),
            "Ownable: new owner is the zero address"
        );
    }

    function testZeroFeeTokenAddress() public {
        invalidInitialization(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(0),
            _formatErrorMessage("zero token address")
        );
    }

    function invalidInitialization(
        address teleporterRegistryAddress,
        address teleporterManagerAddress,
        address wrappedTokenAddress,
        bytes memory expectedErrorMessage
    ) private {
        app = new NativeTokenHome();
        vm.expectRevert(expectedErrorMessage);
        app.initialize(teleporterRegistryAddress, teleporterManagerAddress, wrappedTokenAddress);
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenHome));
        emit TokensWithdrawn(recipient, amount);
        vm.expectCall(
            address(mockWrappedToken), abi.encodeCall(IWrappedNativeToken.withdraw, (amount))
        );
        vm.expectEmit(true, true, true, true, address(mockWrappedToken));
        emit Withdrawal(address(app), amount);
    }

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal override {
        if (targetHasCode) {
            // Non-zero code length
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                INativeSendAndCallReceiver.receiveTokens,
                (sourceBlockchainID, originInfo.bridgeAddress, originInfo.senderAddress, payload)
            );
            if (expectSuccess) {
                vm.mockCall(recipient, amount, expectedCalldata, new bytes(0));
            } else {
                vm.mockCallRevert(recipient, amount, expectedCalldata, new bytes(0));
            }
            vm.expectCall(recipient, amount, uint64(gasLimit), expectedCalldata);
        } else {
            // No code at target
            vm.etch(recipient, new bytes(0));
        }

        if (targetHasCode && expectSuccess) {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        } else {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        }
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal override {
        mockWrappedToken.deposit{value: feeAmount}();
        // Transfer the fee to the bridge if it is greater than 0
        if (feeAmount > 0) {
            bridgedToken.safeIncreaseAllowance(address(tokenBridge), feeAmount);
            vm.expectCall(
                address(bridgedToken),
                abi.encodeCall(
                    IERC20.transferFrom, (address(this), address(tokenBridge), feeAmount)
                )
            );
        }

        vm.expectCall(address(bridgedToken), abi.encodeCall(IWrappedNativeToken.deposit, ()));
        vm.expectEmit(true, true, true, true, address(bridgedToken));
        emit Deposit(address(nativeTokenBridge), amount);
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert("SafeWrappedNativeTokenDeposit: balance not increased");
    }

    function _addCollateral(
        bytes32 remoteBlockchainID,
        address remoteBridgeAddress,
        uint256 amount
    ) internal override {
        app.addCollateral{value: amount}(remoteBlockchainID, remoteBridgeAddress);
    }
}
