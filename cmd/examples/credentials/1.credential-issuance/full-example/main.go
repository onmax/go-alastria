package main

import (
	"fmt"

	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/tokens"

	"github.com/ethereum/go-ethereum/common"
	"github.com/onmax/go-alastria/alastria"
	exampleutil "github.com/onmax/go-alastria/cmd/examples"

	alaTypes "github.com/onmax/go-alastria/types"
)

var entityKsPath = "../../../../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json"
var subjectKsPath = "../../../../../assets/keystores/subject.json"

// See assumptions in README.md
var subjectAddress = common.HexToAddress("d0a0d5a1310a715157c3f81b789d6d9dc447aef5")
var subjectDid = exampleutil.GetDIDGivenAddress(subjectAddress)

var entityAddress = common.HexToAddress("a9728125c573924b2b1ad6a8a8cd9bf6858ced49")
var entityDid = exampleutil.GetDIDGivenAddress(entityAddress)

// entityDatabase simulates the data that the entity should store somewhere in the backend
// database
var entityDatabase = map[string]string{
	"subjectDid": subjectDid.String(),
	"legacy-id":  "123",

	// Data of the user
	"passport":     "12345678A",
	"email":        "example@mail.org",
	"home_address": "",
}

// Temporally data
var jti = ""
var credentialsToIssueTemp []string = []string{}

var levelOfAssurance = map[string]int{
	"SELF-ISSUED": 0,
	"LOW":         1,
	"MEDIUM":      2,
	"HIGH":        3,
}

// More information about this process in the ./README.md

func main() {
	fmt.Printf("\n------ Entity issues credentials to subject ------ \n\n")
	fmt.Printf("To simplify the code, errors are not being checked\n\n")

	subjectDid, credentialsKeys := step1__subjectSelectsCredentials()
	fmt.Printf("\tStep 1: The subject selects credentials\n")
	fmt.Printf("In this case, the subject DID is %s and has selected the credentials %v\n", subjectDid.String(), credentialsKeys)
	fmt.Printf("The subject will send a request to the backend to start the process\n\n")

	signedAT := step2(subjectDid.String(), credentialsKeys)
	fmt.Printf("\tStep 2: The entity receives the request\n")
	fmt.Printf("The credentials that will be issued will be: %s\n", credentialsToIssueTemp)
	fmt.Printf("The entity generates an AT: %v. You can check its contents in https://jwt.io\n", signedAT)
	fmt.Printf("The entity will send the AT to the subject's wallet using QR, Push notification, Deeplink, tinyURL...\n\n")

	signedAS := step3__subjectGeneratesAS(signedAT)
	fmt.Printf("\tStep 3: The subject generates an AS\n")
	fmt.Printf("The subject generates an AS: %v. You can check its contents in https://jwt.io\n", signedAS)
	at, _ := tokens.DecodeAlastriaToken(signedAT)
	fmt.Printf("The subject will send the AS to the entity to the URL in the at.Payload.cbu. In this case: %s\n\n", at.Payload.CallbackURL)

	fmt.Printf("\tStep 4: The entity generates and send signed credentials to subject\n")
	credentials := step4(signedAS)

	fmt.Printf("As a response of the HTTP request made by the subject, the credential will send the signed credentials\n\n")

	fmt.Printf("\tStep 5: The subject receives the signed credentials and store them somewhere secure. Optionally, the user could write in blockchain that he has received the credentials\n\n")

	step5__subjectWriteInBlockchain(credentials)

	// Clear temp data
	jti = ""
	credentialsToIssueTemp = []string{}
}

func step1__subjectSelectsCredentials() (*alaTypes.Did, []string) {
	// This step can be done in any platform: website of the entity, mobile app, etc...
	// In this example, we hardcode the credentials that the subject has selected
	return subjectDid, []string{"passport", "email", "home_address"}
}

func step2(subjectDid string, credentials []string) string {
	step2__filterCredentials(subjectDid, credentials)
	return step2__entityGeneratesAT()
}

// Checks that has the requested credentials. In this example, the entity does not have
// "home_address" so we will ignore it. The issuer could return an error if the
// credential is not available and it is crucial to the process
func step2__filterCredentials(subjectDid string, credentials []string) {
	var credentialsToIssue []string
	for _, credential := range credentials {
		if entityDatabase["subjectDid"] == subjectDid && entityDatabase[credential] != "" {
			credentialsToIssue = append(credentialsToIssue, credential)
		}
	}
	credentialsToIssueTemp = credentialsToIssue
}

func step2__entityGeneratesAT() string {
	entityConf := exampleutil.GetDisconnectedClientConf(entityKsPath)
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
			CallbackURL:       configuration.CallbackUrlCredentialIssuance,
			GatewayURL:        configuration.ProviderURL,
			JSONTokenId:       configuration.JtiAT,
			Types:             configuration.TypesAT221,
			// ignoring context
		},
	}

	// Save the relation of the subject and the AT generated just for him
	jti = at.Payload.JSONTokenId

	signedAT, _ := alastria.SignJWT(entityClient, at)
	return signedAT
}

func step3__subjectGeneratesAS(signedAT string) string {
	// This step is made in the wallet of the subject
	subjectConf := exampleutil.GetDisconnectedClientConf("../../../../../assets/keystores/subject.json")
	subjectClient, _ := alastria.NewClient(subjectConf)

	as := tokens.AS{
		Header: &alaTypes.Header{
			Algorithm:    "ES256K",
			Type:         "JWT",
			JSONWebToken: subjectClient.Client.Ks.HexPublicKey,
		},
		Payload: &tokens.ASPayload{
			IssuedAt:      configuration.Iat,
			ExpiresAt:     configuration.Exp,
			JSONTokenId:   configuration.JtiAS,
			AlastriaToken: signedAT,
			Issuer:        subjectDid.String(),
			Contexts:      configuration.ContextsAIC12,
			Types:         configuration.TypesASCredentialIssuance,
		},
	}
	signedAS, _ := alastria.SignJWT(subjectClient, as)
	return signedAS
}

func step4(signedAS string) []string {
	// AlastriaSession is being used for security reasons only
	step4__verifyAS(signedAS)
	step4__verifyDid(signedAS)

	credentials := step4__generateCredentials()

	fmt.Printf("The entity has generated %v new credentials: %v\n", len(credentials), credentials)

	// Optional step
	// Right now, the function is blocking until the transactions are made. You might want to use TxSendAsync
	step4__issuerWriteInBlockchain(credentials)

	return credentials
}

func step4__verifyAS(signedAS string) {
	entityConf := exampleutil.GetReaderClientConf()
	entityClient, _ := alastria.NewClient(entityConf)
	subjectProxy := common.HexToAddress(subjectDid.ProxyAddress)
	pubKey, _ := alastria.GetCurrentPublicKey(entityClient, subjectProxy)
	res, _ := alastria.VerifyJWT(signedAS, pubKey)

	fmt.Printf("The entity needs to verify the AS. The entity has the DID of the subject (from previous steps): %s. The entity takes the proxy, which is the address of the user and with that address, the entity calls GetCurrentPublicKey\n", subjectDid.String())
	fmt.Printf("IMPORTANT! The public key of the subject is %s. The verification has been successful -> %v\n", pubKey, res)
}

func step4__verifyDid(signedAS string) {
	as, _ := tokens.DecodeAlastriaSession(signedAS)
	iss := as.Payload.Issuer
	fmt.Printf("IMPORTANT! The issuer of the AS should be the subject's DID: %s. AS.payload.iss is the same value that the entity has in its database? %v\n", iss, entityDatabase["subjectDid"] == iss)

	at, _ := tokens.DecodeAlastriaToken(as.Payload.AlastriaToken)
	jtiAt := at.Payload.JSONTokenId
	fmt.Printf("The jti of the AS should be the jti of the AT: %s. AS.payload.jti is the same value that the entity has in its temporal database? %v. So this way we know which user sent the request.\n", jtiAt, jti == jtiAt)
}

func generateCredential(entityDid string, subjectDid string, key string, value string, loa int) tokens.Credential {
	return tokens.Credential{
		Header: &alaTypes.Header{
			Algorithm: "ES256K",
			Type:      "JWT",
		},
		Payload: &tokens.CredentialPayload{
			JSONTokenId: configuration.JtiCredential,
			IssuedAt:    configuration.Iat,
			ExpiresAt:   configuration.Exp,
			Issuer:      entityDid,
			Subject:     subjectDid,
			VerifiableCredential: &tokens.CredentialPayloadVC{
				Contexts: configuration.ContextsCredentials,
				Types:    configuration.TypesCredentials,
				CredentialSubject: &map[string]interface{}{
					"levelOfAssurance": loa,
					"data":             map[string]interface{}{key: value},
				},
			},
		},
	}
}

func step4__generateCredentials() []string {
	var credentials []string

	entityConf := exampleutil.GetDisconnectedClientConf(entityKsPath)
	entityClient, _ := alastria.NewClient(entityConf)

	for _, credentialKey := range credentialsToIssueTemp {
		// The level of assurance can be any value between 0 and 3, depends on the issuer
		credential := generateCredential(entityDid.String(), subjectDid.String(), credentialKey, entityDatabase[credentialKey], levelOfAssurance["LOW"])

		signedCredential, _ := alastria.SignJWT(entityClient, credential)
		credentials = append(credentials, signedCredential)
	}
	return credentials
}

func step4__issuerWriteInBlockchain(credentials []string) {
	entityConf := exampleutil.GetClientConf(entityKsPath)
	entityClient, _ := alastria.NewClient(entityConf)

	var (
		txs       []string = make([]string, len(credentials))
		psmHashes []string = make([]string, len(credentials))
	)

	for _, credential := range credentials {
		// Right now, the function is blocking until the transactions are made. You might want to use
		// another thread for this
		tx, psmHash, _ := alastria.AddIssuerCredential(entityClient, credential, entityDid.String())
		txs = append(txs, tx)
		psmHashes = append(psmHashes, psmHash)
	}

	fmt.Printf("The entity has added %v credentials to the blockchain\n", len(credentials))
	fmt.Printf("The transaction hashes are: %v\n", txs)
	fmt.Printf("The PSMHashes are: %v\n", psmHashes)
}

func step5__subjectWriteInBlockchain(credentials []string) {
	subjectConf := exampleutil.GetClientConf(subjectKsPath)
	subjectClient, _ := alastria.NewClient(subjectConf)

	var (
		txs       []string = make([]string, len(credentials))
		psmHashes []string = make([]string, len(credentials))
	)

	for _, credential := range credentials {
		// Right now, the function is blocking until the transactions are made. You might want to use
		// another thread for this
		tx, psmHash, _ := alastria.AddSubjectCredential(subjectClient, credential, entityDid.String(), "URI")
		txs = append(txs, tx)
		psmHashes = append(psmHashes, psmHash)
	}

	fmt.Printf("The entity has added %v credentials to the blockchain\n", len(credentials))
	fmt.Printf("The transaction hashes are: %v\n", txs)
	fmt.Printf("The PSMHashes are: %v\n", psmHashes)
}
