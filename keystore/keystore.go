package keystore

import (
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onmax/go-alastria/types"
)

// Creates a new keystore and saves it in the path with the given password
func CreateKs(path string, password string) (*types.Keystore, error) {
	if _, err := os.Stat(path); err != nil {
		os.Mkdir(path, 0700)
	}
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	data, err := ks.NewAccount(password)
	if err != nil {
		return nil, err
	}
	return ImportKs(data.URL.Path, password)
}

// getKey get a key from KeyStore
func ImportKs(path string, password string) (*types.Keystore, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, err
	}

	return &types.Keystore{
		HexPublicKey:  hex.EncodeToString(crypto.FromECDSAPub(&key.PrivateKey.PublicKey))[2:],
		HexPrivateKey: hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)),
		PrivateKey:    key.PrivateKey,
		PublicKey:     &key.PrivateKey.PublicKey,
		Account:       key,
	}, nil
}
