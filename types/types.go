package types

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	credential "github.com/onmax/go-alastria/contracts/alastria-credential-registry"
	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"
)

type Header struct {
	Algorithm    string `json:"alg,omitempty"`
	Type         string `json:"typ,omitempty"`
	JSONWebToken string `json:"jwk,omitempty"`
	KeyID        string `json:"kid,omitempty"`
}

type Keystore struct {
	Account       *keystore.Key
	HexPublicKey  string
	HexPrivateKey string
	PrivateKey    *ecdsa.PrivateKey
	PublicKey     *ecdsa.PublicKey
}

type Instances struct {
	IdentityManager    *identity.AlastriaContracts
	PublicKeyRegistry  *pkr.AlastriaContracts
	CredentialRegistry *credential.AlastriaContracts
}
type Addresses struct {
	IdentityManager    common.Address
	PublicKeyRegistry  common.Address
	CredentialRegistry common.Address
}

type Contracts struct {
	Instances *Instances
	Addresses *Addresses
}

type Network struct {
	Id        *big.Int
	NodeUrl   string
	Network   string
	NetworkId string
}

type Client struct {
	Eth *ethclient.Client
	Ks  *Keystore
}

type TxClient struct {
	Opts   *bind.TransactOpts
	Signer types.Signer
}

type Connection struct {
	Contracts *Contracts
	Network   *Network
	Client    *Client
	Tx        *TxClient
}

type KeystoreConfig struct {
	Path     string
	Password string
}
type ClientConf struct {
	Keystore          *KeystoreConfig
	ContractAddresses *Addresses
	Network           *Network
}

type Did struct {
	Protocol     string
	Network      string
	NetworkId    string
	ProxyAddress string
}

func (d *Did) String() string {
	return fmt.Sprintf("did:%s:%s:%s:%s", d.Protocol, d.Network, d.NetworkId, d.ProxyAddress)
}

var (
	ErrNodeUrlNotSet   = errors.New("node url not set")
	ErrKeystoreNotSet  = errors.New("keystore not set")
	ErrAddressNotSet   = errors.New("address not set")
	ErrEthClientNotSet = errors.New("eth client not set")
	ErrTxOptsNotSet    = errors.New("eth tx opts not set")
	ErrInvalidJWT      = errors.New("invalid jwt")
)

type PSMHashStatus struct {
	Exists  bool
	PSMHash string
	Status  uint8
}

type PSMHashStatuses struct {
	IssuerPSMHash  string
	SubjectPSMHash string
	Credential     string
	Status         uint8
}

// TODO Separate the types in files
