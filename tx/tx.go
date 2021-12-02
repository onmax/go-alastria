package tx

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/internal/addresses"
	"github.com/onmax/go-alastria/network"
	alaTypes "github.com/onmax/go-alastria/types"
)

func CreateAlastriaIdentity(conn *alaTypes.Connection, nonce uint64) (*types.Transaction, error) {
	// TODO Check for txOpts, contracts and keystore
	if conn.Contracts.IdentityManager == nil {
		instance, err := network.IdentityManagerContract(conn.Client.Eth)
		if err != nil {
			return nil, err
		}
		conn.Contracts.IdentityManager = instance
	}
	if conn.Contracts.PublicKeyRegistry == nil {
		instance, err := network.PublicKeyRegistryContract(conn.Client.Eth)
		if err != nil {
			return nil, err
		}
		conn.Contracts.PublicKeyRegistry = instance
	}
	fmt.Printf("/0/0// conn.Client.Ks.HexPublicKey: %v\n", conn.Client.Ks.HexPublicKey)
	addKeyTx, err := conn.Contracts.PublicKeyRegistry.AddKey(&bind.TransactOpts{
		From:      conn.Tx.Opts.From,
		Signer:    conn.Tx.Opts.Signer,
		Nonce:     conn.Tx.Opts.Nonce,
		Value:     conn.Tx.Opts.Value,
		GasLimit:  conn.Tx.Opts.GasLimit,
		GasPrice:  conn.Tx.Opts.GasPrice,
		GasFeeCap: conn.Tx.Opts.GasFeeCap,
		GasTipCap: conn.Tx.Opts.GasTipCap,
		Context:   conn.Tx.Opts.Context,
		NoSend:    true,
	}, conn.Client.Ks.HexPublicKey)
	if err != nil {
		return nil, err
	}
	// signedTxAddKey, _ := crypto.SignTx(addKeyTx, conn.Tx.Signer, conn.Client.Ks.PrivateKey, conn.Network.Id)
	fmt.Printf("addKeyTx.Data(): %x\n", addKeyTx.Data())
	// Update nonce by hand
	// TODO Check if this is right!
	conn.Tx.Opts.Nonce = new(big.Int).SetUint64(nonce)
	tx, _ := conn.Contracts.IdentityManager.CreateAlastriaIdentity(conn.Tx.Opts, addKeyTx.Data())
	fmt.Printf(" ---------}}}}}}}}}}}}}- tx.Data(): %x\n", tx.Data())
	fmt.Printf("-------------------")
	fmt.Printf("\ntx: %+v\n", conn.Tx.Opts)
	fmt.Printf("\ntx: %+v\n", tx)
	fmt.Printf("-------------------\n")
	// tx.To = common.HexToAddress(addresses.AlastriaIdentityManager)
	return tx, nil
}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorAddress common.Address, nonce uint64) (*types.Transaction, error) {
	// TODO Check for txOpts, contracts and keystore
	if conn.Contracts.IdentityManager == nil {
		instance, err := network.IdentityManagerContract(conn.Client.Eth)
		if err != nil {
			return nil, err
		}
		conn.Contracts.IdentityManager = instance
	}
	//     const delegatedData = web3.eth.abi.encodeFunctionCall(config_1.config.contractsAbi.AlastriaIdentityManager.prepareAlastriaID, [AddressUtils_1.AddressUtils.getAddressWithHexPrefix(signAddress)]);
	// tx, _ := conn.Contracts.IdentityManager.AlastriaContractsTransactor.PrepareAlastriaID(conn.Tx.Opts, common.HexToAddress(newAgentPubKey))
	delegatedData, _ := conn.Contracts.IdentityManager.PrepareAlastriaID(conn.Tx.Opts, newActorAddress)
	fmt.Printf("newAgentPubKey: %s\n", newActorAddress)
	fmt.Printf("____________________________________delegatedData: %x\n", delegatedData.Data())
	conn.Tx.Opts.Nonce = new(big.Int).SetUint64(nonce)
	delegated, _ := conn.Contracts.IdentityManager.DelegateCall(
		conn.Tx.Opts, common.HexToAddress(addresses.AlastriaIdentityManager), big.NewInt(0), delegatedData.Data())
	fmt.Printf("____________________________________delegated: %x\n", delegated.Data())
	return delegated, nil
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	return conn.Contracts.IdentityManager.IdentityKeys(&bind.CallOpts{}, agentAddress)
}
