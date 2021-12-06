package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	alaDid "github.com/onmax/go-alastria/ala-did"
	"github.com/onmax/go-alastria/alastria"
	"github.com/onmax/go-alastria/cmd"
	"github.com/onmax/go-alastria/internal/configuration"
)

func main() {
	signedTxCreateAID, newActorAddress := step1__newAgentSignsTx()
	step2__entitySignsPrepareAID_And_SendsTxs(signedTxCreateAID, newActorAddress)
	step3__buildNewAgentDid(newActorAddress)
}

func step1__newAgentSignsTx() (*ethTypes.Transaction, common.Address) {
	// The new agent needs to create their own keystore

	// TODO: Then, the Entity 1 should create an AT, sign it and send it to the wallet
	// The new agent needs to connect to the network
	newAgentConnConf := cmd.GetConnectionConf("../../assets/keystores/subject.json")
	newAgentClient, _ := alastria.NewClient(newAgentConnConf)

	// The subject, from the wallet, should build the tx createAlastriaId and sign it
	signedTxCreateAID, _ := alastria.CreateAlastriaIdentity(newAgentClient)

	return signedTxCreateAID, newAgentClient.Client.Ks.Account.Address
}

func step2__entitySignsPrepareAID_And_SendsTxs(signedTxCreateAID *ethTypes.Transaction, newActorAddress common.Address) {
	// The entity needs to connect to the network
	entityArgs := cmd.GetConnectionConf("../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json")
	entityClient, _ := alastria.NewClient(entityArgs)

	// The entity needs to connect to the network
	signedPrepareAITx, _ := alastria.PrepareAlastriaId(entityClient, newActorAddress)

	alastria.SendTx(entityClient, signedPrepareAITx)
	alastria.SendTx(entityClient, signedTxCreateAID)
}

func step3__buildNewAgentDid(newActorAddress common.Address) {
	entityArgs := cmd.GetConnectionConf("../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json")
	entityClient, _ := alastria.NewClient(entityArgs)

	actorProxy, _ := alastria.IdentityKeys(entityClient, newActorAddress)

	did := alaDid.NewDid(configuration.Network, configuration.NetworkId, actorProxy)

	fmt.Printf("DID of the new actor: %v\n", did)
}