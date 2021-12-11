package alastria

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/blockchain/network"
	"github.com/onmax/go-alastria/blockchain/tx"
	"github.com/onmax/go-alastria/blockchain/txutil"
	"github.com/onmax/go-alastria/client"
	"github.com/onmax/go-alastria/crypto"
	"github.com/onmax/go-alastria/did"
	"github.com/onmax/go-alastria/hash"
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
	if conn.Client.Eth == nil {
		return alaTypes.ErrEthClientNotSet
	}
	if conn.Client.Ks == nil {
		return alaTypes.ErrKeystoreNotSet
	}
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

func AddSubjectCredential(conn *alaTypes.Connection, signedJWT string, subjectDid *alaTypes.Did, URI string) (*types.Transaction, string, error) {
	// TODO move this logic
	psmHash, psmHashByteArr := hash.PsmHash(signedJWT, subjectDid)
	tx, err := tx.AddSubjectCredential(conn, psmHashByteArr, URI)
	if err != nil {
		return &types.Transaction{}, "", err
	}
	return tx, psmHash, nil
}

func AddSubjectCredentials(client *alaTypes.Connection, credentials []string, subjectDid *alaTypes.Did, URI string) ([]string, []string, error) {
	// TODO Move this code
	var (
		txs       []string = make([]string, len(credentials))
		psmHashes []string = make([]string, len(credentials))
	)

	for _, credential := range credentials {
		// Right now, the function is blocking until the transactions are made
		// TODO Add non-blocking option
		tx, psmHash, err := AddSubjectCredential(client, credential, subjectDid, URI)
		if err != nil {
			return nil, nil, err
		}
		SendTx(client, tx)
		txutil.UpdateNonce(client)

		txs = append(txs, tx.Hash().Hex())
		psmHashes = append(psmHashes, psmHash)
	}

	return txs, psmHashes, nil
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subject common.Address) ([]common.Address, error) {
	// TODO Move somewhere else
	// Ignoringing the first value as it is the length of credentialsByteArray and in go is easy to calculate
	_, credentialsByteArray, err := tx.GetSubjectCredentialList(conn, subject)
	if err != nil {
		return []common.Address{}, err
	}

	var credentials []common.Address
	for _, credentialByteArray := range credentialsByteArray {
		credential := common.BytesToAddress(credentialByteArray[:])
		credentials = append(credentials, credential)
	}
	return credentials, nil
}

func AddIssuerCredential(conn *alaTypes.Connection, signedJWT string, entityDid *alaTypes.Did) (*types.Transaction, string, error) {
	// TODO Move somewhere else
	psmHash, psmHashByteArr := hash.PsmHash(signedJWT, entityDid)
	addSubjectCredentialTx, err := tx.AddIssuerCredential(conn, psmHashByteArr)
	if err != nil {
		return &types.Transaction{}, "", err
	}
	return addSubjectCredentialTx, psmHash, nil
}

func AddIssuerCredentials(client *alaTypes.Connection, credentials []string, entityDid *alaTypes.Did) ([]string, []string, error) {
	// TODO Move this code
	var (
		txs       []string = make([]string, len(credentials))
		psmHashes []string = make([]string, len(credentials))
	)

	for _, credential := range credentials {
		// Right now, the function is blocking until the transactions are made
		// TODO Add non-blocking option
		tx, psmHash, err := AddIssuerCredential(client, credential, entityDid)
		if err != nil {
			return nil, nil, err
		}
		SendTx(client, tx)
		txutil.UpdateNonce(client)
		txs = append(txs, tx.Hash().Hex())
		psmHashes = append(psmHashes, psmHash)
	}

	return txs, psmHashes, nil
}

func SignJWT(conn *alaTypes.Connection, jwt interface{}) (string, error) {
	// TODO Check Keystore
	return crypto.Sign(jwt, conn.Client.Ks.HexPrivateKey)
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
