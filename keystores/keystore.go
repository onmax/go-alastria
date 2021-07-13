package keystore

import (
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

// getKey get a key from KeyStore
func ImportKs(file string, password string) (string, string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", "", err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return "", "", err
	}

	privKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	pubKey := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()

	return privKey, pubKey, nil
}
