package network

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"
)

// var alastriaIdentityManager = "0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"
var alastriaIdentityManager = "0xeafCAf8c3f9016B14ec65105677658f3D6Eb9079" // Upgradeable contract
// var publicKeyRegistry = "0x017B0F93d1F31666Ae0bfb35f2e35bC1b0AC43C1"
var publicKeyRegistry = "0x1cd13Dc5161Bf2d5A83AAB4e7b66A6542D10Ab5B" // Upgradeable contract

func ConnectToNetwork(node string) (*ethclient.Client, error) {
	return ethclient.Dial("http://34.91.211.67:22000")
}

func NetworkID(client *ethclient.Client) (*big.Int, error) {
	return client.NetworkID(context.Background())
}

func IdentityManagerContract(client *ethclient.Client) (*identity.AlastriaContracts, error) {
	return identity.NewAlastriaContracts(common.HexToAddress(alastriaIdentityManager), client)
}

func PublicKeyRegistryContract(client *ethclient.Client) (*pkr.AlastriaContracts, error) {
	return pkr.NewAlastriaContracts(common.HexToAddress(publicKeyRegistry), client)
}
