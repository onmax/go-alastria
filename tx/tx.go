package tx

import (
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/contracts"
	alaTypes "github.com/onmax/go-alastria/types"
)

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func checkTxOpts(conn *alaTypes.Connection) error {
	if conn.Tx.Opts == nil {
		return alaTypes.ErrTxOptsNotSet
	}
	return nil
}

func checkKeystore(conn *alaTypes.Connection) error {
	if conn.Client.Ks == nil {
		return alaTypes.ErrKeystoreNotSet
	}
	return nil
}

func checkIdentityManager(conn *alaTypes.Connection) error {
	if conn.Contracts.Instances.IdentityManager == nil {
		if isNil(conn.Contracts.Addresses.IdentityManager) {
			return alaTypes.ErrAddressNotSet
		}
		instance, err := contracts.IdentityManagerContract(conn.Client.Eth, conn.Contracts.Addresses.IdentityManager)
		if err != nil {
			return err
		}
		conn.Contracts.Instances.IdentityManager = instance
	}
	return nil
}

func checkPublickeyRegistry(conn *alaTypes.Connection) error {
	if conn.Contracts.Instances.PublicKeyRegistry == nil {
		if isNil(conn.Contracts.Addresses.PublicKeyRegistry) {
			return alaTypes.ErrAddressNotSet
		}
		instance, err := contracts.PublicKeyRegistryContract(conn.Client.Eth, conn.Contracts.Addresses.PublicKeyRegistry)
		if err != nil {
			return err
		}
		conn.Contracts.Instances.PublicKeyRegistry = instance
	}
	return nil
}

func CreateAlastriaIdentity(conn *alaTypes.Connection) (*types.Transaction, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return nil, err
	}
	err = checkPublickeyRegistry(conn)
	if err != nil {
		return nil, err
	}

	err = checkTxOpts(conn)
	if err != nil {
		return nil, err
	}

	err = checkKeystore(conn)
	if err != nil {
		return nil, err
	}

	addKeyTx, err := conn.Contracts.Instances.PublicKeyRegistry.AddKey(conn.Tx.Opts, conn.Client.Ks.HexPublicKey)
	if err != nil {
		return nil, err
	}
	tx, _ := conn.Contracts.Instances.IdentityManager.CreateAlastriaIdentity(conn.Tx.Opts, addKeyTx.Data())
	return tx, nil
}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorAddress common.Address) (*types.Transaction, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return nil, err
	}

	err = checkTxOpts(conn)
	if err != nil {
		return nil, err
	}

	err = checkKeystore(conn)
	if err != nil {
		return nil, err
	}

	delegatedData, _ := conn.Contracts.Instances.IdentityManager.PrepareAlastriaID(conn.Tx.Opts, newActorAddress)
	delegated, _ := conn.Contracts.Instances.IdentityManager.DelegateCall(
		conn.Tx.Opts, common.HexToAddress(conn.Contracts.Addresses.IdentityManager.Hex()), big.NewInt(0), delegatedData.Data())
	return delegated, nil
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return common.Address{}, err
	}

	return conn.Contracts.Instances.IdentityManager.IdentityKeys(&bind.CallOpts{}, agentAddress)
}
