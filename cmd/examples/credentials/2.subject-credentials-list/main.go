package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/onmax/go-alastria/alastria"
	exampleutil "github.com/onmax/go-alastria/cmd/examples"
)

var subjectAddress = common.HexToAddress("d0a0d5a1310a715157c3f81b789d6d9dc447aef5")

// More information about this process in the ./README.md
// To simplify the code, errors are not being checked
func main() {
	fmt.Printf("\n------ Any actor get the psm hashes of credentials of a subject ------ \n\n")

	actorConf := exampleutil.GetReaderClientConf()
	actorClient, _ := alastria.NewClient(actorConf)
	psmHashes, _ := alastria.GetSubjectCredentialList(actorClient, subjectAddress)

	fmt.Printf("The subject %s has %v credentials registered in the blockchain:\n", subjectAddress, len(psmHashes))
	fmt.Printf("The PSMHashes are %v\n", psmHashes)
}
