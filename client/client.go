package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/network"
	"github.com/onmax/go-alastria/tx"
	alaTypes "github.com/onmax/go-alastria/types"
)

// Initializes the client that any actor interacting with the network should use. It can be:
// subject, issuer, or service provider.
// if args.NodeUrl, then the client will not connect to the network
// args.Keystore is not mandatory, but it is required if you want to sign JWT or tx
func NewClient(args *alaTypes.ClientConf) (*alaTypes.Connection, error) {
	conn := &alaTypes.Connection{
		Client: &alaTypes.Client{
			Eth: &ethclient.Client{},
			Ks:  &alaTypes.Keystore{},
		},
		Tx: &alaTypes.TxClient{
			Opts: &bind.TransactOpts{},
		},
		Contracts: &alaTypes.Contracts{
			Instances: &alaTypes.Instances{},
			Addresses: &alaTypes.Addresses{},
		},
		Network:   &alaTypes.Network{},
		Connected: false,
	}

	if args.Keystore != nil {
		err := SetKeystore(conn, args.Keystore)
		if err != nil {
			return nil, err
		}
	}

	// Only connects to the network if NodeUrl is set
	if args.NodeUrl != "" {
		ConnectToNetwork(conn, args.NodeUrl)

		err := setNetwork(conn, args)
		if err != nil {
			return nil, err
		}

		if args.Keystore != nil {
			err = setOpts(conn, args)
			if err != nil {
				return nil, err
			}
		}

		setContracts(conn, args)
	}

	return conn, nil
}

// Configuration of all the tx that the client will make
func setOpts(conn *alaTypes.Connection, args *alaTypes.ClientConf) error {
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
func setNetwork(conn *alaTypes.Connection, args *alaTypes.ClientConf) error {
	networkId, err := network.NetworkID(conn.Client.Eth)
	if err != nil {
		return err
	}
	conn.Network = &alaTypes.Network{
		Id: networkId,
	}
	return nil
}

// Instances of the contracts. This is the only way to interact with the contracts.
// They will be initialized before the first interaction with the contract and the
// instance of the contract will be cached.
// The user using args can set the address of the contract to interact with. If no
// address is set, the default address of the contract will be used
func setContracts(conn *alaTypes.Connection, args *alaTypes.ClientConf) {
	var identityManager, publicKeyRegistry, credentialRegistry common.Address
	if args.ContractAddresses.IdentityManager != (common.Address{}) {
		identityManager = args.ContractAddresses.IdentityManager
	} else {
		identityManager = configuration.AlastriaIdentityManager
	}
	if args.ContractAddresses.PublicKeyRegistry != (common.Address{}) {
		publicKeyRegistry = args.ContractAddresses.PublicKeyRegistry
	} else {
		publicKeyRegistry = configuration.PublicKeyRegistry
	}
	if args.ContractAddresses.PublicKeyRegistry != (common.Address{}) {
		credentialRegistry = args.ContractAddresses.CredentialRegistry
	} else {
		credentialRegistry = configuration.CredentialRegistry
	}

	conn.Contracts = &alaTypes.Contracts{
		Instances: &alaTypes.Instances{},
		Addresses: &alaTypes.Addresses{
			IdentityManager:    identityManager,
			PublicKeyRegistry:  publicKeyRegistry,
			CredentialRegistry: credentialRegistry,
		},
	}
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

// Connects to the network
func ConnectToNetwork(conn *alaTypes.Connection, nodeUrl string) error {
	client, err := network.ConnectToNetwork(nodeUrl)
	if err != nil {
		return err
	}
	conn.Client.Eth = client
	conn.Connected = true
	return nil
}
