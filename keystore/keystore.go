package keystore

import (
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type Keypair struct {
	public_key  string
	private_key string
}

type Keystore struct {
	account  *keystore.Key
	key_pair *Keypair
}

func getKeyPair(key *keystore.Key) *Keypair {
	publicKey := hex.EncodeToString(crypto.FromECDSAPub(&key.PrivateKey.PublicKey))[2:]
	return &Keypair{
		public_key:  publicKey,
		private_key: hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)),
	}
}

// Creates a new keystore and saves it in the path with the given password
func CreateKs(folder string, password string) (*Keystore, error) {
	if _, err := os.Stat(folder); err == nil {
		return nil, err
	}
	ks := keystore.NewKeyStore(folder, keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount(password)
	if err != nil {
		return nil, err
	}
	return ImportKs(folder, password)
}

// getKey get a key from KeyStore
func ImportKs(path string, password string) (*Keystore, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &Keystore{
		key_pair: getKeyPair(key),
		account:  key,
	}, nil
}
