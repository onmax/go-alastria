package main

import (
	"fmt"

	"github.com/onmax/go-alastria/alastria"
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/types"
)

func main() {
	// The new agent needs to create their own keystore

	// TODO: Then, the Entity 1 should create an AT, sign it and send it to the wallet
	// The new agent needs to connect to the network
	newAgentArgs := &types.ConnectionArgs{
		NodeUrl: configuration.NodeUrl,
		Keystore: &types.KeystoreConfig{
			Path:     "../../assets/keystores/subject.json",
			Password: "Passw0rd",
		},
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

	// The entity needs to connect to the network
	entityArgs := &types.ConnectionArgs{
		NodeUrl: configuration.NodeUrl,
		Keystore: &types.KeystoreConfig{
			Path:     "../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json",
			Password: "Passw0rd",
		},
	}
	entityClient, _ := alastria.NewClient(entityArgs)

	// The entity needs to connect to the network
	signedPrepareAITx, _ := alastria.PrepareAlastriaId(entityClient, newAgentClient.Client.Ks.Account.Address)

	alastria.SendTx(entityClient, signedPrepareAITx)
	alastria.SendTx(entityClient, signedTxCreateAID)

	agentProxy, _ := alastria.IdentityKeys(entityClient, newAgentClient.Client.Ks.Account.Address)

	fmt.Printf("agentProxy: %v\n", agentProxy)
}
