package network

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"
	"github.com/onmax/go-alastria/keystore"
)

type AlastriaContracts struct {
	identity_manager    *identity.AlastriaContracts
	public_key_registry *pkr.AlastriaContracts
}

type AlastriaNetwork struct {
	id       *big.Int
	node_url string
}

type AlastriaClient struct {
	eth *ethclient.Client
	ks  *keystore.Keystore
}

type AlastriaConnection struct {
	contracts *AlastriaContracts
	network   *AlastriaNetwork
	client    *AlastriaClient
}

func New(node string) (*AlastriaConnection, error) {
	client, err := ConnectToNetwork(node)
	if err != nil {
		return nil, err
	}
	return &AlastriaConnection{
		client: &AlastriaClient{
			eth: client,
		},
		contracts: &AlastriaContracts{},
		network:   &AlastriaNetwork{},
	}, nil
}
