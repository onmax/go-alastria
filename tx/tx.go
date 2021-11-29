package tx

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/pkg/app"
	"github.com/onmax/go-alastria/pkg/pkg/crypto"
	"github.com/onmax/go-alastria/pkg/pkg/network"
)

func NewAlastriaClient(node string) (*app.AlastriaConnection, error) {
	// client, err := network.ConnectToNetwork(node)
	// if err != nil {
	// 	return nil, err
	// }
	client := &app.AlastriaClient{
		eth: nil,
		ks:  nil,
	}
	p := app.Person{name: ""}
	fmt.Printf("p: %v\n", p)
	return &app.AlastriaConnection{
		client: &app.AlastriaClient{
			eth: nil,
			ks:  nil,
		},
		contracts: &app.AlastriaContracts{
			identity_manager:    nil,
			public_key_registry: nil,
		},
		network: &app.AlastriaNetwork{
			id:       big.NewInt(0),
			node_url: node,
		},
	}, nil
}

func newClient(node string) (*app.AlastriaClient, error) {

	client, err := network.ConnectToNetwork(node)
	if err != nil {
		return nil, err
	}
	return &app.AlastriaClient{
		eth: client,
		ks:  nil,
	}, nil
}

func (conn *app.AlastriaConnection) PrepareAlastriaId(newAgentPubKey string, a *app.AlastriaClient) (*types.Transaction, error) {
	instance, err := network.IdentityManagerContract(conn.client.eth)
	if err != nil {
		return nil, err
	}
	opts, _ := TxOpt(conn.client.ks)
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
