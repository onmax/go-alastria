package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/network"
	"github.com/onmax/go-alastria/tx"
	alaTypes "github.com/onmax/go-alastria/types"
)

// Initialites the client that any actor interacting with the network should use. It can be:
// subject, issuer, or service provider.
// args.NodeUrl is mandatory
// args.Keystore is not mandatory, but it is required if you want to sign JWT or tx
func NewClient(args *alaTypes.ConnectionArgs) (*alaTypes.Connection, error) {
	if args.NodeUrl == "" {
		return nil, alaTypes.ErrNodeUrlNotSet
	}

	client, err := network.ConnectToNetwork(args.NodeUrl)
	if err != nil {
		return nil, err
	}

	conn := &alaTypes.Connection{
		Client: &alaTypes.Client{
			Eth: client,
			Ks:  &alaTypes.Keystore{},
		},
		Tx: &alaTypes.TxClient{
			Opts: &bind.TransactOpts{},
		},
		Contracts: &alaTypes.Contracts{},
		Network:   &alaTypes.Network{},
	}

	err = SetKeystore(conn, args.Keystore)
	if err != nil {
		return nil, err
	}

	err = setNetwork(conn, args)
	if err != nil {
		return nil, err
	}

	err = setOpts(conn, args)
	if err != nil {
		return nil, err
	}

	// Instances of the contracts. This is the only way to interact with the contracts.
	// They will be initialized before the first interaction with the contract and the
	// instance of the contract will be cached.
	conn.Contracts = &alaTypes.Contracts{}

	return conn, nil
}

// Configuration of all the tx that the client will make
func setOpts(conn *alaTypes.Connection, args *alaTypes.ConnectionArgs) error {
	if conn.Client.Ks == nil {
		return nil
	}
	opts, err := tx.TxOpt(conn.Client.Ks.PrivateKey, conn.Network.Id)
	if err != nil {
		return err
	}
	nonce, err := conn.Client.Eth.PendingNonceAt(context.Background(), conn.Client.Ks.Account.Address)
	if err != nil {
		return err
	}
	opts.Nonce = new(big.Int).SetUint64(nonce)
	conn.Tx = &alaTypes.TxClient{
		Opts:   opts,
		Signer: types.NewEIP155Signer(conn.Network.Id),
	}
	return nil
}

// Settings network configurations
func setNetwork(conn *alaTypes.Connection, args *alaTypes.ConnectionArgs) error {
	networkId, err := network.NetworkID(conn.Client.Eth)
	if err != nil {
		return err
	}
	conn.Network = &alaTypes.Network{
		Id: networkId,
	}
	return nil
}

// Set the keystore that will be used to sign the transactions and JWTs
func SetKeystore(conn *alaTypes.Connection, ksConfig *alaTypes.KeystoreConfig) error {
	if ksConfig == nil {
		return alaTypes.ErrKeystoreNotSet
	}
	ks, err := keystore.ImportKs(ksConfig.Path, ksConfig.Password)
	if err != nil {
		return err
	}
	conn.Client.Ks = ks
	return nil
}
