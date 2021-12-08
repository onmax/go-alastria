package main

import (
	"fmt"

	alaDid "github.com/onmax/go-alastria/ala-did"
	"github.com/onmax/go-alastria/alastria"
	exampleutil "github.com/onmax/go-alastria/cmd/examples"
	"github.com/onmax/go-alastria/crypto"
	"github.com/onmax/go-alastria/hex"
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/tokens"
	alaTypes "github.com/onmax/go-alastria/types"
)

// More information about this process in the ../README.md

func main() {
	fmt.Printf("\n------ Creating new actor ------ \n\n")
	fmt.Printf("To simplify the code, errors are not being checked\n\n")

	// The process starts in the entity
	signedAT := step1__entityGeneratesAT()

	fmt.Printf("\tStep 1: Entity generates an AT\n")
	fmt.Printf("AT: %v. You can check its contents in https://jwt.io\n", signedAT)
	fmt.Printf("And somehow sends it to the new agent like a QR code for example.\n\n")

	// Now in the newActor will do step 2
	signedAIC := step2(signedAT)

	fmt.Printf("\tStep 2: New actor generates AIC\n")
	fmt.Printf("The new actor once receives the AT, he generates an AIC: %v. You can check its contents in https://jwt.io\n", signedAIC)

	at, _ := tokens.DecodeAlastriaToken(signedAT)
	fmt.Printf("The new actor needs to send the AIC to the entity using the cbu field from at.payload.cbu. In this case, it will be: %s\n\n", at.Payload.CallbackURL)

	fmt.Printf("\tStep 3: Entity sends the transactions\n")
	fmt.Printf("The entity receives the AIC and destructure the 3 important params mentioned before:\n")
	fmt.Printf("\t- PublicKey -> To verify the AIC and to generate the address of the new actor\n")
	fmt.Printf("\t- Tx of CreateAlastriaID -> The transaction data that the entity needs to run in the SC\n")
	fmt.Printf("\t- The original AT -> In case, the entity needs to keep track of the user from the beginning for his own US\n\n")
	step3__entitySignsPrepareAID_And_SendsTxs(signedAIC)

	// This is not part of the US per se, but it is to show the DID of the new actor
	aic, _ := tokens.DecodeAIC(signedAIC)
	did := step4__buildNewAgentDid(aic.Payload.PublicKey)

	fmt.Printf("\tStep 4: Get the full DID\n")
	fmt.Printf("New actor's Alastria DID: %v\n", did)
}

func step1__entityGeneratesAT() string {
	entityConf := exampleutil.GetDisconnectedClientConf("../../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json")
	entityClient, _ := alastria.NewClient(entityConf)

	at := tokens.AT{
		Header: &alaTypes.Header{
			Algorithm:    "ES256K",
			Type:         "JWT",
			KeyID:        configuration.Kid,
			JSONWebToken: entityClient.Client.Ks.HexPublicKey,
		},
		Payload: &tokens.ATPayload{
			IssuedAt:          configuration.Iat,
			ExpiresAt:         configuration.Exp,
			Issuer:            configuration.DidEntity,
			AlastriaNetworkId: configuration.NetworkId,
			CallbackURL:       configuration.CallbackUrl,
			GatewayURL:        configuration.ProviderURL,
			JSONTokenId:       configuration.Jti,
			Types:             configuration.TypesAT12,
		},
	}
	signedAT, _ := crypto.Sign(at, entityClient.Client.Ks.HexPrivateKey)
	return signedAT
}

func step2(signedAT string) string {
	// The new agent needs to connect to the network
	newAgentConnConf := exampleutil.GetClientConf("../../../assets/keystores/subject.json")
	newAgentClient, _ := alastria.NewClient(newAgentConnConf)

	crypto.Verify(signedAT, newAgentClient.Client.Ks.HexPublicKey) // Remember to check the return value

	signedCreateAIDTx := step2__newAgentSignsTx(newAgentClient)
	aic := step2__newActorGeneratesAIC(newAgentClient, signedAT, signedCreateAIDTx)
	return aic
}

func step2__newAgentSignsTx(newAgentClient *alaTypes.Connection) string {
	// The subject, from the wallet, should build the tx createAlastriaId and sign it
	signedCreateAIDTx, _ := alastria.CreateAlastriaIdentity(newAgentClient)
	signedCreateAIDTxStr, _ := alastria.TxToHex(signedCreateAIDTx)
	return signedCreateAIDTxStr
}

func step2__newActorGeneratesAIC(newAgentClient *alaTypes.Connection, at string, createAlastriaTX string) string {
	aic := tokens.AIC{
		Header: &alaTypes.Header{
			Algorithm:    "ES256K",
			Type:         "JWT",
			JSONWebToken: newAgentClient.Client.Ks.HexPublicKey,
		},
		Payload: &tokens.AICPayload{
			IssuedAt:         configuration.Iat,
			PublicKey:        hex.Add0x(newAgentClient.Client.Ks.HexPublicKey),
			JSONTokenId:      configuration.Jti,
			CreateAlastriaTX: hex.Add0x(createAlastriaTX),
			AlastriaToken:    at,
			Contexts:         configuration.ContextsAIC12,
			Types:            configuration.TypesAIC12,
		},
	}
	signedAIC, _ := crypto.Sign(aic, newAgentClient.Client.Ks.HexPrivateKey)
	return signedAIC
}

func step3__entitySignsPrepareAID_And_SendsTxs(signedAIC string) {
	aic, _ := tokens.DecodeAIC(signedAIC)

	crypto.Verify(signedAIC, aic.Payload.PublicKey) // Remember to check the return value of Verify

	// The entity needs to connect to the network if not already is
	entityArgs := exampleutil.GetClientConf("../../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json")
	entityClient, _ := alastria.NewClient(entityArgs)

	// Convert hex tx to tx struct
	signedCreateAIDTx, _ := alastria.HexToTx(aic.Payload.CreateAlastriaTX)

	// generate and sign signedPrepareAITx
	signedPrepareAITx, _ := alastria.PrepareAlastriaId(entityClient, aic.Payload.PublicKey)

	// Sends both tx in this order
	alastria.SendTx(entityClient, signedPrepareAITx)
	alastria.SendTx(entityClient, signedCreateAIDTx)
}

func step4__buildNewAgentDid(newActorPub string) *alaDid.Did {
	// Any member can connect to the network to do this step, in this case will be the entity
	entityArgs := exampleutil.GetClientConf("../../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json")
	entityClient, _ := alastria.NewClient(entityArgs)

	newActorAddress, _ := alastria.PublicKeyToAddress(newActorPub)
	actorProxy, _ := alastria.IdentityKeys(entityClient, newActorAddress)

	return alaDid.NewDid(configuration.Network, configuration.NetworkId, actorProxy)
}
