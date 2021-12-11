package main

import (
	"fmt"
	"time"

	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/tokens"

	"github.com/ethereum/go-ethereum/common"
	"github.com/onmax/go-alastria/alastria"
	exampleutil "github.com/onmax/go-alastria/cmd/examples"

	alaTypes "github.com/onmax/go-alastria/types"
)

var readerClient, _ = alastria.NewClient(exampleutil.GetReaderClientConf())

var subjectKsPath = "../../../../../assets/keystores/subject.json"
var entityKsPath = "../../../../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json"

// See assumptions in README.md
var subjectAddress = common.HexToAddress("d0a0d5a1310a715157c3f81b789d6d9dc447aef5")
var subjectDid, _ = alastria.GetDIDGivenAddress(readerClient, subjectAddress)

var entityAddress = common.HexToAddress("a9728125c573924b2b1ad6a8a8cd9bf6858ced49")
var entityDid, _ = alastria.GetDIDGivenAddress(readerClient, entityAddress)

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
var credentialsToIssueTemp []string = []string{}

var levelOfAssurance = map[string]int{
	"SELF-ISSUED": 0,
	"LOW":         1,
	"MEDIUM":      2,
	"HIGH":        3,
}

// More information about this process in the ./README.md
// To simplify the code, errors are not being checked
func main() {
	// Only issue the credentials that are stored in our database. In this example see var entityDatabase
	credentialsToIssueTemp = []string{"passport", "email"}

	fmt.Printf("\n------ Entity issues credentials to subject ------ \n\n")

	fmt.Printf("\tOmiting step 1-3 as they are only for establishing a secure handshake between subject and issuer\n")
	fmt.Printf("\tStep 4: The entity generates and send signed credentials to subject\n")
	credentials := step4()

	fmt.Printf("As a response of the HTTP request made by the subject, the credential will send the signed credentials\n\n")

	fmt.Printf("\tStep 5: The subject receives the signed credentials and store them somewhere secure. Optionally, the user could write in blockchain that he has received the credentials\n\n")

	step5__subjectWriteInBlockchain(credentials)

}

func step4() []string {
	credentials := step4__generateCredentials()

	fmt.Printf("The entity has generated %v new credentials: %v\n", len(credentials), credentials)

	// Optional step
	// Right now, the function is blocking until the transactions are made. You might want to use TxSendAsync(which is a WIP)
	step4__writeInBlockchain(credentials)

	return credentials
}

func generateCredential(entityDid string, subjectDid string, key string, value string, loa int) tokens.Credential {
	return tokens.Credential{
		Header: &alaTypes.Header{
			Algorithm: "ES256K",
			Type:      "JWT",
		},
		Payload: &tokens.CredentialPayload{
			JSONTokenId: configuration.JtiCredential,
			IssuedAt:    uint64(time.Now().UnixNano()) / uint64(time.Millisecond),
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

func step4__writeInBlockchain(credentials []string) {
	entityConf := exampleutil.GetClientConf(entityKsPath)
	entityClient, _ := alastria.NewClient(entityConf)

	txs, psmHashes, _ := alastria.AddIssuerCredentials(entityClient, credentials, entityDid)

	fmt.Printf("The entity has added %v credentials to the blockchain\n", len(credentials))
	fmt.Printf("The transaction hashes are: %v\n", txs)
	fmt.Printf("The PSMHashes are: %v\n", psmHashes)
}

func step5__subjectWriteInBlockchain(credentials []string) {
	subjectConf := exampleutil.GetClientConf(subjectKsPath)
	subjectClient, _ := alastria.NewClient(subjectConf)

	txs, psmHashes, _ := alastria.AddSubjectCredentials(subjectClient, credentials, subjectDid, "URI")

	fmt.Printf("The entity has added %v credentials to the blockchain\n", len(credentials))
	fmt.Printf("The transaction hashes are: %v\n", txs)
	fmt.Printf("The PSMHashes are: %v\n", psmHashes)
}
