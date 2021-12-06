package contracts

import (
	"errors"

	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/internal/addresses"
)

func IdentityManagerContract(client *ethclient.Client) (*identity.AlastriaContracts, error) {
	if client == nil {
		return nil, errors.New("no client specified")
	}
	return identity.NewAlastriaContracts(common.HexToAddress(addresses.AlastriaIdentityManager), client)
}

func PublicKeyRegistryContract(client *ethclient.Client) (*pkr.AlastriaContracts, error) {
	if client == nil {
		return nil, errors.New("no client specified")
	}
	return pkr.NewAlastriaContracts(common.HexToAddress(addresses.PublicKeyRegistry), client)
}
