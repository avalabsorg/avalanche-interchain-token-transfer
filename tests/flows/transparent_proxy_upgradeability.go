package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/ERC20TokenHome"
	erc20tokenremote "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenRemote/ERC20TokenRemote"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func TransparentUpgradeableProxy(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the token to be bridged
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy the ERC20TokenHome logic contract on primary network
	erc20TokenHomeAddress, proxyAdmin, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		tokenDecimals,
	)

	// Deploy the ERC20TokenRemote contract on Subnet A
	erc20TokenRemoteAddress, erc20TokenRemote := utils.DeployERC20TokenRemote(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		tokenDecimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetAInfo,
		erc20TokenRemoteAddress,
	)

	// Send a bridge transfer from primary network to Subnet A
	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	input := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenRemoteAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, bridgedAmount := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		input,
		amount,
		fundedKey,
	)

	// Check that the bridge transfer was successful, and expected balances are correct
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20TokenRemoteWithdrawal(
		ctx,
		erc20TokenRemote,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemote.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	// Deploy a new ERC20TokenHome logic contract on primary network
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		cChainInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())

	newLogic, tx, _, err := erc20tokenhome.DeployERC20TokenHome(
		opts,
		cChainInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Upgrade the TransparentUpgradeableProxy contract to use the new logic contract
	opts, err = bind.NewKeyedTransactorWithChainID(
		fundedKey,
		cChainInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())
	tx, err = proxyAdmin.Upgrade(opts, erc20TokenHomeAddress, newLogic)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Send a bridge transfer from Subnet A back to primary network
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputB := erc20tokenremote.SendTokensInput{
		DestinationBlockchainID:  cChainInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenHomeAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   erc20TokenRemoteAddress,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}

	receipt, bridgedAmount = utils.SendERC20TokenRemote(
		ctx,
		subnetAInfo,
		erc20TokenRemote,
		erc20TokenRemoteAddress,
		inputB,
		teleporterUtils.BigIntSub(bridgedAmount, inputB.PrimaryFee),
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		cChainInfo,
		true,
	)

	// Check that the bridge transfer was successful, and expected balances are correct
	utils.CheckERC20TokenHomeWithdrawal(
		ctx,
		erc20TokenHomeAddress,
		exampleERC20,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))
}
