package txsender

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	alaTypes "github.com/onmax/go-alastria/types"
)

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

	return conn.Contracts.Instances.IdentityManager.PrepareAlastriaID(conn.Tx.Opts, newActorAddress)
}

func DelegateCall(conn *alaTypes.Connection, delegatedData *types.Transaction) (*types.Transaction, error) {
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

	identityManagerHex := common.HexToAddress(conn.Contracts.Addresses.IdentityManager.Hex())

	delegated, err := conn.Contracts.Instances.IdentityManager.DelegateCall(conn.Tx.Opts, identityManagerHex, big.NewInt(0), delegatedData.Data())
	if err != nil {
		return nil, err
	}

	return delegated, nil
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return common.Address{}, err
	}

	return conn.Contracts.Instances.IdentityManager.IdentityKeys(&bind.CallOpts{}, agentAddress)
}

func CreateAlastriaIdentity(conn *alaTypes.Connection, addPublicKeyCallData []byte) (*types.Transaction, error) {
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
	return conn.Contracts.Instances.IdentityManager.CreateAlastriaIdentity(conn.Tx.Opts, addPublicKeyCallData)
}
