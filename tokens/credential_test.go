package tokens

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestCreateCredential(t *testing.T) {
	type args struct {
		header  *Header
		payload *CredentialPayload
	}
	tests := []struct {
		name             string
		args             args
		want             *Credential
		wantErr          bool
		errorMsgContains string
	}{
		{
			name: "test with all values",
			args: args{
				header: &Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &CredentialPayload{
					JSONTokenId: "json-token-id",
					IssuedAt:    _CurrentTimestamp,
					ExpiresAt:   TimestampInTheFuture,
					NotBefore:   _CurrentTimestamp,
					Issuer:      "issuer",
					Subject:     "subject",
					VerifiableCredential: &CredentialPayloadVC{
						Contexts: []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
						Types:    []string{"VerifiableCredential", "AlastriaVerifiableCredential"},
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
							"data":             map[string]interface{}{"field_name": "field_value"},
						},
					},
				},
			},
			want: &Credential{
				Header: &Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &CredentialPayload{
					JSONTokenId: "json-token-id",
					IssuedAt:    _CurrentTimestamp,
					ExpiresAt:   TimestampInTheFuture,
					NotBefore:   _CurrentTimestamp,
					Issuer:      "issuer",
					Subject:     "subject",
					VerifiableCredential: &CredentialPayloadVC{
						Contexts: []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
						Types:    []string{"VerifiableCredential", "AlastriaVerifiableCredential"},
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
							"data":             map[string]interface{}{"field_name": "field_value"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test with optional values",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
					VerifiableCredential: &CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
						},
					},
					IssuedAt: _CurrentTimestamp,
				},
			},
			want: &Credential{
				Header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				Payload: &CredentialPayload{
					IssuedAt: _CurrentTimestamp,
					Issuer:   "issuer",
					Subject:  "subject",
					VerifiableCredential: &CredentialPayloadVC{
						Types:    defaultCredentialTypes[:],
						Contexts: defaultCredentialContextURLs[:],
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test invalid levelOfAssurance",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
					VerifiableCredential: &CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 4,
						},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "levelOfAssurance is invalid. Only 0, 1, 2, 3 are valid",
		},
		{
			name: "test without Subject which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer: "issuer",
					VerifiableCredential: &CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
						},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Subject is empty and it is mandatory",
		},
		{
			name: "test without VerifiableCredential which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value VerifiableCredential is empty and it is mandatory",
		},
		{
			name: "test without CredentialSubject which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:               "issuer",
					Subject:              "subject",
					VerifiableCredential: &CredentialPayloadVC{},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value CredentialSubject is empty and it is mandatory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateCredential(tt.args.header, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCredential() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
			if tt.args.payload.IssuedAt == 0 {
				tt.args.payload.IssuedAt = _CurrentTimestamp
			}
			if !reflect.DeepEqual(got, tt.want) {
				a, _ := json.Marshal(got)
				b, _ := json.Marshal(tt.want)
				t.Errorf("CreateCredential() -> %s \n\tgot:  %s,\n\twant: %s", tt.name, string(a), string(b))
			}
		})
	}
}
