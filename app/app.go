package app

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/network"

	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"
)

type AlastriaContracts struct {
	IdentityManager   *identity.AlastriaContracts
	PublicKeyRegistry *pkr.AlastriaContracts
}

type Person struct {
	Name string
	Age  int
}

type AlastriaNetwork struct {
	Id *big.Int
}

type AlastriaClient struct {
	Eth *ethclient.Client
	Ks  *keystore.AlastriaKeystore
}

type AlastriaConnection struct {
	Contracts *AlastriaContracts
	Network   *AlastriaNetwork
	Client    *AlastriaClient
}

type AlastriaConnectionArgs struct {
	NodeUrl  string
	Keystore *keystore.AlastriaKeystore
}

func NewAlastriaClient(args AlastriaConnectionArgs) (*AlastriaConnection, error) {
	client, err := network.ConnectToNetwork(args.NodeUrl)
	if err != nil {
		return nil, err
	}
	networkId, err := network.NetworkID(client)
	if err != nil {
		return nil, err
	}

	return &AlastriaConnection{
		Client: &AlastriaClient{
			Eth: client,
			Ks:  args.Keystore,
		},
		Contracts: &AlastriaContracts{
			IdentityManager:   nil,
			PublicKeyRegistry: nil,
		},
		Network: &AlastriaNetwork{
			Id: networkId,
		},
	}, nil
}
