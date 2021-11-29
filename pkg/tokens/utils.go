package tokens

import (
	"fmt"
)

type Header struct {
	Algorithm    string `json:"alg,omitempty"`
	Type         string `json:"typ,omitempty"`
	JSONWebToken string `json:"jwk,omitempty"`
	KeyID        string `json:"kid,omitempty"`
}

// Checks if list contains the given string
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var emptyPayloadField = "the value %s is empty and it is mandatory"

var validHeaderTypes = [...]string{"JWT"}
var validHeaderAlgorithms = [...]string{"ES256K"}

// Validates that the header is valid following the specification found here:
// https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Sets default values to header.Type and header.Algorithm if they are empty.
// If header.Type or header.Algorithm are invalid also throws an error.
func ValidateHeader(header *Header) error {

	if header == nil {
		return fmt.Errorf("header is nil")
	}

	// If header.Type is empty, then sets a default value. Otherwise, checks if header.Type is valid.
	if header.Type == "" {
		header.Type = validHeaderTypes[0]
	} else if !stringInSlice(header.Type, validHeaderTypes[:]) {
		return fmt.Errorf("invalid Type equals to %s in header. Use: JWT", header.Type)
	}

	// If header.Algorithm is empty, then sets a default value. Otherwise, checks if header.Algorithm is valid.
	if header.Algorithm == "" {
		header.Algorithm = validHeaderAlgorithms[0]
	} else if !stringInSlice(header.Algorithm, validHeaderAlgorithms[:]) {
		return fmt.Errorf("invalid Algorithm equals to %s in header. Use: ES256K", header.Algorithm)
	}

	// TODO Check if header.KeyID is valid
	// TODO Check if header.JSONWebToken is valid

	return nil
}

func checkMandatoryStringFieldsAreNotEmpty(values map[string]string) error {
	for k, v := range values {
		if v == "" {
			return fmt.Errorf(emptyPayloadField, k)
		}
	}
	return nil
}

func checkMandatoryStringArrayFieldsNotEmpty(values map[string][]string) error {
	for k, v := range values {
		if len(v) == 0 {
			return fmt.Errorf(emptyPayloadField, k)
		}
	}
	return nil
}

func checkMandatoryStructFieldsAreNotEmpty(values map[string]interface{}) error {
	for k, v := range values {
		if v == nil {
			return fmt.Errorf(emptyPayloadField, k)
		}
	}
	return nil
}

// Adds default value in values array if they are not set
func addDefaultValues(values []string, defaultValues []string) []string {
	for _, _defaultValue := range defaultValues {
		if !stringInSlice(_defaultValue, values) {
			values = append(values, _defaultValue)
		}
	}
	return values
}

// Validates that the elements in values are in the validValues array
func validateEnum(values []string, validValues []string, field string) error {
	for _, value := range values {
		if !stringInSlice(value, validValues) {
			return fmt.Errorf("the value %s=%s is invalid. Only the following values are accepted %s", field, value, validValues)
		}
	}
	return nil
}

func checkLevelOfAssurance(_data *map[string]interface{}) error {
	if _data == nil {
		return fmt.Errorf("arg is nil")
	}

	data := *_data
	loaI := data["levelOfAssurance"]
	if loaI == nil {
		return fmt.Errorf("levelOfAssurance is empty")
	}
	loa, ok := loaI.(int)
	if !ok {
		return fmt.Errorf("levelOfAssurance is not a int")
	}
	if loa > 3 {
		return fmt.Errorf("levelOfAssurance is invalid. Only 0, 1, 2, 3 are valid")
	}
	return nil
}
