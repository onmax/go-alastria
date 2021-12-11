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

func AddIssuerCredential(conn *alaTypes.Connection, signedJWT string, entityDid *alaTypes.Did) (*types.Transaction, string, error) {
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
		network.SendTx(client, tx)
		txutil.UpdateNonce(client)
		txs = append(txs, tx.Hash().Hex())
		psmHashes = append(psmHashes, psmHash)
	}

	return txs, psmHashes, nil
}

func GetIssuerCredentialStatus(conn *alaTypes.Connection, subjectDid *alaTypes.Did, psmHash common.Address) (*alaTypes.PSMHashStatus, error) {
	psmHashByteArr := hex.StringTo32ByteArray(psmHash.Hex())
	subjectDidAsAddress := common.HexToAddress(subjectDid.ProxyAddress)
	return GetSubjectCredentialStatus(conn, subjectDidAsAddress, psmHashByteArr)
}
