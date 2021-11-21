package main

import (
	"context"
	"fmt"
	"log"

	"github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/network"
	"github.com/onmax/go-alastria/tx"
)

func main() {
	client := network.ConnectToNetwork("http://34.91.211.67:22000")

	fmt.Println("Success! you are connected to the Alastria Network http://34.91.211.67:22000")

	entityKs, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
	agentKs, _, agentPubKey, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")

	signedAITx := tx.CreateAlastriaIdentity(client, entityKs, agentKs)
	fmt.Printf("createAITx: %v\n\n", signedAITx)

	signedPrepareAITx := tx.PrepareAlastriaId(client, entityKs, agentPubKey)
	fmt.Printf("createPrepareAITx: %v\n\n", signedPrepareAITx)

	err := client.SendTransaction(context.Background(), signedPrepareAITx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("signedPrepareAITx sent: %s\n", signedPrepareAITx.Hash().Hex())

	// time.Sleep(time.Second * 10)

	err = client.SendTransaction(context.Background(), signedAITx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("signedAITx sent: %s\n", signedAITx.Hash().Hex())
	// time.Sleep(5 * time.Second)

	did := tx.IdentityKeys(client, agentKs.Address)
	fmt.Printf("DID: %v\n", did)
}
