package credentials

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/blockchain/network"
	"github.com/onmax/go-alastria/blockchain/tx"
	"github.com/onmax/go-alastria/blockchain/txutil"
	"github.com/onmax/go-alastria/did"
	"github.com/onmax/go-alastria/hash"
	"github.com/onmax/go-alastria/tokens"
	alaTypes "github.com/onmax/go-alastria/types"
)

func AddIssuerCredential(conn *alaTypes.Connection, signedJWT string, entityDid *alaTypes.Did) (*types.Transaction, string, error) {
	psmHash := hash.PsmHash(signedJWT, entityDid)
	addSubjectCredentialTx, err := tx.AddIssuerCredential(conn, psmHash)
	if err != nil {
		return &types.Transaction{}, "", err
	}

	return addSubjectCredentialTx, psmHash, nil
}

func AddIssuerCredentials(client *alaTypes.Connection, credentials []string, entityDid *alaTypes.Did) ([]string, []string, error) {
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

func GetIssuerCredentialStatus(conn *alaTypes.Connection, signedCredential string) (*alaTypes.PSMHashStatus, error) {
	credential, err := tokens.DecodeCredential(signedCredential)
	if err != nil {
		return nil, err
	}
	entityDid, err := did.NewDidFromString(credential.Payload.Issuer)
	if err != nil {
		return nil, err
	}

	issuerPsmHash := hash.PsmHash(signedCredential, entityDid)
	_, status, err := tx.GetIssuerCredentialStatus(conn, entityDid, issuerPsmHash)
	if err != nil {
		return &alaTypes.PSMHashStatus{}, err
	}

	return &alaTypes.PSMHashStatus{
		PSMHash: issuerPsmHash,
		Status:  status,
	}, nil
}
