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
func AddSubjectCredential(conn *alaTypes.Connection, psmHash string, URI string) (*types.Transaction, error) {
	psmHashByteArr := hex.StringTo32ByteArray(psmHash)
	return txsender.AddSubjectCredential(conn, psmHashByteArr, URI)
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subject common.Address) (*big.Int, []string, error) {
	length, psmHashByteArr, err := txsender.GetSubjectCredentialList(conn, subject)
	if err != nil {
		return nil, nil, err
	}
	var psmHashList []string

	for _, psmHashByte := range psmHashByteArr {
		psmHashList = append(psmHashList, hex.ByteArrToHex(psmHashByte[:]))
	}
	return length, psmHashList, nil
}

func GetSubjectCredentialStatus(conn *alaTypes.Connection, subject *alaTypes.Did, psmHash string) (bool, uint8, error) {
	psmHashByteArr := hex.StringTo32ByteArray(psmHash)
	return txsender.GetSubjectCredentialStatus(conn, common.HexToAddress(subject.ProxyAddress), psmHashByteArr)
}

// Credentials - Issuer
func AddIssuerCredential(conn *alaTypes.Connection, psmHash string) (*types.Transaction, error) {
	psmHashByteArr := hex.StringTo32ByteArray(psmHash)
	return txsender.AddIssuerCredential(conn, psmHashByteArr)
}

func GetIssuerCredentialStatus(conn *alaTypes.Connection, issuer *alaTypes.Did, psmHash string) (bool, uint8, error) {
	psmHashByteArr := hex.StringTo32ByteArray(psmHash)
	return txsender.GetIssuerCredentialStatus(conn, common.HexToAddress(issuer.ProxyAddress), psmHashByteArr)
}
