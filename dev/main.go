package main

import (
	"context"
	"fmt"
	"time"

	"github.com/onmax/go-alastria/alastria"
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/types"
)

func main() {
	// The new agent needs to create their own keystore
	// keystoreNewActor, err := keystore.ImportKs("../assets/keystores/new-subject.json/UTC--2021-11-30T15-08-53.335386742Z--d0a0d5a1310a715157c3f81b789d6d9dc447aef5", "Passw0rd")
	keystoreNewActor, err := keystore.ImportKs("../assets/keystores/subject1.json", "Passw0rd")
	if err != nil {
		fmt.Println(err)
	}

	// TODO: Then, the Entity 1 should create an AT, sign it and send it to the wallet

	// The new agent needs to connect to the network
	newAgentArgs := &types.ConnectionArgs{
		NodeUrl:  configuration.NodeUrl,
		Keystore: keystoreNewActor,
	}
	newAgentClient, err := alastria.NewClient(newAgentArgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	// The subject, from the wallet, should build the tx createAlastriaId and sign it
	signedTxCreateAID, err := alastria.CreateAlastriaIdentity(newAgentClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("signedTxCreateAID: %v\n", signedTxCreateAID.Hash().Hex())
	fmt.Printf("signedTxCreateAID: %v\n", signedTxCreateAID)

	// The entity needs to connect to the network
	keystoreEntity, err := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
	if err != nil {
		fmt.Println(err)
	}
	entityArgs := &types.ConnectionArgs{
		NodeUrl:  configuration.NodeUrl,
		Keystore: keystoreEntity,
	}
	entityClient, _ := alastria.NewClient(entityArgs)

	// The entity needs to connect to the network
	signedPrepareAITx, _ := alastria.PrepareAlastriaId(entityClient, keystoreNewActor.HexPublicKey)
	fmt.Printf("signedPrepareAITx: %v\n", signedPrepareAITx.Hash().Hex())
	fmt.Printf("newAgentClient.Client.Ks.Account.Address: %v\n", newAgentClient.Client.Ks.Account.Address)

	// At the end, the entity 1 should send both tx (signedPrepareAITx and signedTxCreateAID, in that order) to the blockchain as follows:
	entityClient.Client.Eth.SendTransaction(context.Background(), signedPrepareAITx)
	time.Sleep(10 * time.Second)
	entityClient.Client.Eth.SendTransaction(context.Background(), signedTxCreateAID)
	time.Sleep(10 * time.Second)

	// Now, retrieve the proxy from new subject to display it which is part of the did
	agentProxy, _ := alastria.IdentityKeys(entityClient, newAgentClient.Client.Ks.Account.Address)

	fmt.Printf("agentProxy: %v\n", agentProxy)

	// signedAITx := tx.CreateAlastriaIdentity(client, keystoreEntity., agentKs)
	// signedAITx = tx.CreateAlastriaIdentity(client, entityKs, agentKs)
	// fmt.Printf("createAITx: %v\n\n", signedAITx)

	// signedPrepareAITx := tx.PrepareAlastriaId(client, entityKs, agentPubKey)
	// fmt.Printf("createPrepareAITx: %v\n\n", signedPrepareAITx)

	// err := client.SendTransaction(context.Background(), signedPrepareAITx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("signedPrepareAITx sent: %s\n", signedPrepareAITx.Hash().Hex())

	// // time.Sleep(time.Second * 10)

	// err = client.SendTransaction(context.Background(), signedAITx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("signedAITx sent: %s\n", signedAITx.Hash().Hex())
	// // time.Sleep(5 * time.Second)

	// did := tx.IdentityKeys(client, agentKs.Address)
	// fmt.Printf("DID: %v\n", did)
}
