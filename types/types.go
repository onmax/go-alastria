package types

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"
)

type Keystore struct {
	Account       *keystore.Key
	HexPublicKey  string
	HexPrivateKey string
	PrivateKey    *ecdsa.PrivateKey
	PublicKey     *ecdsa.PublicKey
}

type Contracts struct {
	IdentityManager   *identity.AlastriaContracts
	PublicKeyRegistry *pkr.AlastriaContracts
}

type Network struct {
	Id *big.Int
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
type ConnectionArgs struct {
	NodeUrl  string
	Keystore *KeystoreConfig
}

var (
	ErrNodeUrlNotSet   = errors.New("node url not set")
	ErrKeystoreNotSet  = errors.New("keystore not set")
	ErrEthClientNotSet = errors.New("eth client not set")
)
