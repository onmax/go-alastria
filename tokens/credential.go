package tokens

import "fmt"

type Credential struct {
	Header  *Header            `json:"header"`
	Payload *CredentialPayload `json:"payload"`
}

type CredentialPayload struct {
	JSONTokenId          string               `json:"jti"`
	IssuedAt             uint64               `json:"iat"`
	ExpiresAt            uint64               `json:"exp"`
	NotBefore            uint64               `json:"nbf"`
	Issuer               string               `json:"iss"`
	Subject              string               `json:"sub"`
	VerifiableCredential *CredentialPayloadVC `json:"vc"`
}

type CredentialPayloadVC struct {
	Contexts          []string                `json:"@context"`
	Types             []string                `json:"type"`
	CredentialSubject *map[string]interface{} `json:"credentialSubject"`
}

var defaultCredentialTypes = [2]string{"VerifiableCredential", "AlastriaVerifiableCredential"}
var defaultCredentialContextURLs = [2]string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"}

// Validates the VerifiableCredential(aka Credential) according to the specification
// Header: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Payload: https://github.com/alastria/alastria-identity/wiki/Alastria-DID-Method-Specification-(Quorum-version)#3-credentials
// The validation with timestamp will be done with the machine timestamp. This can cause a problem, if the time is not syncronize.
// Sets default values if they are empty and they are required
// Returns an error if a mandatory field is empty
// Mandatory fields are: payload.AlastriaToken, payload.CreateAlastriaTX and payload.PublicKey
func CreateCredential(header *Header, payload *CredentialPayload) (*Credential, error) {
	err := ValidateCredential(header, payload)

	if err != nil {
		return nil, err
	}

	return &Credential{
			Header:  header,
			Payload: payload},
		nil
}

// Validates the Credential according to the specification
// https://github.com/alastria/alastria-identity/wiki/Alastria-DID-Method-Specification-(Quorum-version)#3-credentials
// Sets default values if they are empty and they are required
func ValidateCredential(header *Header, payload *CredentialPayload) error {

	if err := ValidateHeader(header); err != nil {
		return err
	}

	mandatoryStringValues := map[string]string{
		"Issuer":  payload.Issuer,
		"Subject": payload.Subject,
	}

	if err := CheckMandatoryStringFieldsAreNotEmpty(mandatoryStringValues); err != nil {
		return err
	}

	if payload.VerifiableCredential == nil {
		return fmt.Errorf(emptyPayloadField, "VerifiableCredential")
	}

	if payload.VerifiableCredential.CredentialSubject == nil {
		return fmt.Errorf(emptyPayloadField, "CredentialSubject")
	}

	if err := CheckLevelOfAssurance(payload.VerifiableCredential.CredentialSubject); err != nil {
		return err
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.VerifiableCredential.Types = AddDefaultValues(payload.VerifiableCredential.Types, defaultCredentialTypes[:])
	payload.VerifiableCredential.Contexts = AddDefaultValues(payload.VerifiableCredential.Contexts, defaultCredentialContextURLs[:])

	// TODO Check iss and sub are valid
	return nil
}
