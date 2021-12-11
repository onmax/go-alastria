package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onmax/go-alastria/alastria"
	exampleutil "github.com/onmax/go-alastria/cmd/examples"
	"github.com/onmax/go-alastria/internal/configuration"
	alaTypes "github.com/onmax/go-alastria/types"
)

var entityKsPath = "../../../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json"
var subjectKsPath = "../../../../assets/keystores/subject.json"

// More information about this process in the ../README.md

func main() {
	fmt.Printf("\n------ Creating new actor | Only blockchain part ------ \n\n")
	fmt.Printf("To simplify the code, errors are not being checked\n\n")

	fmt.Printf("\tOmitting step 1 as it has nothing to do with blockchain\n")

	fmt.Printf("\tStep 2: New agent generates CreateAlastriaID transaction\n")
	signedTxCreateAID, newActorPublicKey := step2__newAgentSignsTx()
	fmt.Printf("The signed tx CreateAID generated by new agent: %v\n", signedTxCreateAID)
	fmt.Printf("The public key of the new agent: %v\n", newActorPublicKey)
	fmt.Printf("Both of them should go in the AIC\n")

	fmt.Printf("\tStep 3: Send both transaction\n")
	fmt.Printf("The entity needs first to parse the signedTxCreateAID transaction, and then, the entity send both transactions in a specific order\n")
	step3__entitySignsPrepareAID_And_SendsTxs(signedTxCreateAID, newActorPublicKey)

	// This is not part of the US per se, but it is to show the DID of the new actor
	did := step4__buildNewAgentDid(newActorPublicKey)

	fmt.Printf("\tStep 4: Get the full DID\n")
	fmt.Printf("DID of the new actor: %v\n", did)

}

func step2__newAgentSignsTx() (*types.Transaction, string) {
	// The new agent needs to connect to the network
	newAgentConnConf := exampleutil.GetClientConf(subjectKsPath)
	newAgentClient, _ := alastria.NewClient(newAgentConnConf)

	// The subject, from the wallet, should build the tx createAlastriaId and sign it
	signedTxCreateAID, _ := alastria.CreateAlastriaIdentity(newAgentClient)

	return signedTxCreateAID, newAgentClient.Client.Ks.HexPublicKey
}

func step3__entitySignsPrepareAID_And_SendsTxs(signedTxCreateAID *types.Transaction, newActorPublicKey string) {
	// The entity needs to connect to the network
	entityArgs := exampleutil.GetClientConf(entityKsPath)
	entityClient, _ := alastria.NewClient(entityArgs)

	signedPrepareAITx, _ := alastria.PrepareAlastriaId(entityClient, newActorPublicKey)

	alastria.SendTx(entityClient, signedPrepareAITx)
	alastria.SendTx(entityClient, signedTxCreateAID)
}

func step4__buildNewAgentDid(newActorPub string) *alaTypes.Did {
	entityArgs := exampleutil.GetReaderClientConf()
	entityClient, _ := alastria.NewClient(entityArgs)

	newActorAddress, _ := alastria.PublicKeyToAddress(newActorPub)
	actorProxy, _ := alastria.IdentityKeys(entityClient, newActorAddress)

	return alastria.NewDid(configuration.Network, configuration.NetworkId, actorProxy)
}
