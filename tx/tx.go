package tx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"
	"github.com/onmax/go-alastria/crypto"
	"github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/network"
)

type AlastriaContracts struct {
	identity_manager    *identity.AlastriaContracts
	public_key_registry *pkr.AlastriaContracts
}

type AlastriaNetwork struct {
	id       *big.Int
	node_url string
}

type AlastriaClient struct {
	eth *ethclient.Client
	ks  *keystore.Keystore
}

type AlastriaConnection struct {
	contracts *AlastriaContracts
	network   *AlastriaNetwork
	client    *AlastriaClient
}

func (conn *AlastriaConnection) PrepareAlastriaId(client *ethclient.Client, newAgentPubKey string) (*types.Transaction, error) {
	instance, err := network.IdentityManagerContract(client)
	if err != nil {
		return nil, err
	}
	opts, _ := TxOpt(conn.client.ks.)
	opts.NoSend = true
	tx, _ := instance.CreateAlastriaIdentity(opts, []byte(newAgentPubKey))
	signedTx := crypto.SignTx(conn.network.id, tx, conn.client.ks.private_key)
	return signedTx, nil
}

// func addKeyTx(client *ethclient.Client, entityPrivKey *ecdsa.PrivateKey, agentKeystore *keystore.Key) *types.Transaction {
// 	instance := network.PublicKeyRegistryContract(client)
// 	fmt.Println("Connected to PublicKeyRegistry contract")

// 	agentPrivKey := agentKeystore.PrivateKey

// 	opts, _ := TxOpt(entityPrivKey)
// 	opts.NoSend = true
// 	tx, _ := instance.AddKey(opts, alastriaKs.PublicKeyToString(&agentPrivKey.PublicKey))
// 	fmt.Println("Created tx")

// 	signedTxAddKey := crypto.SignTx(network.NetworkID(client), tx, agentPrivKey)
// 	fmt.Println("Signed Add Key Tx")
// 	return signedTxAddKey
// }

// func createAlastriaIdentity(client *ethclient.Client, entityPrivKey *ecdsa.PrivateKey, agentKeystore *keystore.Key, data []byte) (*types.Transaction, error) {

// 	instance, err := network.IdentityManagerContract(client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	opts, _ := TxOpt(entityPrivKey)
// 	opts.NoSend = true
// 	nonce, _ := client.PendingNonceAt(context.Background(), opts.From)
// 	opts.Nonce = new(big.Int).SetUint64(nonce + 1)
// 	tx, _ := instance.CreateAlastriaIdentity(opts, data)
// 	fmt.Println("Created tx")

// 	signedTx := crypto.SignTx(network.NetworkID(client), tx, entityPrivKey)
// 	fmt.Println("Signed Add Key Tx")
// 	return signedTx

// }

// func CreateAlastriaIdentity(client *ethclient.Client, entityKeystore *keystore.Key, agentKeystore *keystore.Key) *types.Transaction {
// 	fmt.Println("----- Starting CreateAlastriaIdentity -----")

// 	addKeyTx := addKeyTx(client, entityKeystore.PrivateKey, agentKeystore)
// 	createAlastriaIdTx := createAlastriaIdentity(client, entityKeystore.PrivateKey, agentKeystore, addKeyTx.Data())

// 	fmt.Println("----- End of CreateAlastriaIdentity -----")
// 	return createAlastriaIdTx
// }

// func IdentityKeys(client *ethclient.Client, agentAddress common.Address) common.Address {
// 	instance := network.IdentityManagerContract(client)
// 	fmt.Println("Connected to AlastriaIdentityManager contract")

// 	address, err := instance.IdentityKeys(&bind.CallOpts{}, agentAddress)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// opts, _ := TxOpt(entityPrivKey)
// 	// opts.NoSend = true
// 	// nonce, _ := client.PendingNonceAt(context.Background(), opts.From)
// 	// opts.Nonce = new(big.Int).SetUint64(nonce + 1)
// 	// tx, _ := instance.CreateAlastriaIdentity(opts, data)
// 	fmt.Println("Got address tx")
// 	return address
// }
