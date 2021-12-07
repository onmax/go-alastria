package configuration

import "github.com/ethereum/go-ethereum/common"

//
// Configuration for network package
//
// Node deployed and mantained by Alastria
var NodeUrl = "http://63.33.206.111/rpc"

var (
	Network   = "quor"
	NetworkId = "redT"
)

// Well-known contracts. These are the contracts that are used mostly by testing purposes.
// Users might want to use their own contracts as this contracts have been deployed by
// alastria, and alastria-identity-example use them as well.
// Upgradeable contracts
var (
	AlastriaIdentityManager = common.HexToAddress("0xeafCAf8c3f9016B14ec65105677658f3D6Eb9079")
	PublicKeyRegistry       = common.HexToAddress("0x1cd13Dc5161Bf2d5A83AAB4e7b66A6542D10Ab5B")
)

// Non-upgradeable contracts
// var (
// 	AlastriaIdentityManager = common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a")
// 	PublicKeyRegistry       = common.HexToAddress("0x017B0F93d1F31666Ae0bfb35f2e35bC1b0AC43C1")
// )

//
// End of configuration for network package
//

// Configuration for tokens
var (
	Kid         = "did:ala:quor:redt:d2f868f056ef3a48bbc8d446dfed411e9bf93ab0#keys-1"
	DidEntity   = "did:ala:quor:redT:d2f868f056ef3a48bbc8d446dfed411e9bf93ab0"
	ProviderURL = "https://your-provider-url.com"
	CallbackUrl = "https://your-backend-url.com"
	Exp         = uint64(2036798552159)
	Iat         = uint64(1636798552159)
	Jti         = "your-unique-at-id"

	TypesAT12     = []string{"AlastriaToken", "US12"}
	TypesAIC12    = []string{"AlastriaIdentityCreation", "US12"}
	ContextsAIC12 = []string{"https://alastria.github.io/identity/artifacts/v1"}
)
