// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenTransferrer} from "../../interfaces/ITokenTransferrer.sol";
import {TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";

/**
 * THIS CONTRACT HAS BEEN AUDITED, BUT MAY CONTAIN UN-AUDITED CHANGES.
 * PLEASE CHECK `audits/README.md` FOR INFORMATION ON AUDITED VERSIONS OF THIS CONTRACT.
 * DO NOT USE UN-AUDITED CODE IN PRODUCTION.
 */

/**
 * @notice Settings for constructing a {ITokenRemote} contract.
 * @param teleporterRegistryAddress The current blockchain ID's Teleporter registry
 * address. See here for details: https://github.com/ava-labs/teleporter/tree/main/contracts/src/Teleporter/upgrades.
 * @param teleporterManager Address that manages this contract's integration with the
 * Teleporter registry and Teleporter versions.
 * @param tokenHomeBlockchainID The blockchain ID of the associated TokenHome instance.
 * @param tokenHomeAddress The address of the associated token home contract.
 * @param tokenHomeDecimals The number of decimal places used by the token home's token.
 */
struct TokenRemoteSettings {
    address teleporterRegistryAddress;
    address teleporterManager;
    bytes32 tokenHomeBlockchainID;
    address tokenHomeAddress;
    uint8 tokenHomeDecimals;
}

/**
 * @dev Interface for a remote token transfer contract that mints a representation token on its chain, and allows
 * for burning that token to redeem the backing asset on the home chain, or transferring to other remotes.
 */
interface ITokenRemote is ITokenTransferrer {
    /**
     * @notice Sends a Teleporter message to register the TokenRemote instance with its configured home.
     * @param feeInfo The optional fee asset and amount for the Teleporter message, for relayer incentivization.
     */
    function registerWithHome(TeleporterFeeInfo calldata feeInfo) external;
}
