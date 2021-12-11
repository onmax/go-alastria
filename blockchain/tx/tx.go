package tx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/blockchain/txsender"
	"github.com/onmax/go-alastria/hex"
	alaTypes "github.com/onmax/go-alastria/types"
)

func CreateAlastriaIdentity(conn *alaTypes.Connection) (*types.Transaction, error) {
	addKeyTx, err := txsender.AddKey(conn)
	if err != nil {
		return &types.Transaction{}, err
	}
	return txsender.CreateAlastriaIdentity(conn, addKeyTx.Data())
}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorPublicKey string) (*types.Transaction, error) {
	newActorAddress, err := hex.PublicKeyToAddress(newActorPublicKey)

	if err != nil {
		return nil, err
	}
	delegatedData, err := txsender.PrepareAlastriaId(conn, newActorAddress)
	if err != nil {
		return &types.Transaction{}, err
	}
	return txsender.DelegateCall(conn, delegatedData)
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	return txsender.IdentityKeys(conn, agentAddress)
}

func GetCurrentPublicKey(conn *alaTypes.Connection, agentAddress common.Address) (string, error) {
	return txsender.GetCurrentPublicKey(conn, agentAddress)
}

// Credentials - Subject
func AddSubjectCredential(conn *alaTypes.Connection, psmHash [32]byte, URI string) (*types.Transaction, error) {
	return txsender.AddSubjectCredential(conn, psmHash, URI)
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subject common.Address) (*big.Int, [][32]byte, error) {
	return txsender.GetSubjectCredentialList(conn, subject)
}

func GetSubjectCredentialStatus(conn *alaTypes.Connection, subject common.Address, psmHash [32]byte) (bool, uint8, error) {
	return txsender.GetSubjectCredentialStatus(conn, subject, psmHash)
}

// Credentials - Issuer
func AddIssuerCredential(conn *alaTypes.Connection, psmHash [32]byte) (*types.Transaction, error) {
	return txsender.AddIssuerCredential(conn, psmHash)
}

func GetIssuerCredentialStatus(conn *alaTypes.Connection, issuer common.Address, psmHash [32]byte) (bool, uint8, error) {
	return txsender.GetIssuerCredentialStatus(conn, issuer, psmHash)
}
