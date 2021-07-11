package tokens

type AIC struct {
	Header  *Header     `json:"header"`
	Payload *AICPayload `json:"payload"`
}

type AICPayload struct {
	IssuedAt         uint64   `json:"iat"`
	ExpiresAt        uint64   `json:"exp"`
	NotBefore        uint64   `json:"nbf"`
	PublicKey        string   `json:"publicKey"`
	JSONTokenId      string   `json:"jti"`
	CreateAlastriaTX string   `json:"createAlastriaTX"` // ! Not the best name
	AlastriaToken    string   `json:"alastriaToken"`
	Contexts         []string `json:"@context"`
	Types            []string `json:"type"`
}

var defaultAICContextURLs = [...]string{"https://alastria.github.io/identity/artifacts/v1"}
var defaultAICType = [...]string{"AlastriaIdentityCreation"}
var validAICTypes = append([]string{"US12"}, defaultAICType[:]...)

// Validates the AIC according to the specification
// Header: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Payload: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#03-alastria-identity-creation-aic
// The validation with timestamp will be done with the machine timestamp. This can cause a problem, if the time is not syncronize.
// Sets default values if they are empty and they are required
// Returns an error if a mandatory field is empty
// Mandatory fields are: payload.AlastriaToken, payload.CreateAlastriaTX and payload.PublicKey
func CreateAlastriaIdentityCreation(header *Header, payload *AICPayload) (*AIC, error) {
	err := ValidateAlastriaIdentityCreation(header, payload)

	if err != nil {
		return nil, err
	}

	return &AIC{
			Header:  header,
			Payload: payload},
		nil
}

// Validates the AIC according to the specification
// https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#03-alastria-identity-creation-aic
// Sets default values if they are empty and they are required
func ValidateAlastriaIdentityCreation(header *Header, payload *AICPayload) error {

	if err := ValidateHeader(header); err != nil {
		return err
	}

	mandatoryValues := map[string]string{
		"AlastriaToken":    payload.AlastriaToken,
		"CreateAlastriaTX": payload.CreateAlastriaTX,
		"PublicKey":        payload.PublicKey,
	}

	if err := CheckMandatoryStringFieldsAreNotEmpty(mandatoryValues); err != nil {
		return err
	}

	// ! In the docs, it is not explicitly mentioned that the payload.Types can have more values that the ones that they mention
	if err := ValidateEnum(payload.Types, validAICTypes, "Types"); err != nil {
		return err
	}
	// TODO validAICTypes should be a oneOfs

	// ! In the docs, it is not explicitly mentioned that the payload.Contexts can have more values that the ones that they mention
	if err := ValidateEnum(payload.Contexts, defaultAICContextURLs[:], "Context"); err != nil {
		return err
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.Types = AddDefaultValues(payload.Types, defaultAICType[:])
	payload.Contexts = AddDefaultValues(payload.Contexts, defaultAICContextURLs[:])

	// TODO Validates contexts are URLs
	// TODO CreateAlastria, TXPublicKey and JSONTokenId can be a regex?

	return nil
}
