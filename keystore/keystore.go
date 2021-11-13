package keystore

import (
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// Creates a new keystore and saves it in the path with the given password
func CreateKs(folder string, password string) (accounts.Account, error) {
	if _, err := os.Stat(folder); err == nil {
		return accounts.Account{}, err
	}
	ks := keystore.NewKeyStore(folder, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return accounts.Account{}, err
	}

	return account, nil
}

func PublicKeyToString(publicKey *ecdsa.PublicKey) string {
	return hex.EncodeToString(crypto.FromECDSAPub(publicKey))[2:]
}

// getKey get a key from KeyStore
func ImportKs(path string, password string) (*keystore.Key, string, string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, "", "", err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, "", "", err
	}

	if err != nil {
		return nil, "", "", err
	}

	privKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	pubKey := PublicKeyToString(&key.PrivateKey.PublicKey)

	return key, privKey, pubKey, nil
}
