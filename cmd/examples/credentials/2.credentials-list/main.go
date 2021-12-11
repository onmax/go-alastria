package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/onmax/go-alastria/alastria"
	exampleutil "github.com/onmax/go-alastria/cmd/examples"
	alaTypes "github.com/onmax/go-alastria/types"
)

var subjectAddress = common.HexToAddress("d0a0d5a1310a715157c3f81b789d6d9dc447aef5")
var subjectDid = exampleutil.GetDIDGivenAddress(subjectAddress)

// More information about this process in the ./README.md
// To simplify the code, errors are not being checked

func main() {
	fmt.Printf("\n------ Any actor gets the PSMhashes of credentials of a subject ------ \n\n")

	actorConf := exampleutil.GetReaderClientConf()
	actorClient, _ := alastria.NewClient(actorConf)

	psmHashes := step1__getCredentialList(actorClient)
	fmt.Printf("The subject %s (with did: %v) has %v credentials registered in the blockchain:\n", subjectAddress, subjectDid, len(psmHashes))

	credentialsStatues := step2__getStatusForEveryCredential(actorClient, subjectDid, psmHashes)

	for _, status := range credentialsStatues {
		fmt.Printf("\tCredential %v || Status %v\n", status.PSMHash, status.Status)
	}
}

func step1__getCredentialList(actorClient *alaTypes.Connection) []common.Address {
	psmHashes, _ := alastria.GetSubjectCredentialList(actorClient, subjectAddress)
	return psmHashes
}

func step2__getStatusForEveryCredential(actorClient *alaTypes.Connection, subjectDid *alaTypes.Did, psmHashes []common.Address) []*alaTypes.PSMHashStatus {
	statues, _ := alastria.GetSubjectCredentialsStatus(actorClient, subjectDid, psmHashes)
	return statues
}
