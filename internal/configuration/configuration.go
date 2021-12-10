package configuration

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

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
	CredentialRegistry      = common.HexToAddress("0x33CDa1389Ddc944A8082949Ff1c62c5F450d38d2")
	PresentationRegistry    = common.HexToAddress("0x337268ddd7C6403375531f130740004cEEf2808e")
)

// Non-upgradeable contracts
// var (
// 	AlastriaIdentityManager = common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a")
// 	PublicKeyRegistry       = common.HexToAddress("0x017B0F93d1F31666Ae0bfb35f2e35bC1b0AC43C1")
// 	CredentialRegistry      = common.HexToAddress("0x7bbca11cbd86b562136d5708eba40f4bc0aa1ddc")
// 	PresentationRegistry    = common.HexToAddress("0x54d1dbfacada17ff39f2bac08e05fbdb4659f671")
// )

//
// End of configuration for network package
//

//
// Configuration for tokens package
//
var (
	Kid         = "did:ala:quor:redt:d2f868f056ef3a48bbc8d446dfed411e9bf93ab0#keys-1"
	DidEntity   = "did:ala:quor:redT:d2f868f056ef3a48bbc8d446dfed411e9bf93ab0"
	ProviderURL = "https://your-provider-url.com"
	CallbackUrl = "https://your-backend-url.com"

	// TODO Review this CBU and add docs
	CallbackUrlCredentialIssuance = "https://your-backend-url.com/alastria/credential"
	Exp                           = uint64(2036798552159)
	Iat                           = uint64(1636798552159)
	JtiAS                         = "your-unique-at-id-only-one-use-alastria-session"
	JtiAT                         = "your-unique-at-id-only-one-use-alastria-token"
	JtiAIC                        = "your-unique-at-id-only-one-use-aic"
	JtiCredential                 = "your-unique-at-id-only-one-use-credential"
	JtiPresentation               = "your-unique-at-id-only-one-use-presentation"
	JtiPR                         = "your-unique-at-id-only-one-use-pr"

	// TODO Use human readable names
	TypesAT12  = []string{"AlastriaToken", "US12"}
	TypesAT221 = []string{"AlastriaToken", "US221"}

	TypesAIC12 = []string{"AlastriaIdentityCreation", "US12"}

	TypesASAuthentication     = []string{"AlastriaSession", "US211"}
	TypesASCredentialIssuance = []string{"AlastriaSession", "US221"}
	TypesASVinculation        = []string{"AlastriaSession", "US142"}
	ContextsAIC12             = []string{"https://alastria.github.io/identity/artifacts/v1"}

	TypesCredentials    = []string{"VerifiableCredential", "AlastriaVerifiableCredential"}
	ContextsCredentials = []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"}
)

//
// End of configuration for tokens package
//

//
// Configuration for tx package
//
var (
	TxValue    = big.NewInt(0)
	TxGasLimit = uint64(600000)
	TxGasPrice = big.NewInt(0)
	TxNoSend   = true
)

//
// End of configuration for tx package
//
