package credentials

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/blockchain/network"
	"github.com/onmax/go-alastria/blockchain/tx"
	"github.com/onmax/go-alastria/blockchain/txutil"
	"github.com/onmax/go-alastria/hash"
	"github.com/onmax/go-alastria/hex"
	alaTypes "github.com/onmax/go-alastria/types"
)

func AddSubjectCredential(conn *alaTypes.Connection, signedCredential string, subjectDid *alaTypes.Did, URI string) (*types.Transaction, string, error) {
	// TODO move this logic
	psmHash, psmHashByteArr := hash.PsmHash(signedCredential, subjectDid)
	tx, err := tx.AddSubjectCredential(conn, psmHashByteArr, URI)
	if err != nil {
		return &types.Transaction{}, "", err
	}
	return tx, psmHash, nil
}

// TODO: URI right now is just one string, users might want to one URI for each credentials
func AddSubjectCredentials(client *alaTypes.Connection, signedCredentials []string, subjectDid *alaTypes.Did, URI string) ([]string, []string, error) {
	var (
		txs       []string = make([]string, len(signedCredentials))
		psmHashes []string = make([]string, len(signedCredentials))
	)

	for _, credential := range signedCredentials {
		// Right now, the function is blocking until the transactions are made
		// TODO Add non-blocking option
		tx, psmHash, err := AddSubjectCredential(client, credential, subjectDid, URI)
		if err != nil {
			return nil, nil, err
		}
		network.SendTx(client, tx)
		txutil.UpdateNonce(client)

		txs = append(txs, tx.Hash().Hex())
		psmHashes = append(psmHashes, psmHash)
	}

	return txs, psmHashes, nil
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subject common.Address) ([]common.Address, error) {
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

func GetSubjectCredentialStatus(conn *alaTypes.Connection, subject common.Address, psmHash [32]byte) (*alaTypes.PSMHashStatus, error) {
	exists, status, err := tx.GetSubjectCredentialStatus(conn, subject, psmHash)
	if err != nil {
		return nil, err
	}
	return &alaTypes.PSMHashStatus{
		Exists:  exists,
		PSMHash: common.BytesToAddress(psmHash[:]),
		Status:  status,
	}, nil
}

func GetSubjectCredentialsStatus(conn *alaTypes.Connection, subjectDid *alaTypes.Did, psmHashes []common.Address) ([]*alaTypes.PSMHashStatus, error) {
	var statuses []*alaTypes.PSMHashStatus
	for _, psmHash := range psmHashes {
		psmHashByteArr := hex.StringTo32ByteArray(psmHash.Hex())
		subjectDidAsAddress := common.HexToAddress(subjectDid.ProxyAddress)
		status, err := GetSubjectCredentialStatus(conn, subjectDidAsAddress, psmHashByteArr)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}
