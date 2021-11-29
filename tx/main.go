// package tx

// import (
// 	"context"
// 	"encoding/hex"
// 	"fmt"
// 	"log"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"

// 	"github.com/onmax/go-alastria/pkg/crypto"
// 	"github.com/onmax/go-alastria/pkg/keystore"
// 	"github.com/onmax/go-alastria/pkg/tokens"
// )

// // var alastriaIdentityManager = "0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"
// var alastriaIdentityManager = "0xeafCAf8c3f9016B14ec65105677658f3D6Eb9079" // truffle in alastria
// // var publicKeyRegistry = "0x017B0F93d1F31666Ae0bfb35f2e35bC1b0AC43C1"
// var publicKeyRegistry = "0x1cd13Dc5161Bf2d5A83AAB4e7b66A6542D10Ab5B" // truffle in alastria

// // k, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
// // k2, _, pubKey, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")

// // _, _, pubk2, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")
// // k, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")

// // func main() {
// // 	client := network.ConnectToNetwork("http://34.91.211.67:22000")

// // 	fmt.Println("Success! you are connected to the Alastria Network http://34.91.211.67:22000")

// // 	entityKs, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
// // 	agentKs, _, agentPubKey, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")

// // 	createAITx := CreateAlastriaIdentity(client, entityKs, agentKs)
// // 	fmt.Printf("createAITx: %v\n\n", createAITx)
// // 	// aic := signAIC(*tx)
// // 	createPrepareAITx := PrepareAlastriaId(client, entityKs, agentPubKey)
// // 	fmt.Printf("createPrepareAITx: %v\n\n", createPrepareAITx)

// // if tx2 == nil {
// // 	fmt.Println("No tx2")
// // }

// // fmt.Println("Sending tx2")
// // err = client.SendTransaction(context.Background(), tx2)
// // if err != nil {
// // 	log.Fatal(err)
// // }

// // pending := true
// // fmt.Printf("OUTLOOP: Waiting for tx2 to be mined %v %v \n", tx2.Hash(), pending)
// // for pending {
// // 	_, pending, _ = client.TransactionByHash(context.Background(), tx2.Hash())
// // 	fmt.Printf("[IN] Waiting for tx2 to be mined %v %v \n", tx2.Hash(), pending)
// // 	time.Sleep(time.Second * 1)
// // }

// // fmt.Println("Sending tx")
// // tx := CreateAlastriaIdentity(client)
// // if tx == nil {
// // 	fmt.Println("No tx")
// // }
// // err = client.SendTransaction(context.Background(), tx)
// // if err != nil {
// // 	log.Fatal(err)
// // }
// // pending = true
// // fmt.Printf("OUTLOOP: Waiting for tx2 to be mined  %v %v\n", tx.Hash().String(), pending)
// // for pending {
// // 	_, pending, _ = client.TransactionByHash(context.Background(), tx.Hash())
// // 	fmt.Printf("[IN] Waiting for tx to be mined %v %v \n", tx.Hash(), pending)
// // 	time.Sleep(time.Second * 1)
// // }

// // fmt.Printf("---------\ntx2 %v\n---------\n", tx2.Hash().Hex())
// // fmt.Printf("%v\n\n", tx2.Data())
// // fmt.Printf("%x\n\n", tx2.Data())
// // fmt.Printf("---------\ntx %v\n---------\n", tx.Hash().Hex())
// // fmt.Printf("%v\n\n", tx.Data())
// // fmt.Printf("%x\n\n", tx.Data())

// // r, _ := client.TransactionReceipt(context.Background(), tx.Hash())
// // fmt.Printf("---------  TransactionReceipt %v---------\n", r)
// // // fmt.Printf("%v\n\n", r.TxHash.Hex())
// // // fmt.Printf("%v\n\n", r.Bloom)

// // // subscribePAID(client, tx2.Hash())
// // queryPAID := ethereum.FilterQuery{
// // 	Addresses: []common.Address{common.HexToAddress(alastriaIdentityManager)},
// // }
// // logs, err := client.FilterLogs(context.Background(), queryPAID)
// // if err != nil {
// // 	log.Fatal(err)
// // }
// // for _, vLog := range logs {
// // 	fmt.Println("NEW CREATE EVENT")
// // 	fmt.Println(vLog)
// // }
// // // end subscribePAID
// // // subscribeCreate(client, tx.Hash())
// // queryCra := ethereum.FilterQuery{
// // 	Addresses: []common.Address{common.HexToAddress(publicKeyRegistry)},
// // }

// // logs2, err := client.FilterLogs(context.Background(), queryCra)
// // if err != nil {
// // 	log.Fatal(err)
// // }
// // for _, vLog := range logs2 {
// // 	fmt.Println("NEW CREATE EVENT")
// // 	fmt.Println(vLog)
// // }
// // end subscribePAID

// // fmt.Printf("---------\naic %v\n---------\n", aic)
// // }

// func subscribePAID(client *ethclient.Client, hash common.Hash) {
// 	queryPAID := ethereum.FilterQuery{
// 		Addresses: []common.Address{common.HexToAddress(alastriaIdentityManager)},
// 	}
// 	logs, err := client.FilterLogs(context.Background(), queryPAID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, vLog := range logs {
// 		fmt.Println("NEW CREATE EVENT")
// 		fmt.Println(vLog)
// 	}
// }

// func subscribeCreate(client *ethclient.Client, hash common.Hash) {
// 	address := common.HexToAddress(publicKeyRegistry)
// 	queryPAID := ethereum.FilterQuery{
// 		Addresses: []common.Address{address},
// 	}

// 	logs, err := client.FilterLogs(context.Background(), queryPAID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, vLog := range logs {
// 		fmt.Println("NEW CREATE EVENT")
// 		fmt.Println(vLog)
// 	}
// }

// func signAIC(tx types.Transaction) string {
// 	_, privKeyEntity1, pubKeyEntity1, err := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
// 	didEntity1 := "did:ala:quor:redT:d2f868f056ef3a48bbc8d446dfed411e9bf93ab0"
// 	providerURL := "https://regular.telsius.blockchainbyeveris.io:2000"
// 	callbackUrl := "https://serviceprovider.alastria.blockchainbyeveris.io/api/login/"
// 	networkId := "Alastria network"
// 	exp := uint64(2036798552159)
// 	kidCredential := "did:ala:quor:redt:12eeaCCA9eEbB78eB97d7cac6b#keys-1"
// 	iat := uint64(1636798552159)
// 	jti := "ze298y42sba"
// 	at := tokens.AT{
// 		Header: &tokens.Header{
// 			Algorithm:    "ES256K",
// 			Type:         "JWT",
// 			KeyID:        kidCredential,
// 			JSONWebToken: pubKeyEntity1,
// 		},
// 		Payload: &tokens.ATPayload{
// 			IssuedAt:          iat,
// 			ExpiresAt:         exp,
// 			Issuer:            didEntity1,
// 			AlastriaNetworkId: networkId,
// 			CallbackURL:       callbackUrl,
// 			GatewayURL:        providerURL,
// 			JSONTokenId:       jti,
// 			Types:             []string{"AlastriaToken"},
// 		},
// 	}

// 	signedAT, err := crypto.Sign(at, privKeyEntity1)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Printf("\nSigned AT %s\n\n", signedAT)
// 	}

// 	_, pk2, pubk2, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")

// 	b, _ := tx.MarshalBinary()
// 	createAlastriaIdTx := "0x" + hex.EncodeToString(b)

// 	aic := tokens.AIC{
// 		Header: &tokens.Header{
// 			Algorithm:    "ES256K",
// 			Type:         "JWT",
// 			KeyID:        "key-id",
// 			JSONWebToken: "json-web-token",
// 		},
// 		Payload: &tokens.AICPayload{
// 			IssuedAt:         1,
// 			PublicKey:        pubk2,
// 			JSONTokenId:      "json-token-id",
// 			CreateAlastriaTX: createAlastriaIdTx,
// 			AlastriaToken:    signedAT,
// 			Contexts:         []string{"https://alastria.github.io/identity/artifacts/v1"},
// 			Types:            []string{"AlastriaIdentityCreation", "US12"},
// 		},
// 	}
// 	aicSigned, _ := crypto.Sign(aic, pk2)
// 	return aicSigned
// }