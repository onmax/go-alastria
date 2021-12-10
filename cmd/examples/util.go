package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/onmax/go-alastria/alastria"
	"github.com/onmax/go-alastria/did"
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/types"
)

func GetClientConf(ksPath string) *types.ClientConf {
	return &types.ClientConf{
		NodeUrl: configuration.NodeUrl,
		Keystore: &types.KeystoreConfig{
			Path:     ksPath,
			Password: "Passw0rd",
		},
		ContractAddresses: &types.Addresses{
			IdentityManager:    configuration.AlastriaIdentityManager,
			PublicKeyRegistry:  configuration.PublicKeyRegistry,
			CredentialRegistry: configuration.CredentialRegistry,
		},
	}
}

func GetReaderClientConf() *types.ClientConf {
	return &types.ClientConf{
		NodeUrl: configuration.NodeUrl,
		ContractAddresses: &types.Addresses{
			IdentityManager:    configuration.AlastriaIdentityManager,
			PublicKeyRegistry:  configuration.PublicKeyRegistry,
			CredentialRegistry: configuration.CredentialRegistry,
		},
	}
}

func GetDisconnectedClientConf(ksPath string) *types.ClientConf {
	return &types.ClientConf{
		Keystore: &types.KeystoreConfig{
			Path:     ksPath,
			Password: "Passw0rd",
		},
	}
}

func GetDIDGivenAddress(address common.Address) *types.Did {
	// Any member can connect to the network to execute this function, in this case will be the entity
	entityArgs := GetReaderClientConf()
	entityClient, _ := alastria.NewClient(entityArgs)

	actorProxy, _ := alastria.IdentityKeys(entityClient, address)

	did := did.NewDid(configuration.Network, configuration.NetworkId, actorProxy)

	return did
}
