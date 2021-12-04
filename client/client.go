package alastria

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	networkId, err := network.NetworkID(client)
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), args.Keystore.Account.Address)
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
	conn := &alaTypes.Connection{
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
	}
	conn.Tx.Opts.Nonce = new(big.Int).SetUint64(nonce)
	return conn, nil
}
