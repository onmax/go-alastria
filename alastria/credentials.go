package alastria

import (
	"github.com/onmax/go-alastria/credentials"
	alaTypes "github.com/onmax/go-alastria/types"
)

func AddSubjectCredentials(client *alaTypes.Connection, signedCredentials []string, subjectDid *alaTypes.Did, URI string) ([]string, []string, error) {
	// TODO Return also the status of the psmHash
	return credentials.AddSubjectCredentials(client, signedCredentials, subjectDid, URI)
}

func AddIssuerCredentials(client *alaTypes.Connection, signedCredentials []string, entityDid *alaTypes.Did) ([]string, []string, error) {
	// TODO Return also the status of the psmHash
	return credentials.AddIssuerCredentials(client, signedCredentials, entityDid)
}

func GetCredentials(conn *alaTypes.Connection, signedCredentials []string, subjectDid *alaTypes.Did) ([]*alaTypes.PSMHashStatuses, error) {
	return credentials.GetCredentialsStatus(conn, signedCredentials, subjectDid)

}
