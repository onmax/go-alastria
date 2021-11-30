package alastria

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/network"
	"github.com/onmax/go-alastria/tx"
	alaTypes "github.com/onmax/go-alastria/types"
)

func NewClient(args *alaTypes.ConnectionArgs) (*alaTypes.Connection, error) {
	client, err := network.ConnectToNetwork(args.NodeUrl)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connected!: client: %v\n", client)
	networkId, err := network.NetworkID(client)
	if err != nil {
		return nil, err
	}

	opts := &bind.TransactOpts{}
	if args.Keystore != nil {
		opts, err = tx.TxOpt(args.Keystore.PrivateKey, networkId)
		if err != nil {
			return nil, err
		}
	}

	return &alaTypes.Connection{
		Client: &alaTypes.Client{
			Eth: client,
			Ks:  args.Keystore,
		},
		Contracts: &alaTypes.Contracts{
			IdentityManager:   nil,
			PublicKeyRegistry: nil,
		},
		Network: &alaTypes.Network{
			Id: networkId,
		},
		Tx: &alaTypes.TxClient{
			Opts:   opts,
			Signer: types.NewEIP155Signer(networkId),
		},
	}, nil
}

func PrepareAlastriaId(conn *alaTypes.Connection, newAgentPubKey string) (*types.Transaction, error) {
	return tx.PrepareAlastriaId(conn, newAgentPubKey)
}

func CreateAlastriaIdentity(conn *alaTypes.Connection) (*types.Transaction, error) {
	nonce, err := Nonce(conn, conn.Client.Ks.Account.Address)
	if err != nil {
		return nil, err
	}

	return tx.CreateAlastriaIdentity(conn, nonce)
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	if conn.Contracts.IdentityManager == nil {
		instance, err := network.IdentityManagerContract(conn.Client.Eth)
		if err != nil {
			return common.Address{}, err
		}
		conn.Contracts.IdentityManager = instance
	}
	return tx.IdentityKeys(conn, agentAddress)
}

func Nonce(conn *alaTypes.Connection, from common.Address) (uint64, error) {
	return conn.Client.Eth.PendingNonceAt(context.Background(), from)
}

// func SignTx(conn *alaTypes.Connection, tx *types.Transaction) (*types.Transaction, error) {
// 	return crypto.SignTx(tx, conn.Tx.Signer, conn.Client.Ks.PrivateKey, conn.Network.Id)
// }
