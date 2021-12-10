package txsender

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	alaTypes "github.com/onmax/go-alastria/types"
)

func AddSubjectCredential(conn *alaTypes.Connection, psmHash [32]byte, URI string) (*types.Transaction, error) {
	err := checkCredentialRegistry(conn)
	if err != nil {
		return &types.Transaction{}, err
	}
	return conn.Contracts.Instances.CredentialRegistry.AddSubjectCredential(conn.Tx.Opts, psmHash, URI)
}

func GetSubjectCredentialList(conn *alaTypes.Connection, subject common.Address) (*big.Int, [][32]byte, error) {
	err := checkCredentialRegistry(conn)
	if err != nil {
		return &big.Int{}, [][32]byte{}, err
	}
	return conn.Contracts.Instances.CredentialRegistry.GetSubjectCredentialList(&bind.CallOpts{}, subject)
}

func AddIssuerCredential(conn *alaTypes.Connection, psmHash [32]byte) (*types.Transaction, error) {
	err := checkCredentialRegistry(conn)
	if err != nil {
		return &types.Transaction{}, err
	}
	return conn.Contracts.Instances.CredentialRegistry.AddIssuerCredential(conn.Tx.Opts, psmHash)
}
