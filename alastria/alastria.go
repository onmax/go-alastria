package alastria

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/blockchain/network"
	"github.com/onmax/go-alastria/blockchain/tx"
	"github.com/onmax/go-alastria/blockchain/txutil"
	"github.com/onmax/go-alastria/client"
	"github.com/onmax/go-alastria/credentials"
	"github.com/onmax/go-alastria/crypto"
	"github.com/onmax/go-alastria/did"
	"github.com/onmax/go-alastria/hex"
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
	return network.SendTx(conn, tx)
}

func PublicKeyToAddress(publicKey string) (common.Address, error) {
	return hex.PublicKeyToAddress(publicKey)
}

func HexToTx(txStr string) (*types.Transaction, error) {
	return txutil.HexToTx(txStr)
}

func TxToHex(tx_ *types.Transaction) (string, error) {
	return txutil.TxToHex(tx_)
}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorPublicKey string) (*types.Transaction, error) {
	return tx.PrepareAlastriaId(conn, newActorPublicKey)
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

func GetCurrentPublicKey(conn *alaTypes.Connection, agentAddress common.Address) (string, error) {
	return tx.GetCurrentPublicKey(conn, agentAddress)
}

func AddSubjectCredentials(client *alaTypes.Connection, signedCredentials []string, subjectDid *alaTypes.Did, URI string) ([]string, []string, error) {
	// TODO Return also the status of the psmHash
	return credentials.AddSubjectCredentials(client, signedCredentials, subjectDid, URI)
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subjectAddress common.Address) ([]common.Address, error) {
	return credentials.GetSubjectCredentialList(conn, subjectAddress)
}

func GetSubjectCredentialsStatus(conn *alaTypes.Connection, subjectDid *alaTypes.Did, psmHashes []common.Address) ([]*alaTypes.PSMHashStatus, error) {
	return credentials.GetSubjectCredentialsStatus(conn, subjectDid, psmHashes)
}

func AddIssuerCredentials(client *alaTypes.Connection, signedCredentials []string, entityDid *alaTypes.Did) ([]string, []string, error) {
	// TODO Return also the status of the psmHash
	return credentials.AddIssuerCredentials(client, signedCredentials, entityDid)
}

func GetIssuerCredentialStatus(conn *alaTypes.Connection, issuerDid *alaTypes.Did, psmHash common.Address) (*alaTypes.PSMHashStatus, error) {
	return credentials.GetIssuerCredentialStatus(conn, issuerDid, psmHash)
}

func SignJWT(conn *alaTypes.Connection, jwt interface{}) (string, error) {
	return crypto.Sign(conn, jwt)
}

func VerifyJWT(signedJWT string, publicKey string) (bool, error) {
	e := crypto.Verify(signedJWT, publicKey)
	return e == nil, e
}

func NewDid(network, networkId, proxyAddress string) *alaTypes.Did {
	return did.NewDid(network, networkId, proxyAddress)
}

func NewDidFromString(didStr string) (*alaTypes.Did, error) {
	return did.NewDidFromString(didStr)
}
