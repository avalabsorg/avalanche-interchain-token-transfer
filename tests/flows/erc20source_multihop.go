package flows

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func ERC20SourceMultihop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the source token to be bridged
	sourceTokenAddress, sourceToken := teleporterUtils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create an ERC20Source for bridging the source token
	erc20SourceAddress, erc20Source := utils.DeployERC20Source(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		sourceTokenAddress,
	)

	// Token representation on subnet B will have same name, symbol, and decimals
	tokenName, err := sourceToken.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := sourceToken.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := sourceToken.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination for Subnet A
	erc20DestinationAddress_A, erc20Destination_A := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Deploy an ERC20Destination for Subnet B
	erc20DestinationAddress_B, erc20Destination_B := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-chain to Subnet A
	input := erc20source.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress_A,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}
	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, bridgedAmount := utils.SendERC20Source(
		ctx,
		cChainInfo,
		erc20Source,
		erc20SourceAddress,
		sourceToken,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to subnet A and check for ERC20Destination withdrawal
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20DestinationWithdrawal(
		ctx,
		erc20Destination_A,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20Destination_A.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	// Multihop transfer to Subnet B
	sendERC20MultihopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetAInfo,
		erc20Destination_A,
		erc20DestinationAddress_A,
		subnetBInfo,
		erc20Destination_B,
		erc20DestinationAddress_B,
		cChainInfo,
		erc20Source,
		bridgedAmount,
	)

	sendERC20MultihopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetBInfo,
		erc20Destination_B,
		erc20DestinationAddress_B,
		subnetAInfo,
		erc20Destination_A,
		erc20DestinationAddress_A,
		cChainInfo,
		erc20Source,
		bridgedAmount,
	)
}

func sendERC20MultihopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	fundedKey *ecdsa.PrivateKey,
	recipientKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *erc20destination.ERC20Destination,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *erc20destination.ERC20Destination,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	erc20Source *erc20source.ERC20Source,
	bridgedAmount *big.Int,
) {
	teleporterUtils.SendNativeTransfer(
		ctx,
		fromSubnet,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	input_A := erc20destination.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}

	receipt, bridgedAmount := utils.SendERC20Destination(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input_A,
		bridgedAmount,
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		fromSubnet,
		cChainInfo,
		true,
	)

	hopSendEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseSendTokens)
	Expect(err).Should(BeNil())

	Expect(hopSendEvent.Amount).Should(Equal(bridgedAmount))

	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		toSubnet,
		true,
	)

	utils.CheckERC20DestinationWithdrawal(
		ctx,
		toBridge,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	balance, err := toBridge.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))
}
