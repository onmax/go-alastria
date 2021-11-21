package tx

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/crypto"
	alastriaKs "github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/network"
)

func PrepareAlastriaId(client *ethclient.Client, entityKeystore *keystore.Key, agentPubKey string) *types.Transaction {
	fmt.Println("Starting PrepareAlastriaId")

	instance := network.IdentityManagerContract(client)
	fmt.Println("Connected to AlastriaIdentityManager contract")

	opts, _ := TxOpt(entityKeystore.PrivateKey)
	opts.NoSend = true
	tx, _ := instance.CreateAlastriaIdentity(opts, []byte(agentPubKey))
	fmt.Println("Created tx")

	signedTx := crypto.SignTx(network.NetworkID(client), tx, entityKeystore.PrivateKey)
	fmt.Println("Signed tx")

	fmt.Println("End of PrepareAlastriaId")

	return signedTx
}

func addKeyTx(client *ethclient.Client, entityPrivKey *ecdsa.PrivateKey, agentKeystore *keystore.Key) *types.Transaction {
	instance := network.PublicKeyRegistryContract(client)
	fmt.Println("Connected to PublicKeyRegistry contract")

	agentPrivKey := agentKeystore.PrivateKey

	opts, _ := TxOpt(entityPrivKey)
	opts.NoSend = true
	tx, _ := instance.AddKey(opts, alastriaKs.PublicKeyToString(&agentPrivKey.PublicKey))
	fmt.Println("Created tx")

	signedTxAddKey := crypto.SignTx(network.NetworkID(client), tx, agentPrivKey)
	fmt.Println("Signed Add Key Tx")
	return signedTxAddKey
}

func createAlastriaIdentity(client *ethclient.Client, entityPrivKey *ecdsa.PrivateKey, agentKeystore *keystore.Key, data []byte) *types.Transaction {

	instance := network.IdentityManagerContract(client)
	fmt.Println("Connected to AlastriaIdentityManager contract")

	opts, _ := TxOpt(entityPrivKey)
	opts.NoSend = true
	tx, _ := instance.CreateAlastriaIdentity(opts, data)
	fmt.Println("Created tx")

	signedTx := crypto.SignTx(network.NetworkID(client), tx, entityPrivKey)
	fmt.Println("Signed Add Key Tx")
	return signedTx

}

func CreateAlastriaIdentity(client *ethclient.Client, entityKeystore *keystore.Key, agentKeystore *keystore.Key) *types.Transaction {
	fmt.Println("Starting CreateAlastriaIdentity")

	addKeyTx := addKeyTx(client, entityKeystore.PrivateKey, agentKeystore)
	createAlastriaIdTx := createAlastriaIdentity(client, entityKeystore.PrivateKey, agentKeystore, addKeyTx.Data())

	fmt.Println("End of CreateAlastriaIdentity")
	return createAlastriaIdTx
}
