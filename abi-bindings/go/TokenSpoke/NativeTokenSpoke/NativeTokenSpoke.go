// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenspoke

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	RecipientContract        common.Address
	RecipientPayload         []byte
	RequiredGasLimit         *big.Int
	RecipientGasLimit        *big.Int
	MultiHopFallback         common.Address
	FallbackRecipient        common.Address
	PrimaryFeeTokenAddress   common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	PrimaryFeeTokenAddress   common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
	RequiredGasLimit         *big.Int
	MultiHopFallback         common.Address
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TokenSpokeSettings is an auto generated low-level Go binding around an user-defined struct.
type TokenSpokeSettings struct {
	TeleporterRegistryAddress common.Address
	TeleporterManager         common.Address
	TokenHubBlockchainID      [32]byte
	TokenHubAddress           common.Address
	TokenHubDecimals          uint8
}

// NativeTokenSpokeMetaData contains all meta data concerning the NativeTokenSpoke contract.
var NativeTokenSpokeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenHubBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenHubAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenHubDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenSpokeSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HUB_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_SPOKE_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnedFeesReportingRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastestBurnedFeesReported\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnSpoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenHubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenHubBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101c06040523480156200001257600080fd5b50604051620057d5380380620057d583398101604081905262000035916200080d565b83826012826000015183602001518188604051602001620000579190620008d9565b60408051601f198184030181529190528960036200007683826200099a565b5060046200008582826200099a565b50506001600555506001600160a01b0381166200010f5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200015a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000180919062000a66565b600755506200018f3362000578565b6200019a81620005ca565b505060016009819055507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001f7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200021d919062000a66565b60a0526040830151620002845760405162461bcd60e51b815260206004820152602860248201527f546f6b656e53706f6b653a207a65726f20746f6b656e2068756220626c6f636b60448201526718da185a5b88125160c21b606482015260840162000106565b60a051836040015103620003015760405162461bcd60e51b815260206004820152603960248201527f546f6b656e53706f6b653a2063616e6e6f74206465706c6f7920746f2073616d60448201527f6520626c6f636b636861696e20617320746f6b656e2068756200000000000000606482015260840162000106565b60608301516001600160a01b0316620003685760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a207a65726f20746f6b656e20687562206164647265604482015261737360f01b606482015260840162000106565b6012836080015160ff161115620003d25760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a20746f6b656e2068756220646563696d616c7320746044820152660dede40d0d2ced60cb1b606482015260840162000106565b60128160ff161115620004345760405162461bcd60e51b815260206004820152602360248201527f546f6b656e53706f6b653a20746f6b656e20646563696d616c7320746f6f20686044820152620d2ced60eb1b606482015260840162000106565b604083015160c05260608301516001600160a01b031660e052610180829052600a8054831560ff19909116179055608083015160ff90811661010081905290821661012081905262000493919062000649602090811b6200165117901c565b1515610160526101405250505060008290036200050c5760405162461bcd60e51b815260206004820152603060248201527f4e6174697665546f6b656e53706f6b653a207a65726f20696e697469616c207260448201526f65736572766520696d62616c616e636560801b606482015260840162000106565b606481106200056a5760405162461bcd60e51b8152602060048201526024808201527f4e6174697665546f6b656e53706f6b653a20696e76616c69642070657263656e6044820152637461676560e01b606482015260840162000106565b6101a0525062000bca915050565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b620005d46200069a565b6001600160a01b0381166200063b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840162000106565b620006468162000578565b50565b60008060ff808516908416118181620006715762000668858762000a96565b60ff1662000681565b6200067d868662000a96565b60ff165b6200068e90600a62000bb5565b96919550909350505050565b6008546001600160a01b03163314620006f65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000106565b565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620007335762000733620006f8565b60405290565b80516001600160a01b03811681146200075157600080fd5b919050565b60005b838110156200077357818101518382015260200162000759565b50506000910152565b600082601f8301126200078e57600080fd5b81516001600160401b0380821115620007ab57620007ab620006f8565b604051601f8301601f19908116603f01168101908282118183101715620007d657620007d6620006f8565b81604052838152866020858801011115620007f057600080fd5b6200080384602083016020890162000756565b9695505050505050565b6000806000808486036101008112156200082657600080fd5b60a08112156200083557600080fd5b50620008406200070e565b6200084b8662000739565b81526200085b6020870162000739565b602082015260408601516040820152620008786060870162000739565b6060820152608086015160ff811681146200089257600080fd5b608082015260a08601519094506001600160401b03811115620008b457600080fd5b620008c2878288016200077c565b60c087015160e09097015195989097509350505050565b6702bb930b83832b2160c51b815260008251620008fe81600885016020870162000756565b9190910160080192915050565b600181811c908216806200092057607f821691505b6020821081036200094157634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200099557600081815260208120601f850160051c81016020861015620009705750805b601f850160051c820191505b8181101562000991578281556001016200097c565b5050505b505050565b81516001600160401b03811115620009b657620009b6620006f8565b620009ce81620009c784546200090b565b8462000947565b602080601f83116001811462000a065760008415620009ed5750858301515b600019600386901b1c1916600185901b17855562000991565b600085815260208120601f198616915b8281101562000a375788860151825594840194600190910190840162000a16565b508582101562000a565787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006020828403121562000a7957600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60ff828116828216039081111562000ab25762000ab262000a80565b92915050565b600181815b8085111562000af957816000190482111562000add5762000add62000a80565b8085161562000aeb57918102915b93841c939080029062000abd565b509250929050565b60008262000b125750600162000ab2565b8162000b215750600062000ab2565b816001811462000b3a576002811462000b455762000b65565b600191505062000ab2565b60ff84111562000b595762000b5962000a80565b50506001821b62000ab2565b5060208310610133831016604e8410600b841016171562000b8a575081810a62000ab2565b62000b96838362000ab8565b806000190482111562000bad5762000bad62000a80565b029392505050565b600062000bc3838362000b01565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051614aa862000d2d600039600081816104fb0152610e2601526000818161068601528181610a8a015261124c01526000818161095901528181610eac015281816137fa015261384501526000818161078f01528181610e8b015281816137d9015261382401526000818161052f015261129c015260008181610925015261127101526000818161042a01528181610fc301528181611355015281816122ff01528181612acb01528181612d8a015281816130450152818161328101526136ab0152600081816108b101528181610f9d0152818161132f0152818161207a015281816121a00152818161227f01528181612aa501528181612d640152818161301f015261325b0152600081816108480152818161299501526138d901526000818161039f015281816113e401528181611eac01526124e80152614aa86000f3fe6080604052600436106102cd5760003560e01c806371717c1811610175578063c452165e116100dc578063dd62ed3e11610095578063f2fde38b1161006f578063f2fde38b146108d3578063f3f981d8146108f3578063fdd7546f14610913578063feea2b7214610947576102dc565b8063dd62ed3e1461087f578063eae0544f14610571578063eb3da9091461089f576102dc565b8063c452165e146107e8578063c868efaa14610800578063d0e30db0146102dc578063d10a5b8c14610820578063d127dc9b14610836578063d2cc7a701461086a576102dc565b8063a2309ff81161012e578063a2309ff814610727578063a457c2d71461073d578063a9059cbb1461075d578063ba3f5a121461077d578063bf48d36b146107b1578063c03f4013146107d1576102dc565b806371717c181461065d5780638ac7dd20146106745780638bf2fa94146106a85780638da5cb5b146106bb57806395d89b41146106d957806397314297146106ee576102dc565b8063329c3e121161023457806349e3284e116101ed5780635eb99514116101c75780635eb99514146105df5780636e6eef8d146105ff57806370a0823114610612578063715018a614610648576102dc565b806349e3284e1461058e5780635507f3d1146105a857806355538c8b146105bf576102dc565b8063329c3e12146104ae57806339509351146104c95780633a23dfe2146104e95780633b97e8561461051d5780634511243e1461055157806347a9a22c14610571576102dc565b8063223668441161028657806322366844146103d957806323b872dd146103f857806324427629146104185780632b0d8f181461044c5780632e1a7d4d1461046c578063313ce5671461048c576102dc565b806306fdde03146102e4578063095ea7b31461030f57806315beb59f1461033f57806318160ddd146103635780631906529c146103785780631a7f5bec1461038d576102dc565b366102dc576102da61097b565b005b6102da61097b565b3480156102f057600080fd5b506102f96109bc565b6040516103069190613c86565b60405180910390f35b34801561031b57600080fd5b5061032f61032a366004613cb9565b610a4e565b6040519015158152602001610306565b34801561034b57600080fd5b506103556105dc81565b604051908152602001610306565b34801561036f57600080fd5b50600254610355565b34801561038457600080fd5b50610355610a68565b34801561039957600080fd5b506103c17f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610306565b3480156103e557600080fd5b50600a5461032f90610100900460ff1681565b34801561040457600080fd5b5061032f610413366004613ce5565b610ac9565b34801561042457600080fd5b506103c17f000000000000000000000000000000000000000000000000000000000000000081565b34801561045857600080fd5b506102da610467366004613d26565b610aed565b34801561047857600080fd5b506102da610487366004613d43565b610bef565b34801561049857600080fd5b5060125b60405160ff9091168152602001610306565b3480156104ba57600080fd5b506103c16001600160991b0181565b3480156104d557600080fd5b5061032f6104e4366004613cb9565b610c3b565b3480156104f557600080fd5b506103557f000000000000000000000000000000000000000000000000000000000000000081565b34801561052957600080fd5b5061049c7f000000000000000000000000000000000000000000000000000000000000000081565b34801561055d57600080fd5b506102da61056c366004613d26565b610c5d565b34801561057d57600080fd5b506103c162010203600160981b0181565b34801561059a57600080fd5b50600a5461032f9060ff1681565b3480156105b457600080fd5b506103556205302081565b3480156105cb57600080fd5b506102da6105da366004613d43565b610d5a565b3480156105eb57600080fd5b506102da6105fa366004613d43565b6110d4565b6102da61060d366004613d5c565b6110e5565b34801561061e57600080fd5b5061035561062d366004613d26565b6001600160a01b031660009081526020819052604090205490565b34801561065457600080fd5b506102da611111565b34801561066957600080fd5b506103556205573081565b34801561068057600080fd5b506103557f000000000000000000000000000000000000000000000000000000000000000081565b6102da6106b6366004613d98565b611123565b3480156106c757600080fd5b506008546001600160a01b03166103c1565b3480156106e557600080fd5b506102f961114f565b3480156106fa57600080fd5b5061032f610709366004613d26565b6001600160a01b031660009081526006602052604090205460ff1690565b34801561073357600080fd5b50610355600b5481565b34801561074957600080fd5b5061032f610758366004613cb9565b61115e565b34801561076957600080fd5b5061032f610778366004613cb9565b6111d9565b34801561078957600080fd5b506103557f000000000000000000000000000000000000000000000000000000000000000081565b3480156107bd57600080fd5b506102da6107cc366004613db1565b6111e7565b3480156107dd57600080fd5b506103556201c13881565b3480156107f457600080fd5b506103c1600160981b81565b34801561080c57600080fd5b506102da61081b366004613dc3565b6113cf565b34801561082c57600080fd5b50610355600c5481565b34801561084257600080fd5b506103557f000000000000000000000000000000000000000000000000000000000000000081565b34801561087657600080fd5b50600754610355565b34801561088b57600080fd5b5061035561089a366004613e4c565b611599565b3480156108ab57600080fd5b506103557f000000000000000000000000000000000000000000000000000000000000000081565b3480156108df57600080fd5b506102da6108ee366004613d26565b6115c4565b3480156108ff57600080fd5b5061035561090e366004613d43565b61163a565b34801561091f57600080fd5b5061049c7f000000000000000000000000000000000000000000000000000000000000000081565b34801561095357600080fd5b5061032f7f000000000000000000000000000000000000000000000000000000000000000081565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26109ba333461169c565b565b6060600380546109cb90613e85565b80601f01602080910402602001604051908101604052809291908181526020018280546109f790613e85565b8015610a445780601f10610a1957610100808354040283529160200191610a44565b820191906000526020600020905b815481529060010190602001808311610a2757829003601f168201915b5050505050905090565b600033610a5c81858561175c565b60019150505b92915050565b600080610a8462010203600160981b0131600160981b31613ecf565b905060007f0000000000000000000000000000000000000000000000000000000000000000600b54610ab69190613ecf565b9050610ac28282613ee2565b9250505090565b600033610ad7858285611881565b610ae28585856118f5565b506001949350505050565b610af5611a99565b6001600160a01b038116610b245760405162461bcd60e51b8152600401610b1b90613ef5565b60405180910390fd5b6001600160a01b03811660009081526006602052604090205460ff1615610ba35760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610b1b565b6001600160a01b038116600081815260066020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610c2e3382611aa1565b610c383382611bd0565b50565b600033610a5c818585610c4e8383611599565b610c589190613ecf565b61175c565b610c65611a99565b6001600160a01b038116610c8b5760405162461bcd60e51b8152600401610b1b90613ef5565b6001600160a01b03811660009081526006602052604090205460ff16610d055760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610b1b565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600660205260409020805460ff19169055565b600160095414610d7c5760405162461bcd60e51b8152600401610b1b90613f43565b6002600955600c54600160981b31908111610e0b5760405162461bcd60e51b815260206004820152604360248201527f4e6174697665546f6b656e53706f6b653a206275726e2061646472657373206260448201527f616c616e6365206e6f742067726561746572207468616e206c617374207265706064820152621bdc9d60ea1b608482015260a401610b1b565b6000600c5482610e1b9190613ee2565b905060006064610e4b7f000000000000000000000000000000000000000000000000000000000000000084613f87565b610e559190613f9e565b90506000610e638284613ee2565b600c85905590508115610e8457610e7a3083611ce9565b610e84308361169c565b6000610ed17f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084611d6b565b11610f3a5760405162461bcd60e51b815260206004820152603360248201527f4e6174697665546f6b656e53706f6b653a207a65726f207363616c656420616d60448201527237bab73a103a37903932b837b93a10313ab93760691b6064820152608401610b1b565b604080518082018252600181528151808301835262010203600160981b0181526020808201859052925160009380840192610f7792909101613fd6565b6040516020818303038152906040528152509050600061108a6040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b03168152602001888152508152602001898152602001600067ffffffffffffffff81111561103357611033613ff6565b60405190808252806020026020018201604052801561105c578160200160208202803683370190505b50815260200184604051602001611073919061400c565b604051602081830303815290604052815250611d82565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5846040516110be91815260200190565b60405180910390a2505060016009555050505050565b6110dc611a99565b610c3881611ea8565b600a5460ff166111075760405162461bcd60e51b8152600401610b1b90614051565b610c388134612048565b6111196120c2565b6109ba600061211c565b600a5460ff166111455760405162461bcd60e51b8152600401610b1b90614051565b610c38813461216e565b6060600480546109cb90613e85565b6000338161116c8286611599565b9050838110156111cc5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610b1b565b610ae2828686840361175c565b600033610a5c8185856118f5565b600a54610100900460ff161561123f5760405162461bcd60e51b815260206004820152601e60248201527f546f6b656e53706f6b653a20616c7265616479207265676973746572656400006044820152606401610b1b565b60408051606080820183527f0000000000000000000000000000000000000000000000000000000000000000825260ff7f0000000000000000000000000000000000000000000000000000000000000000811660208085019182527f00000000000000000000000000000000000000000000000000000000000000008316858701908152865180880188526000808252885188518186015294518616858a015291519094168386015286518084039095018552608090920190955281850192909252919261131d9061131390860186613d26565b85602001356121da565b90506113c86040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808860000160208101906113a09190613d26565b6001600160a01b0316815260209081018690529082526201c138908201526040016000611033565b5050505050565b6113d7612224565b6007546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561144e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611472919061409f565b10156114d95760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610b1b565b6114e233610709565b156115485760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610b1b565b611589848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061227d92505050565b6115936001600555565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6115cc6120c2565b6001600160a01b0381166116315760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610b1b565b610c388161211c565b6000600561164983601f613ecf565b901c92915050565b60008060ff8085169084161181816116755761166d85876140b8565b60ff16611683565b61167f86866140b8565b60ff165b61168e90600a6141b5565b9350909150505b9250929050565b6001600160a01b0382166116f25760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610b1b565b80600260008282546117049190613ecf565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b5050565b6001600160a01b0383166117be5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610b1b565b6001600160a01b03821661181f5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610b1b565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b600061188d8484611599565b9050600019811461159357818110156118e85760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610b1b565b611593848484840361175c565b6001600160a01b0383166119595760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610b1b565b6001600160a01b0382166119bb5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610b1b565b6001600160a01b03831660009081526020819052604090205481811015611a335760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610b1b565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3611593565b6109ba6120c2565b6001600160a01b038216611b015760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610b1b565b6001600160a01b03821660009081526020819052604090205481811015611b755760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610b1b565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611874565b505050565b80471015611c205760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610b1b565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611c6d576040519150601f19603f3d011682016040523d82523d6000602084013e611c72565b606091505b5050905080611bcb5760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610b1b565b80600b6000828254611cfb9190613ecf565b90915550506040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba90604401600060405180830381600087803b158015611d4f57600080fd5b505af1158015611d63573d6000803e3d6000fd5b505050505050565b6000611d7a84848460006124b2565b949350505050565b600080611d8d6124e3565b60408401516020015190915015611e32576040830151516001600160a01b0316611e0f5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610b1b565b604083015160208101519051611e32916001600160a01b039091169083906125f7565b604051630624488560e41b81526001600160a01b03821690636244885090611e5e9086906004016141c1565b6020604051808303816000875af1158015611e7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea1919061409f565b9392505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2c919061409f565b60075490915081831115611f9c5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610b1b565b8083116120115760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610b1b565b6007839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b60016009541461206a5760405162461bcd60e51b8152600401610b1b90613f43565b6002600955612078826126dc565b7f00000000000000000000000000000000000000000000000000000000000000008235036120af576120aa8282612916565b6120b9565b6120aa8282612b8d565b50506001600955565b6008546001600160a01b031633146109ba5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610b1b565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600954146121905760405162461bcd60e51b8152600401610b1b90613f43565b600260095561219e82612e64565b7f00000000000000000000000000000000000000000000000000000000000000008235036121d0576120aa8282612f51565b6120b982826130f7565b6000816000036121ec57506000610a62565b306001600160a01b0384160361221957612207333084611881565b6122123330846118f5565b5080610a62565b611ea18333846132f4565b6002600554036122765760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610b1b565b6002600555565b7f000000000000000000000000000000000000000000000000000000000000000083146122fd5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e53706f6b653a20696e76616c696420736f7572636520626c6f636b60448201526718da185a5b88125160c21b6064820152608401610b1b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316146123905760405162461bcd60e51b815260206004820152602960248201527f546f6b656e53706f6b653a20696e76616c6964206f726967696e2073656e646560448201526872206164647265737360b81b6064820152608401610b1b565b6000818060200190518101906123a69190614345565b600a54909150610100900460ff1615806123c35750600a5460ff16155b156123d857600a805461ffff19166101011790555b6001815160048111156123ed576123ed613fc0565b03612426576000816020015180602001905181019061240c91906143d4565b90506124208160000151826020015161345d565b50611593565b60028151600481111561243b5761243b613fc0565b0361246a576000816020015180602001905181019061245a919061440e565b90506124208182608001516134aa565b60405162461bcd60e51b815260206004820181905260248201527f546f6b656e53706f6b653a20696e76616c6964206d65737361676520747970656044820152606401610b1b565b6000811515841515036124d0576124c98584613f87565b9050611d7a565b6124da8584613f9e565b95945050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612544573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061256891906144de565b905061258c816001600160a01b031660009081526006602052604090205460ff1690565b156125f25760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610b1b565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015612648573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061266c919061409f565b6126769190613ecf565b6040516001600160a01b03851660248201526044810182905290915061159390859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526135d7565b80356126fa5760405162461bcd60e51b8152600401610b1b906144fb565b600061270c6040830160208401613d26565b6001600160a01b0316036127325760405162461bcd60e51b8152600401610b1b90614545565b60006127446060830160408401613d26565b6001600160a01b0316036127ae5760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a207a65726f20726563697069656e7420636f6e747260448201526a616374206164647265737360a81b6064820152608401610b1b565b60008160800135116127d25760405162461bcd60e51b8152600401610b1b90614590565b60008160a00135116128325760405162461bcd60e51b8152602060048201526024808201527f546f6b656e53706f6b653a207a65726f20726563697069656e7420676173206c6044820152631a5b5a5d60e21b6064820152608401610b1b565b80608001358160a00135106128995760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a20696e76616c696420726563697069656e742067616044820152661cc81b1a5b5a5d60ca1b6064820152608401610b1b565b60006128ac610100830160e08401613d26565b6001600160a01b031603610c385760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a207a65726f2066616c6c6261636b2072656369706960448201526a656e74206164647265737360a81b6064820152608401610b1b565b6129446129296040840160208501613d26565b61014084013561293f60e0860160c08701613d26565b6136a9565b600061296d8261295c61012086016101008701613d26565b8561012001358661014001356137ba565b60408051808201909152919350915060009080600281526020016040518061010001604052807f00000000000000000000000000000000000000000000000000000000000000008152602001306001600160a01b031681526020016129cf3390565b6001600160a01b031681526020016129ed6060890160408a01613d26565b6001600160a01b0316815260208101879052604001612a0f60608901896145d3565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a08801356020820152604001612a64610100890160e08a01613d26565b6001600160a01b03169052604051612a7f919060200161461a565b60405160208183030381529060405281525090506000612b406040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602001604051806040016040528089610100016020810190612b179190613d26565b6001600160a01b0316815260209081018890529082526080890135908201526040016000611033565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168787604051612b7e929190614728565b60405180910390a35050505050565b612bb78235612ba26040850160208601613d26565b612bb260e0860160c08701613d26565b6138d7565b6000612bcf8261295c61012086016101008701613d26565b6040805180820190915291935091506000908060048152602001604051806101600160405280612bfc3390565b6001600160a01b0316815260200187600001358152602001876020016020810190612c279190613d26565b6001600160a01b03168152602001612c456060890160408a01613d26565b6001600160a01b0316815260208101879052604001612c6760608901896145d3565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a08801356020820152604001612cbc610100890160e08a01613d26565b6001600160a01b0316815260808801356020820152604001612ce460e0890160c08a01613d26565b6001600160a01b03168152610140880135602091820152604051612d09929101614837565b60408051601f198184030181529190529052905060006105dc612d39612d3260608801886145d3565b905061163a565b612d439190613f87565b612d509062055730613ecf565b90506000612e166040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808a610100016020810190612dd69190613d26565b6001600160a01b0316815260209081018990529082528181018690526040805160008152808301825281840152516060909201916110739188910161400c565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168888604051612e54929190614728565b60405180910390a3505050505050565b6000612e766060830160408401613d26565b6001600160a01b031603612ed75760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a207a65726f20726563697069656e74206164647265604482015261737360f01b6064820152608401610b1b565b60008160c0013511612efb5760405162461bcd60e51b8152600401610b1b90614590565b8035612f195760405162461bcd60e51b8152600401610b1b906144fb565b6000612f2b6040830160208401613d26565b6001600160a01b031603610c385760405162461bcd60e51b8152600401610b1b90614545565b612f7a612f646040840160208501613d26565b60a084013561293f610100860160e08701613d26565b6000612f9f82612f906080860160608701613d26565b85608001358660a001356137ba565b60408051808201909152919350915060009080600181526020016040518060400160405280876040016020810190612fd79190613d26565b6001600160a01b0316815260200186815250604051602001612ff99190613fd6565b604051602081830303815290604052815250905060006130b96040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808960600160208101906130909190613d26565b6001600160a01b03168152602090810188905290825260c0890135908201526040016000611033565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb528787604051612b7e929190614915565b61311d823561310c6040850160208601613d26565b612bb2610100860160e08701613d26565b600061313382612f906080860160608701613d26565b60408051808201825260038152815160e081018352873581529395509193506000926020808401928282019161316d918a01908a01613d26565b6001600160a01b0316815260200161318b6060890160408a01613d26565b6001600160a01b031681526020810187905260a0880135604082015260c088013560608201526080016131c5610100890160e08a01613d26565b6001600160a01b031690526040516132359190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b604051602081830303815290604052815250905060006130b96040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808960600160208101906132cc9190613d26565b6001600160a01b03168152602090810188905290825262053020908201526040016000611033565b6040516370a0823160e01b815230600482015260009081906001600160a01b038616906370a0823190602401602060405180830381865afa15801561333d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613361919061409f565b90506133786001600160a01b038616853086613988565b6040516370a0823160e01b81523060048201526000906001600160a01b038716906370a0823190602401602060405180830381865afa1580156133bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133e3919061409f565b90508181116134495760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610b1b565b6134538282613ee2565b9695505050505050565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b8260405161349891815260200190565b60405180910390a26117588282611ce9565b6134b43082611ce9565b60008260000151836020015184604001518560a001516040516024016134dd94939291906149b4565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c0840151606085015191925060009161352491908590856139c0565b905080156135785783606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161356b91815260200190565b60405180910390a2611593565b83606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0846040516135b791815260200190565b60405180910390a260e0840151611593906001600160a01b031684611bd0565b600061362c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613a959092919063ffffffff16565b805190915015611bcb578080602001905181019061364a91906149e6565b611bcb5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610b1b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316146136fa5760405162461bcd60e51b8152600401610b1b90614a08565b81156137535760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a206e6f6e2d7a65726f207365636f6e646172792066604482015261656560f01b6064820152608401610b1b565b6001600160a01b03811615611bcb5760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a206e6f6e2d7a65726f206d756c74692d686f702066604482015266616c6c6261636b60c81b6064820152608401610b1b565b6000806137c785856121da565b93506137d286613aa4565b955061381f7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085611d6b565b61386a7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089611d6b565b116138cb5760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a20696e73756666696369656e7420746f6b656e732060448201526a3a37903a3930b739b332b960a91b6064820152608401610b1b565b50939491935090915050565b7f0000000000000000000000000000000000000000000000000000000000000000830361392657306001600160a01b038316036139265760405162461bcd60e51b8152600401610b1b90614a08565b6001600160a01b038116611bcb5760405162461bcd60e51b815260206004820152602360248201527f546f6b656e53706f6b653a207a65726f206d756c74692d686f702066616c6c6260448201526261636b60e81b6064820152608401610b1b565b6040516001600160a01b03808516602483015283166044820152606481018290526115939085906323b872dd60e01b906084016126a5565b6000845a1015613a125760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610b1b565b83471015613a625760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610b1b565b826001600160a01b03163b600003613a7c57506000611d7a565b600080600084516020860188888bf19695505050505050565b6060611d7a8484600085613abd565b6000613ab962010203600160981b0183611bd0565b5090565b606082471015613b1e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610b1b565b600080866001600160a01b03168587604051613b3a9190614a56565b60006040518083038185875af1925050503d8060008114613b77576040519150601f19603f3d011682016040523d82523d6000602084013e613b7c565b606091505b5091509150613b8d87838387613b98565b979650505050505050565b60608315613c07578251600003613c00576001600160a01b0385163b613c005760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b1b565b5081611d7a565b611d7a8383815115613c1c5781518083602001fd5b8060405162461bcd60e51b8152600401610b1b9190613c86565b60005b83811015613c51578181015183820152602001613c39565b50506000910152565b60008151808452613c72816020860160208601613c36565b601f01601f19169290920160200192915050565b602081526000611ea16020830184613c5a565b6001600160a01b0381168114610c3857600080fd5b80356125f281613c99565b60008060408385031215613ccc57600080fd5b8235613cd781613c99565b946020939093013593505050565b600080600060608486031215613cfa57600080fd5b8335613d0581613c99565b92506020840135613d1581613c99565b929592945050506040919091013590565b600060208284031215613d3857600080fd5b8135611ea181613c99565b600060208284031215613d5557600080fd5b5035919050565b600060208284031215613d6e57600080fd5b813567ffffffffffffffff811115613d8557600080fd5b82016101608185031215611ea157600080fd5b60006101008284031215613dab57600080fd5b50919050565b600060408284031215613dab57600080fd5b60008060008060608587031215613dd957600080fd5b843593506020850135613deb81613c99565b9250604085013567ffffffffffffffff80821115613e0857600080fd5b818701915087601f830112613e1c57600080fd5b813581811115613e2b57600080fd5b886020828501011115613e3d57600080fd5b95989497505060200194505050565b60008060408385031215613e5f57600080fd5b8235613e6a81613c99565b91506020830135613e7a81613c99565b809150509250929050565b600181811c90821680613e9957607f821691505b602082108103613dab57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610a6257610a62613eb9565b81810381811115610a6257610a62613eb9565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b8082028115828204841417610a6257610a62613eb9565b600082613fbb57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b81516001600160a01b031681526020808301519082015260408101610a62565b634e487b7160e01b600052604160045260246000fd5b60208152600082516005811061403257634e487b7160e01b600052602160045260246000fd5b806020840152506020830151604080840152611d7a6060840182613c5a565b6020808252602e908201527f4e6174697665546f6b656e53706f6b653a20636f6e747261637420756e64657260408201526d18dbdb1b185d195c985b1a5e995960921b606082015260800190565b6000602082840312156140b157600080fd5b5051919050565b60ff8281168282160390811115610a6257610a62613eb9565b600181815b8085111561410c5781600019048211156140f2576140f2613eb9565b808516156140ff57918102915b93841c93908002906140d6565b509250929050565b60008261412357506001610a62565b8161413057506000610a62565b816001811461414657600281146141505761416c565b6001915050610a62565b60ff84111561416157614161613eb9565b50506001821b610a62565b5060208310610133831016604e8410600b841016171561418f575081810a610a62565b61419983836140d1565b80600019048211156141ad576141ad613eb9565b029392505050565b6000611ea18383614114565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b808410156142505784518316825293850193600193909301929085019061422e565b5060a0880151878203601f190160e0890152945061426e8186613c5a565b98975050505050505050565b6040805190810167ffffffffffffffff8111828210171561429d5761429d613ff6565b60405290565b604051610100810167ffffffffffffffff8111828210171561429d5761429d613ff6565b600082601f8301126142d857600080fd5b815167ffffffffffffffff808211156142f3576142f3613ff6565b604051601f8301601f19908116603f0116810190828211818310171561431b5761431b613ff6565b8160405283815286602085880101111561433457600080fd5b613453846020830160208901613c36565b60006020828403121561435757600080fd5b815167ffffffffffffffff8082111561436f57600080fd5b908301906040828603121561438357600080fd5b61438b61427a565b82516005811061439a57600080fd5b81526020830151828111156143ae57600080fd5b6143ba878286016142c7565b60208301525095945050505050565b80516125f281613c99565b6000604082840312156143e657600080fd5b6143ee61427a565b82516143f981613c99565b81526020928301519281019290925250919050565b60006020828403121561442057600080fd5b815167ffffffffffffffff8082111561443857600080fd5b90830190610100828603121561444d57600080fd5b6144556142a3565b82518152614465602084016143c9565b6020820152614476604084016143c9565b6040820152614487606084016143c9565b60608201526080830151608082015260a0830151828111156144a857600080fd5b6144b4878286016142c7565b60a08301525060c083015160c08201526144d060e084016143c9565b60e082015295945050505050565b6000602082840312156144f057600080fd5b8151611ea181613c99565b6020808252602a908201527f546f6b656e53706f6b653a207a65726f2064657374696e6174696f6e20626c6f60408201526918dad8da185a5b88125160b21b606082015260800190565b6020808252602b908201527f546f6b656e53706f6b653a207a65726f2064657374696e6174696f6e2062726960408201526a646765206164647265737360a81b606082015260800190565b60208082526023908201527f546f6b656e53706f6b653a207a65726f20726571756972656420676173206c696040820152621b5a5d60ea1b606082015260800190565b6000808335601e198436030181126145ea57600080fd5b83018035915067ffffffffffffffff82111561460557600080fd5b60200191503681900382131561169557600080fd5b60208152815160208201526000602083015160018060a01b0380821660408501528060408601511660608501525050606083015161466360808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c085015261468a610120850183613c5a565b915060c085015160e085015260e08501516146af828601826001600160a01b03169052565b5090949350505050565b6000808335601e198436030181126146d057600080fd5b830160208101925035905067ffffffffffffffff8111156146f057600080fd5b80360382131561169557600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6040815282356040820152600061474160208501613cae565b6001600160a01b0316606083015261475b60408501613cae565b6001600160a01b0316608083015261477660608501856146b9565b6101608060a086015261478e6101a0860183856146ff565b9250608087013560c086015260a087013560e08601526147b060c08801613cae565b91506101006147c9818701846001600160a01b03169052565b6147d560e08901613cae565b92506101206147ee818801856001600160a01b03169052565b6147f9828a01613cae565b93506101409150614814828801856001600160a01b03169052565b880135918601919091529095013561018084015260209092019290925292915050565b602081526148516020820183516001600160a01b03169052565b602082015160408201526000604083015161487760608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c08501526148b3610180850183613c5a565b915060c085015160e085015260e08501516101006148db818701836001600160a01b03169052565b860151610120868101919091528601519050610140614904818701836001600160a01b03169052565b959095015193019290925250919050565b823581526101208101602084013561492c81613c99565b6001600160a01b03908116602084015260408501359061494b82613c99565b16604083015261495d60608501613cae565b6001600160a01b0381166060840152506080840135608083015260a084013560a083015260c084013560c083015261499760e08501613cae565b6001600160a01b031660e083015261010090910191909152919050565b8481526001600160a01b0384811660208301528316604082015260806060820181905260009061345390830184613c5a565b6000602082840312156149f857600080fd5b81518015158114611ea157600080fd5b6020808252602e908201527f546f6b656e53706f6b653a20696e76616c69642064657374696e6174696f6e2060408201526d627269646765206164647265737360901b606082015260800190565b60008251614a68818460208701613c36565b919091019291505056fea264697066735822122015a3f13c9733aa3506309e4c1bc9f6b3a4418904c60d5aaf35c5fa27c63578db64736f6c63430008120033",
}

// NativeTokenSpokeABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenSpokeMetaData.ABI instead.
var NativeTokenSpokeABI = NativeTokenSpokeMetaData.ABI

// NativeTokenSpokeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenSpokeMetaData.Bin instead.
var NativeTokenSpokeBin = NativeTokenSpokeMetaData.Bin

// DeployNativeTokenSpoke deploys a new Ethereum contract, binding an instance of NativeTokenSpoke to it.
func DeployNativeTokenSpoke(auth *bind.TransactOpts, backend bind.ContractBackend, settings TokenSpokeSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (common.Address, *types.Transaction, *NativeTokenSpoke, error) {
	parsed, err := NativeTokenSpokeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenSpokeBin), backend, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenSpoke{NativeTokenSpokeCaller: NativeTokenSpokeCaller{contract: contract}, NativeTokenSpokeTransactor: NativeTokenSpokeTransactor{contract: contract}, NativeTokenSpokeFilterer: NativeTokenSpokeFilterer{contract: contract}}, nil
}

// NativeTokenSpoke is an auto generated Go binding around an Ethereum contract.
type NativeTokenSpoke struct {
	NativeTokenSpokeCaller     // Read-only binding to the contract
	NativeTokenSpokeTransactor // Write-only binding to the contract
	NativeTokenSpokeFilterer   // Log filterer for contract events
}

// NativeTokenSpokeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenSpokeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSpokeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenSpokeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSpokeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenSpokeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSpokeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenSpokeSession struct {
	Contract     *NativeTokenSpoke // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeTokenSpokeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenSpokeCallerSession struct {
	Contract *NativeTokenSpokeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// NativeTokenSpokeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenSpokeTransactorSession struct {
	Contract     *NativeTokenSpokeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// NativeTokenSpokeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenSpokeRaw struct {
	Contract *NativeTokenSpoke // Generic contract binding to access the raw methods on
}

// NativeTokenSpokeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenSpokeCallerRaw struct {
	Contract *NativeTokenSpokeCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenSpokeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenSpokeTransactorRaw struct {
	Contract *NativeTokenSpokeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenSpoke creates a new instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpoke(address common.Address, backend bind.ContractBackend) (*NativeTokenSpoke, error) {
	contract, err := bindNativeTokenSpoke(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpoke{NativeTokenSpokeCaller: NativeTokenSpokeCaller{contract: contract}, NativeTokenSpokeTransactor: NativeTokenSpokeTransactor{contract: contract}, NativeTokenSpokeFilterer: NativeTokenSpokeFilterer{contract: contract}}, nil
}

// NewNativeTokenSpokeCaller creates a new read-only instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpokeCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenSpokeCaller, error) {
	contract, err := bindNativeTokenSpoke(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeCaller{contract: contract}, nil
}

// NewNativeTokenSpokeTransactor creates a new write-only instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpokeTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenSpokeTransactor, error) {
	contract, err := bindNativeTokenSpoke(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTransactor{contract: contract}, nil
}

// NewNativeTokenSpokeFilterer creates a new log filterer instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpokeFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenSpokeFilterer, error) {
	contract, err := bindNativeTokenSpoke(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeFilterer{contract: contract}, nil
}

// bindNativeTokenSpoke binds a generic wrapper to an already deployed contract.
func bindNativeTokenSpoke(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenSpokeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSpoke *NativeTokenSpokeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSpoke.Contract.NativeTokenSpokeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSpoke *NativeTokenSpokeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.NativeTokenSpokeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSpoke *NativeTokenSpokeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.NativeTokenSpokeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSpoke *NativeTokenSpokeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSpoke.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSpoke *NativeTokenSpokeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSpoke *NativeTokenSpokeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BURNEDFORBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "BURNED_FOR_BRIDGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenSpoke.CallOpts)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenSpoke.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDTXFEESADDRESS(&_NativeTokenSpoke.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDTXFEESADDRESS(&_NativeTokenSpoke.CallOpts)
}

// HUBCHAINBURNADDRESS is a free data retrieval call binding the contract method 0xeae0544f.
//
// Solidity: function HUB_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) HUBCHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "HUB_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HUBCHAINBURNADDRESS is a free data retrieval call binding the contract method 0xeae0544f.
//
// Solidity: function HUB_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) HUBCHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.HUBCHAINBURNADDRESS(&_NativeTokenSpoke.CallOpts)
}

// HUBCHAINBURNADDRESS is a free data retrieval call binding the contract method 0xeae0544f.
//
// Solidity: function HUB_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) HUBCHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.HUBCHAINBURNADDRESS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenSpoke.Contract.NATIVEMINTER(&_NativeTokenSpoke.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenSpoke.Contract.NATIVEMINTER(&_NativeTokenSpoke.CallOpts)
}

// REGISTERSPOKEREQUIREDGAS is a free data retrieval call binding the contract method 0xc03f4013.
//
// Solidity: function REGISTER_SPOKE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) REGISTERSPOKEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "REGISTER_SPOKE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERSPOKEREQUIREDGAS is a free data retrieval call binding the contract method 0xc03f4013.
//
// Solidity: function REGISTER_SPOKE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) REGISTERSPOKEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.REGISTERSPOKEREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// REGISTERSPOKEREQUIREDGAS is a free data retrieval call binding the contract method 0xc03f4013.
//
// Solidity: function REGISTER_SPOKE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) REGISTERSPOKEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.REGISTERSPOKEREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.Allowance(&_NativeTokenSpoke.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.Allowance(&_NativeTokenSpoke.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BalanceOf(&_NativeTokenSpoke.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BalanceOf(&_NativeTokenSpoke.CallOpts, account)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.BlockchainID(&_NativeTokenSpoke.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.BlockchainID(&_NativeTokenSpoke.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BurnedFeesReportingRewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "burnedFeesReportingRewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenSpoke.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenSpoke.CallOpts)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.CalculateNumWords(&_NativeTokenSpoke.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.CalculateNumWords(&_NativeTokenSpoke.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Decimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.Decimals(&_NativeTokenSpoke.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Decimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.Decimals(&_NativeTokenSpoke.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.GetMinTeleporterVersion(&_NativeTokenSpoke.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.GetMinTeleporterVersion(&_NativeTokenSpoke.CallOpts)
}

// HubTokenDecimals is a free data retrieval call binding the contract method 0xfdd7546f.
//
// Solidity: function hubTokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) HubTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "hubTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HubTokenDecimals is a free data retrieval call binding the contract method 0xfdd7546f.
//
// Solidity: function hubTokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeSession) HubTokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.HubTokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// HubTokenDecimals is a free data retrieval call binding the contract method 0xfdd7546f.
//
// Solidity: function hubTokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) HubTokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.HubTokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) InitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "initialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.InitialReserveImbalance(&_NativeTokenSpoke.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.InitialReserveImbalance(&_NativeTokenSpoke.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) IsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "isCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IsCollateralized() (bool, error) {
	return _NativeTokenSpoke.Contract.IsCollateralized(&_NativeTokenSpoke.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) IsCollateralized() (bool, error) {
	return _NativeTokenSpoke.Contract.IsCollateralized(&_NativeTokenSpoke.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) IsRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "isRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IsRegistered() (bool, error) {
	return _NativeTokenSpoke.Contract.IsRegistered(&_NativeTokenSpoke.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) IsRegistered() (bool, error) {
	return _NativeTokenSpoke.Contract.IsRegistered(&_NativeTokenSpoke.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenSpoke.Contract.IsTeleporterAddressPaused(&_NativeTokenSpoke.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenSpoke.Contract.IsTeleporterAddressPaused(&_NativeTokenSpoke.CallOpts, teleporterAddress)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) LastestBurnedFeesReported(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "lastestBurnedFeesReported")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.LastestBurnedFeesReported(&_NativeTokenSpoke.CallOpts)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.LastestBurnedFeesReported(&_NativeTokenSpoke.CallOpts)
}

// MultiplyOnSpoke is a free data retrieval call binding the contract method 0xfeea2b72.
//
// Solidity: function multiplyOnSpoke() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MultiplyOnSpoke(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "multiplyOnSpoke")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplyOnSpoke is a free data retrieval call binding the contract method 0xfeea2b72.
//
// Solidity: function multiplyOnSpoke() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MultiplyOnSpoke() (bool, error) {
	return _NativeTokenSpoke.Contract.MultiplyOnSpoke(&_NativeTokenSpoke.CallOpts)
}

// MultiplyOnSpoke is a free data retrieval call binding the contract method 0xfeea2b72.
//
// Solidity: function multiplyOnSpoke() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MultiplyOnSpoke() (bool, error) {
	return _NativeTokenSpoke.Contract.MultiplyOnSpoke(&_NativeTokenSpoke.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Name() (string, error) {
	return _NativeTokenSpoke.Contract.Name(&_NativeTokenSpoke.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Name() (string, error) {
	return _NativeTokenSpoke.Contract.Name(&_NativeTokenSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Owner() (common.Address, error) {
	return _NativeTokenSpoke.Contract.Owner(&_NativeTokenSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Owner() (common.Address, error) {
	return _NativeTokenSpoke.Contract.Owner(&_NativeTokenSpoke.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Symbol() (string, error) {
	return _NativeTokenSpoke.Contract.Symbol(&_NativeTokenSpoke.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Symbol() (string, error) {
	return _NativeTokenSpoke.Contract.Symbol(&_NativeTokenSpoke.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TeleporterRegistry(&_NativeTokenSpoke.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TeleporterRegistry(&_NativeTokenSpoke.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.TokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.TokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// TokenHubAddress is a free data retrieval call binding the contract method 0x24427629.
//
// Solidity: function tokenHubAddress() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenHubAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenHubAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenHubAddress is a free data retrieval call binding the contract method 0x24427629.
//
// Solidity: function tokenHubAddress() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenHubAddress() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TokenHubAddress(&_NativeTokenSpoke.CallOpts)
}

// TokenHubAddress is a free data retrieval call binding the contract method 0x24427629.
//
// Solidity: function tokenHubAddress() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenHubAddress() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TokenHubAddress(&_NativeTokenSpoke.CallOpts)
}

// TokenHubBlockchainID is a free data retrieval call binding the contract method 0xeb3da909.
//
// Solidity: function tokenHubBlockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenHubBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenHubBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenHubBlockchainID is a free data retrieval call binding the contract method 0xeb3da909.
//
// Solidity: function tokenHubBlockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenHubBlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.TokenHubBlockchainID(&_NativeTokenSpoke.CallOpts)
}

// TokenHubBlockchainID is a free data retrieval call binding the contract method 0xeb3da909.
//
// Solidity: function tokenHubBlockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenHubBlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.TokenHubBlockchainID(&_NativeTokenSpoke.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TokenMultiplier(&_NativeTokenSpoke.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TokenMultiplier(&_NativeTokenSpoke.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalMinted(&_NativeTokenSpoke.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalMinted(&_NativeTokenSpoke.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalNativeAssetSupply(&_NativeTokenSpoke.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalNativeAssetSupply(&_NativeTokenSpoke.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalSupply(&_NativeTokenSpoke.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalSupply(&_NativeTokenSpoke.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Approve(&_NativeTokenSpoke.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Approve(&_NativeTokenSpoke.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.DecreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.DecreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Deposit(&_NativeTokenSpoke.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Deposit(&_NativeTokenSpoke.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.IncreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.IncreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, addedValue)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.PauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.PauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReceiveTeleporterMessage(&_NativeTokenSpoke.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReceiveTeleporterMessage(&_NativeTokenSpoke.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHub is a paid mutator transaction binding the contract method 0xbf48d36b.
//
// Solidity: function registerWithHub((address,uint256) feeInfo) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) RegisterWithHub(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "registerWithHub", feeInfo)
}

// RegisterWithHub is a paid mutator transaction binding the contract method 0xbf48d36b.
//
// Solidity: function registerWithHub((address,uint256) feeInfo) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) RegisterWithHub(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RegisterWithHub(&_NativeTokenSpoke.TransactOpts, feeInfo)
}

// RegisterWithHub is a paid mutator transaction binding the contract method 0xbf48d36b.
//
// Solidity: function registerWithHub((address,uint256) feeInfo) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) RegisterWithHub(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RegisterWithHub(&_NativeTokenSpoke.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RenounceOwnership(&_NativeTokenSpoke.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RenounceOwnership(&_NativeTokenSpoke.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReportBurnedTxFees(&_NativeTokenSpoke.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReportBurnedTxFees(&_NativeTokenSpoke.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Send(&_NativeTokenSpoke.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Send(&_NativeTokenSpoke.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.SendAndCall(&_NativeTokenSpoke.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.SendAndCall(&_NativeTokenSpoke.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Transfer(&_NativeTokenSpoke.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Transfer(&_NativeTokenSpoke.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferFrom(&_NativeTokenSpoke.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferFrom(&_NativeTokenSpoke.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferOwnership(&_NativeTokenSpoke.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferOwnership(&_NativeTokenSpoke.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UnpauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UnpauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UpdateMinTeleporterVersion(&_NativeTokenSpoke.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UpdateMinTeleporterVersion(&_NativeTokenSpoke.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Withdraw(&_NativeTokenSpoke.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Withdraw(&_NativeTokenSpoke.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Fallback(&_NativeTokenSpoke.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Fallback(&_NativeTokenSpoke.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Receive() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Receive(&_NativeTokenSpoke.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Receive(&_NativeTokenSpoke.TransactOpts)
}

// NativeTokenSpokeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeApprovalIterator struct {
	Event *NativeTokenSpokeApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeApproval represents a Approval event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenSpokeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeApprovalIterator{contract: _NativeTokenSpoke.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeApproval)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseApproval(log types.Log) (*NativeTokenSpokeApproval, error) {
	event := new(NativeTokenSpokeApproval)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallFailedIterator struct {
	Event *NativeTokenSpokeCallFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeCallFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeCallFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeCallFailed represents a CallFailed event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenSpokeCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeCallFailedIterator{contract: _NativeTokenSpoke.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeCallFailed)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseCallFailed(log types.Log) (*NativeTokenSpokeCallFailed, error) {
	event := new(NativeTokenSpokeCallFailed)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallSucceededIterator struct {
	Event *NativeTokenSpokeCallSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeCallSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeCallSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeCallSucceeded represents a CallSucceeded event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenSpokeCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeCallSucceededIterator{contract: _NativeTokenSpoke.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeCallSucceeded)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenSpokeCallSucceeded, error) {
	event := new(NativeTokenSpokeCallSucceeded)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeDepositIterator struct {
	Event *NativeTokenSpokeDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeDeposit represents a Deposit event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenSpokeDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeDepositIterator{contract: _NativeTokenSpoke.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeDeposit)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseDeposit(log types.Log) (*NativeTokenSpokeDeposit, error) {
	event := new(NativeTokenSpokeDeposit)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenSpokeMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeMinTeleporterVersionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeMinTeleporterVersionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenSpokeMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeMinTeleporterVersionUpdatedIterator{contract: _NativeTokenSpoke.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeMinTeleporterVersionUpdated)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenSpokeMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenSpokeMinTeleporterVersionUpdated)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeOwnershipTransferredIterator struct {
	Event *NativeTokenSpokeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenSpokeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeOwnershipTransferredIterator{contract: _NativeTokenSpoke.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeOwnershipTransferred)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenSpokeOwnershipTransferred, error) {
	event := new(NativeTokenSpokeOwnershipTransferred)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeReportBurnedTxFeesIterator struct {
	Event *NativeTokenSpokeReportBurnedTxFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeReportBurnedTxFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeReportBurnedTxFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenSpokeReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeReportBurnedTxFeesIterator{contract: _NativeTokenSpoke.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeReportBurnedTxFees)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportBurnedTxFees is a log parse operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenSpokeReportBurnedTxFees, error) {
	event := new(NativeTokenSpokeReportBurnedTxFees)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressPausedIterator struct {
	Event *NativeTokenSpokeTeleporterAddressPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTeleporterAddressPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeTeleporterAddressPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenSpokeTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTeleporterAddressPausedIterator{contract: _NativeTokenSpoke.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTeleporterAddressPaused)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenSpokeTeleporterAddressPaused, error) {
	event := new(NativeTokenSpokeTeleporterAddressPaused)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenSpokeTeleporterAddressUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTeleporterAddressUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeTeleporterAddressUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenSpokeTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTeleporterAddressUnpausedIterator{contract: _NativeTokenSpoke.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTeleporterAddressUnpaused)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenSpokeTeleporterAddressUnpaused, error) {
	event := new(NativeTokenSpokeTeleporterAddressUnpaused)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensAndCallSentIterator struct {
	Event *NativeTokenSpokeTokensAndCallSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTokensAndCallSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeTokensAndCallSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenSpokeTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTokensAndCallSentIterator{contract: _NativeTokenSpoke.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTokensAndCallSent)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenSpokeTokensAndCallSent, error) {
	event := new(NativeTokenSpokeTokensAndCallSent)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensSentIterator struct {
	Event *NativeTokenSpokeTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTokensSent represents a TokensSent event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenSpokeTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTokensSentIterator{contract: _NativeTokenSpoke.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTokensSent)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensSent is a log parse operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTokensSent(log types.Log) (*NativeTokenSpokeTokensSent, error) {
	event := new(NativeTokenSpokeTokensSent)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensWithdrawnIterator struct {
	Event *NativeTokenSpokeTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenSpokeTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTokensWithdrawnIterator{contract: _NativeTokenSpoke.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTokensWithdrawn)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenSpokeTokensWithdrawn, error) {
	event := new(NativeTokenSpokeTokensWithdrawn)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTransferIterator struct {
	Event *NativeTokenSpokeTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTransfer represents a Transfer event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenSpokeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTransferIterator{contract: _NativeTokenSpoke.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTransfer)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTransfer(log types.Log) (*NativeTokenSpokeTransfer, error) {
	event := new(NativeTokenSpokeTransfer)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeWithdrawalIterator struct {
	Event *NativeTokenSpokeWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenSpokeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenSpokeWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenSpokeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeWithdrawal represents a Withdrawal event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenSpokeWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeWithdrawalIterator{contract: _NativeTokenSpoke.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeWithdrawal)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseWithdrawal(log types.Log) (*NativeTokenSpokeWithdrawal, error) {
	event := new(NativeTokenSpokeWithdrawal)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}