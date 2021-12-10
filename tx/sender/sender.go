package txSender

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	alaTypes "github.com/onmax/go-alastria/types"
)

// TODO Refactor somehow

func AddKey(conn *alaTypes.Connection) (*types.Transaction, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return nil, err
	}

	err = checkPublickeyRegistry(conn)
	if err != nil {
		return nil, err
	}

	err = checkTxOpts(conn)
	if err != nil {
		return nil, err
	}

	err = checkKeystore(conn)
	if err != nil {
		return nil, err
	}
	return conn.Contracts.Instances.PublicKeyRegistry.AddKey(conn.Tx.Opts, conn.Client.Ks.HexPublicKey)
}

func CreateAlastriaIdentity(conn *alaTypes.Connection, addPublicKeyCallData []byte) (*types.Transaction, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return nil, err
	}

	err = checkTxOpts(conn)
	if err != nil {
		return nil, err
	}

	err = checkKeystore(conn)
	if err != nil {
		return nil, err
	}
	
	return conn.Contracts.Instances.IdentityManager.CreateAlastriaIdentity(conn.Tx.Opts, addPublicKeyCallData)
}

func PrepareAlastriaId(conn *alaTypes.Connection, newActorAddress common.Address) (*types.Transaction, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return nil, err
	}

	err = checkTxOpts(conn)
	if err != nil {
		return nil, err
	}

	err = checkKeystore(conn)
	if err != nil {
		return nil, err
	}

	delegatedData, _ := conn.Contracts.Instances.IdentityManager.PrepareAlastriaID(conn.Tx.Opts, newActorAddress)
	delegated, _ := conn.Contracts.Instances.IdentityManager.DelegateCall(
		conn.Tx.Opts, common.HexToAddress(conn.Contracts.Addresses.IdentityManager.Hex()), big.NewInt(0), delegatedData.Data())
	return delegated, nil
}

func IdentityKeys(conn *alaTypes.Connection, agentAddress common.Address) (common.Address, error) {
	err := checkIdentityManager(conn)
	if err != nil {
		return common.Address{}, err
	}

	return conn.Contracts.Instances.IdentityManager.IdentityKeys(&bind.CallOpts{}, agentAddress)
}

func GetCurrentPublicKey(conn *alaTypes.Connection, agentAddress common.Address) (string, error) {
	err := checkPublickeyRegistry(conn)
	if err != nil {
		return "", err
	}

	return conn.Contracts.Instances.PublicKeyRegistry.GetCurrentPublicKey(&bind.CallOpts{}, agentAddress)
}

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
