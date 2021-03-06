package txsender

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/client"
	alaTypes "github.com/onmax/go-alastria/types"
)

func AddKey(conn *alaTypes.Connection) (*types.Transaction, error) {
	err := client.CheckIdentityManager(conn)
	if err != nil {
		return nil, err
	}

	err = client.CheckPublickeyRegistry(conn)
	if err != nil {
		return nil, err
	}

	err = client.CheckTxOpts(conn)
	if err != nil {
		return nil, err
	}

	err = client.CheckKeystore(conn)
	if err != nil {
		return nil, err
	}
	return conn.Contracts.Instances.PublicKeyRegistry.AddKey(conn.Tx.Opts, conn.Client.Ks.HexPublicKey)
}

func GetCurrentPublicKey(conn *alaTypes.Connection, agentAddress common.Address) (string, error) {
	err := client.CheckPublickeyRegistry(conn)
	if err != nil {
		return "", err
	}

	return conn.Contracts.Instances.PublicKeyRegistry.GetCurrentPublicKey(&bind.CallOpts{}, agentAddress)
}
