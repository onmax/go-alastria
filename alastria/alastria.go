package alastria

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/client"
	"github.com/onmax/go-alastria/crypto"
	"github.com/onmax/go-alastria/hex"
	"github.com/onmax/go-alastria/network"
	"github.com/onmax/go-alastria/tx"
	alaTypes "github.com/onmax/go-alastria/types"
)

// Initializes the client that any actor interacting with the network should use. It can be:
// subject, issuer, or service provider.
// args.NodeUrl is mandatory
// args.Keystore is not mandatory, but it is required if you want to sign JWT or tx
func NewClient(args *alaTypes.ClientConf) (*alaTypes.Connection, error) {
	return client.NewClient(args)
}

// Set the keystore that will be used to sign the transactions and JWTs
func SetKeystore(conn *alaTypes.Connection, ksConfig *alaTypes.KeystoreConfig) error {
	return client.SetKeystore(conn, ksConfig)
}

// Sends a transaction to the network. It will wait until the transaction is mined blocking
// the current thread checking once every second.
// conn needs to have a client and keystore set
func SendTx(conn *alaTypes.Connection, tx *types.Transaction) error {
	if conn.Client.Eth == nil {
		return alaTypes.ErrEthClientNotSet
	}
	if conn.Client.Ks == nil {
		return alaTypes.ErrKeystoreNotSet
	}
	return network.SendTx(conn, tx)
}

func PublicKeyToAddress(publicKey string) (common.Address, error) {
	return crypto.PublicKeyToAddress(publicKey)
}

func HexToTx(txStr string) (*types.Transaction, error) {
	return tx.HexToTx(txStr)
}

func TxToHex(tx_ *types.Transaction) (string, error) {
	return tx.TxToHex(tx_)
}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorPublicKey string) (*types.Transaction, error) {
	newActorAddress, err := crypto.PublicKeyToAddress(newActorPublicKey)
	if err != nil {
		return nil, err
	}
	return tx.PrepareAlastriaId(conn, newActorAddress)
}

func CreateAlastriaIdentity(conn *alaTypes.Connection) (*types.Transaction, error) {
	return tx.CreateAlastriaIdentity(conn)
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (string, error) {
	proxyAddress, err := tx.IdentityKeys(conn, agentAddress)
	if err != nil {
		return "", err
	}

	return hex.Remove0x(proxyAddress.Hash().Hex())[24:], nil
}
