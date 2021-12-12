package main

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/onmax/go-alastria/alastria"
	exampleutil "github.com/onmax/go-alastria/cmd/examples"
	"github.com/onmax/go-alastria/internal/configuration"
	"github.com/onmax/go-alastria/tokens"
	alaTypes "github.com/onmax/go-alastria/types"
)

var readerClient, _ = alastria.NewClient(exampleutil.GetReaderClientConf())

var subjectAddress = common.HexToAddress("d0a0d5a1310a715157c3f81b789d6d9dc447aef5")
var subjectDid, _ = alastria.GetDIDGivenAddress(readerClient, subjectAddress)

var entityAddress = common.HexToAddress("a9728125c573924b2b1ad6a8a8cd9bf6858ced49")
var entityDid, _ = alastria.GetDIDGivenAddress(readerClient, entityAddress)

var subjectKsPath = "../../../../assets/keystores/subject.json"
var entityKsPath = "../../../../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json"

// More information about this process in the ./README.md
// To simplify the code, errors are not being checked

func main() {
	// This is only to generate real new data in the blockchain
	signedCredentials := generateCredentials()

	// Example begins here
	fmt.Printf("\n------ Any actor gets the PSMhashes of credentials of a subject ------ \n\n")

	actorConf := exampleutil.GetReaderClientConf()
	actorClient, _ := alastria.NewClient(actorConf)

	credentials, e := alastria.GetCredentials(actorClient, signedCredentials, subjectDid)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("There are %d PSMHashes in the blockchain\n", len(credentials))
	fmt.Printf("Subject DID: %s\n", subjectDid)
	fmt.Printf("Subject subject: %s\n", subjectAddress)
	fmt.Printf("The status shown here is the highest between subject status and issuer status\n")

	for _, credential := range credentials {
		fmt.Printf("credential: %v\n", credential.Credential[:20])
		shortCredential := credential.Credential[:20] + "..." + credential.Credential[len(credential.Credential)-20:]
		fmt.Printf("\tCredential %v || Status %v || IssuerPSMHash %v || SubjectPSMHash %v\n", shortCredential, credential.Status, credential.IssuerPSMHash, credential.SubjectPSMHash)
	}
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

// this function is used to generate new data in the blockchain to show the example.
// To see more details on how the issuer and subject writes a new credential in the blockchain,
// see example 1.credential-issuance
func generateCredentials() []string {
	entityConf := exampleutil.GetClientConf(entityKsPath)
	entityClient, _ := alastria.NewClient(entityConf)

	credential1 := generateCredential(entityDid.String(), subjectDid.String(), "passport", "01234567A", 1)
	signedCredential1, _ := alastria.SignJWT(entityClient, credential1)

	credential2 := generateCredential(entityDid.String(), subjectDid.String(), "email", "example@mail.org", 1)
	signedCredential2, _ := alastria.SignJWT(entityClient, credential2)

	var credentials []string = []string{signedCredential1, signedCredential2}

	alastria.AddIssuerCredentials(entityClient, credentials, entityDid)

	subjectConf := exampleutil.GetClientConf(subjectKsPath)
	subjectClient, _ := alastria.NewClient(subjectConf)

	alastria.AddSubjectCredentials(subjectClient, credentials, subjectDid, "URI")

	return credentials
}
