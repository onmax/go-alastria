package tx

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	addKeyTx, err := conn.Contracts.PublicKeyRegistry.AddKey(conn.Tx.Opts, conn.Client.Ks.HexPublicKey)
	if err != nil {
		return nil, err
	}
	// signedTxAddKey, _ := crypto.SignTx(addKeyTx, conn.Tx.Signer, conn.Client.Ks.PrivateKey, conn.Network.Id)
	fmt.Printf("addKeyTx.Data(): %v\n", addKeyTx.Data())
	// Update nonce by hand
	// TODO Check if this is right!
	// conn.Tx.Opts.Nonce = new(big.Int).SetUint64(nonce + 1)

	return conn.Contracts.IdentityManager.CreateAlastriaIdentity(conn.Tx.Opts, addKeyTx.Data())
}

func PrepareAlastriaId(conn *alaTypes.Connection, newAgentPubKey string) (*types.Transaction, error) {
	// TODO Check for txOpts, contracts and keystore
	if conn.Contracts.IdentityManager == nil {
		instance, err := network.IdentityManagerContract(conn.Client.Eth)
		if err != nil {
			return nil, err
		}
		conn.Contracts.IdentityManager = instance
	}

	return conn.Contracts.IdentityManager.PrepareAlastriaID(conn.Tx.Opts, common.HexToAddress(newAgentPubKey))
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	return conn.Contracts.IdentityManager.IdentityKeys(&bind.CallOpts{}, agentAddress)
}
