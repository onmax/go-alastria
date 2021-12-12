package credentials

import (
	alaTypes "github.com/onmax/go-alastria/types"
)

func GetCredentialStatus(conn *alaTypes.Connection, signedCredential string, subject *alaTypes.Did) (*alaTypes.PSMHashStatuses, error) {
	issuerStatus, err := GetIssuerCredentialStatus(conn, signedCredential)
	if err != nil {
		return nil, err
	}

	subjectStatus, err := GetSubjectCredentialStatus(conn, subject, signedCredential)
	if err != nil {
		return nil, err
	}

	var status uint8
	if issuerStatus.Status < subjectStatus.Status {
		status = issuerStatus.Status
	} else {
		status = subjectStatus.Status
	}

	if err != nil {
		return nil, err
	}

	return &alaTypes.PSMHashStatuses{
		IssuerPSMHash:  issuerStatus.PSMHash,
		SubjectPSMHash: subjectStatus.PSMHash,
		Credential:     signedCredential,
		Status:         status,
	}, nil

}

func GetCredentialsStatus(conn *alaTypes.Connection, signedCredentials []string, subject *alaTypes.Did) ([]*alaTypes.PSMHashStatuses, error) {
	var statuses []*alaTypes.PSMHashStatuses
	for _, signedCredential := range signedCredentials {
		status, err := GetCredentialStatus(conn, signedCredential, subject)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}
