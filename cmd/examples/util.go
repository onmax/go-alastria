package cmd

import (
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/types"
)

func GetClientConf(ksPath string) *types.ClientConf {
	return &types.ClientConf{
		Network: &types.Network{
			NodeUrl:   configuration.NodeUrl,
			Network:   configuration.Network,
			NetworkId: configuration.NetworkId,
		},
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
		Network: &types.Network{
			NodeUrl:   configuration.NodeUrl,
			Network:   configuration.Network,
			NetworkId: configuration.NetworkId,
		},
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
