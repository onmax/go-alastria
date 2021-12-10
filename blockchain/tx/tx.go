package tx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/blockchain/txsender"
	alaTypes "github.com/onmax/go-alastria/types"
)

func CreateAlastriaIdentity(conn *alaTypes.Connection) (*types.Transaction, error) {
	addKeyTx, err := txsender.AddKey(conn)
	if err != nil {
		return &types.Transaction{}, err
	}
	return txsender.CreateAlastriaIdentity(conn, addKeyTx.Data())

}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorAddress common.Address) (*types.Transaction, error) {
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

func AddSubjectCredential(conn *alaTypes.Connection, psmHash [32]byte, URI string) (*types.Transaction, error) {
	return txsender.AddSubjectCredential(conn, psmHash, URI)
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subject common.Address) (*big.Int, [][32]byte, error) {
	return txsender.GetSubjectCredentialList(conn, subject)
}

func AddIssuerCredential(conn *alaTypes.Connection, psmHash [32]byte) (*types.Transaction, error) {
	return txsender.AddIssuerCredential(conn, psmHash)
}
