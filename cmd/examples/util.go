package cmd

import (
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
			IdentityManager:   configuration.AlastriaIdentityManager,
			PublicKeyRegistry: configuration.PublicKeyRegistry,
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
