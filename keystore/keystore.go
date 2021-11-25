package keystore

import (
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type Keypair struct {
	hex_public_key  string
	hex_private_key string
	private_key     *ecdsa.PrivateKey
	public_key      *ecdsa.PublicKey
}

type Keystore struct {
	account  *keystore.Key
	key_pair *Keypair
}

func getKeyPair(key *keystore.Key) *Keypair {
	publicKey := hex.EncodeToString(crypto.FromECDSAPub(&key.PrivateKey.PublicKey))[2:]
	return &Keypair{
		hex_public_key:  publicKey,
		hex_private_key: hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)),
		private_key:     key.PrivateKey,
		public_key:      &key.PrivateKey.PublicKey,
	}
}

// Creates a new keystore and saves it in the path with the given password
func CreateKs(path string, password string) (*Keystore, error) {
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
func ImportKs(path string, password string) (*Keystore, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, err
	}

	return &Keystore{
		key_pair: getKeyPair(key),
		account:  key,
	}, nil
}
