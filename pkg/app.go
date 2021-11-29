package app

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/pkg/keystore"

	identity "github.com/onmax/go-alastria/pkg/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/pkg/contracts/alastria-public-key-registry"
)

type AlastriaContracts struct {
	identity_manager    *identity.AlastriaContracts
	public_key_registry *pkr.AlastriaContracts
}

type Person struct {
	name string
	age  int
}

type AlastriaNetwork struct {
	id       *big.Int
	node_url string
}

type AlastriaClient struct {
	eth *ethclient.Client
	ks  *keystore.AlastriaKeystore
}

type AlastriaConnection struct {
	contracts *AlastriaContracts
	network   *AlastriaNetwork
	client    *AlastriaClient
}
