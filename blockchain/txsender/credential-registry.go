package txsender

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/client"
	alaTypes "github.com/onmax/go-alastria/types"
)

// Subject

func AddSubjectCredential(conn *alaTypes.Connection, psmHash [32]byte, URI string) (*types.Transaction, error) {
	err := client.CheckCredentialRegistry(conn)
	if err != nil {
		return &types.Transaction{}, err
	}
	return conn.Contracts.Instances.CredentialRegistry.AddSubjectCredential(conn.Tx.Opts, psmHash, URI)
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subject common.Address) (*big.Int, [][32]byte, error) {
	err := client.CheckCredentialRegistry(conn)
	if err != nil {
		return &big.Int{}, [][32]byte{}, err
	}
	return conn.Contracts.Instances.CredentialRegistry.GetSubjectCredentialList(&bind.CallOpts{}, subject)
}

func GetSubjectCredentialStatus(conn *alaTypes.Connection, subject common.Address, psmHash [32]byte) (bool, uint8, error) {
	err := client.CheckCredentialRegistry(conn)
	if err != nil {
		return false, 0, err
	}
	res, err := conn.Contracts.Instances.CredentialRegistry.GetSubjectCredentialStatus(&bind.CallOpts{}, subject, psmHash)
	if err != nil {
		return false, 0, err
	}
	return res.Exists, res.Status, nil
}

// Issuer

func AddIssuerCredential(conn *alaTypes.Connection, psmHash [32]byte) (*types.Transaction, error) {
	err := client.CheckCredentialRegistry(conn)
	if err != nil {
		return &types.Transaction{}, err
	}
	return conn.Contracts.Instances.CredentialRegistry.AddIssuerCredential(conn.Tx.Opts, psmHash)
}

func GetIssuerCredentialStatus(conn *alaTypes.Connection, issuer common.Address, psmHash [32]byte) (bool, uint8, error) {
	err := client.CheckCredentialRegistry(conn)
	if err != nil {
		return false, 0, err
	}
	res, err := conn.Contracts.Instances.CredentialRegistry.GetIssuerCredentialStatus(&bind.CallOpts{}, issuer, psmHash)
	if err != nil {
		return false, 0, err
	}
	return res.Exists, res.Status, nil
}
