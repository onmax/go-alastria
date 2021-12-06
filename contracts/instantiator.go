package contracts

import (
	"errors"

	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IdentityManagerContract(client *ethclient.Client, contract common.Address) (*identity.AlastriaContracts, error) {
	if client == nil {
		return nil, errors.New("no client specified")
	}
	return identity.NewAlastriaContracts(contract, client)
}

func PublicKeyRegistryContract(client *ethclient.Client, contract common.Address) (*pkr.AlastriaContracts, error) {
	if client == nil {
		return nil, errors.New("no client specified")
	}
	return pkr.NewAlastriaContracts(contract, client)
}
