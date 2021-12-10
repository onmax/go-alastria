package sender

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	alaTypes "github.com/onmax/go-alastria/types"
)

func AddKey(conn *alaTypes.Connection) (*types.Transaction, error) {
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
	return conn.Contracts.Instances.PublicKeyRegistry.AddKey(conn.Tx.Opts, conn.Client.Ks.HexPublicKey)
}

func GetCurrentPublicKey(conn *alaTypes.Connection, agentAddress common.Address) (string, error) {
	err := checkPublickeyRegistry(conn)
	if err != nil {
		return "", err
	}

	return conn.Contracts.Instances.PublicKeyRegistry.GetCurrentPublicKey(&bind.CallOpts{}, agentAddress)
}
