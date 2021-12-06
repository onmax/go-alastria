package dev

import (
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/types"
)

func GetConnectionConf(ksPath string) *types.ConnectionConf {
	return &types.ConnectionConf{
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
