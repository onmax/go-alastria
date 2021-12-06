package tx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/contracts"
	"github.com/onmax/go-alastria/internal/addresses"
	alaTypes "github.com/onmax/go-alastria/types"
)

func CreateAlastriaIdentity(conn *alaTypes.Connection) (*types.Transaction, error) {
	// TODO Check for txOpts, contracts and keystore
	if conn.Contracts.IdentityManager == nil {
		instance, err := contracts.IdentityManagerContract(conn.Client.Eth)
		if err != nil {
			return nil, err
		}
		conn.Contracts.IdentityManager = instance
	}
	if conn.Contracts.PublicKeyRegistry == nil {
		instance, err := contracts.PublicKeyRegistryContract(conn.Client.Eth)
		if err != nil {
			return nil, err
		}
		conn.Contracts.PublicKeyRegistry = instance
	}
	addKeyTx, err := conn.Contracts.PublicKeyRegistry.AddKey(conn.Tx.Opts, conn.Client.Ks.HexPublicKey)
	if err != nil {
		return nil, err
	}
	tx, _ := conn.Contracts.IdentityManager.CreateAlastriaIdentity(conn.Tx.Opts, addKeyTx.Data())
	return tx, nil
}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorAddress common.Address) (*types.Transaction, error) {
	// TODO Check for txOpts, contracts and keystore
	if conn.Contracts.IdentityManager == nil {
		instance, err := contracts.IdentityManagerContract(conn.Client.Eth)
		if err != nil {
			return nil, err
		}
		conn.Contracts.IdentityManager = instance
	}
	delegatedData, _ := conn.Contracts.IdentityManager.PrepareAlastriaID(conn.Tx.Opts, newActorAddress)
	// conn.Tx.Opts.Nonce = new(big.Int).SetUint64(nonce)
	delegated, _ := conn.Contracts.IdentityManager.DelegateCall(
		conn.Tx.Opts, common.HexToAddress(addresses.AlastriaIdentityManager), big.NewInt(0), delegatedData.Data())
	return delegated, nil
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	return conn.Contracts.IdentityManager.IdentityKeys(&bind.CallOpts{}, agentAddress)
}
