package tx

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/onmax/go-alastria/contracts"
	alaTypes "github.com/onmax/go-alastria/types"
)

func checkTxOpts(conn *alaTypes.Connection) error {
	if conn.Tx.Opts == nil {
		return alaTypes.ErrTxOptsNotSet
	}
	return nil
}

func checkKeystore(conn *alaTypes.Connection) error {
	if conn.Client.Ks == nil {
		return alaTypes.ErrKeystoreNotSet
	}
	return nil
}

func checkIdentityManager(conn *alaTypes.Connection) error {
	if conn.Contracts.Instances.IdentityManager == nil {
		if conn.Contracts.Addresses.IdentityManager == (common.Address{}) {
			return alaTypes.ErrAddressNotSet
		}
		instance, err := contracts.IdentityManagerContract(conn.Client.Eth, conn.Contracts.Addresses.IdentityManager)
		if err != nil {
			return err
		}
		conn.Contracts.Instances.IdentityManager = instance
	}
	return nil
}

func checkPublickeyRegistry(conn *alaTypes.Connection) error {
	if conn.Contracts.Instances.PublicKeyRegistry == nil {
		if conn.Contracts.Addresses.PublicKeyRegistry == (common.Address{}) {
			return alaTypes.ErrAddressNotSet
		}
		instance, err := contracts.PublicKeyRegistryContract(conn.Client.Eth, conn.Contracts.Addresses.PublicKeyRegistry)
		if err != nil {
			return err
		}
		conn.Contracts.Instances.PublicKeyRegistry = instance
	}
	return nil
}

func checkCredentialRegistry(conn *alaTypes.Connection) error {
	if conn.Contracts.Instances.CredentialRegistry == nil {
		if conn.Contracts.Addresses.CredentialRegistry == (common.Address{}) {
			return alaTypes.ErrAddressNotSet
		}
		instance, err := contracts.CredentialRegistryContract(conn.Client.Eth, conn.Contracts.Addresses.CredentialRegistry)
		if err != nil {
			return err
		}
		conn.Contracts.Instances.CredentialRegistry = instance
	}
	return nil
}
