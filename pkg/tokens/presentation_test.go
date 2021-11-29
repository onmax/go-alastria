package tokens

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestCreatePresentation(t *testing.T) {
	type args struct {
		header  *Header
		payload *PresentationPayload
	}
	tests := []struct {
		name             string
		args             args
		want             *Presentation
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
				payload: &PresentationPayload{
					JSONTokenId:                    "json-token-id",
					ExpiresAt:                      TimestampInTheFuture,
					NotBefore:                      _CurrentTimestamp,
					IssuedAt:                       _CurrentTimestamp,
					Issuer:                         "issuer",
					Audience:                       "audience",
					PresentationRequestJSONTokenId: "pr-jti",
					VerifiablePresentation: &PresentationPayloadVP{
						Contexts:              []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
						Types:                 []string{"VerifiablePresentation", "AlastriaVerifiablePresentation"},
						ProcessHash:           "process-hash",
						ProcessUrl:            "process-url",
						ProcessDescription:    "process-description",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
				},
			},
			want: &Presentation{
				Header: &Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &PresentationPayload{
					JSONTokenId:                    "json-token-id",
					IssuedAt:                       _CurrentTimestamp,
					ExpiresAt:                      TimestampInTheFuture,
					NotBefore:                      _CurrentTimestamp,
					Issuer:                         "issuer",
					Audience:                       "audience",
					PresentationRequestJSONTokenId: "pr-jti",
					VerifiablePresentation: &PresentationPayloadVP{
						Contexts:              []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
						Types:                 []string{"VerifiablePresentation", "AlastriaVerifiablePresentation"},
						ProcessHash:           "process-hash",
						ProcessUrl:            "process-url",
						ProcessDescription:    "process-description",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test with optional values",
			args: args{
				header: &Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PresentationPayload{
					Issuer:   "issuer",
					Audience: "audience",
					VerifiablePresentation: &PresentationPayloadVP{
						ProcessHash:           "process-hash",
						ProcessUrl:            "process-url",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
					IssuedAt: _CurrentTimestamp,
				},
			},
			want: &Presentation{
				Header: &Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &PresentationPayload{
					IssuedAt: _CurrentTimestamp,
					Issuer:   "issuer",
					Audience: "audience",
					VerifiablePresentation: &PresentationPayloadVP{
						Contexts:              []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
						Types:                 []string{"VerifiablePresentation", "AlastriaVerifiablePresentation"},
						ProcessHash:           "process-hash",
						ProcessUrl:            "process-url",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test without Issuer which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &PresentationPayload{
					Audience: "audience",
					VerifiablePresentation: &PresentationPayloadVP{
						ProcessHash:           "process-hash",
						ProcessUrl:            "process-url",
						ProcessDescription:    "process-description",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Issuer is empty and it is mandatory",
		},
		{
			name: "test without Audience which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &PresentationPayload{
					Issuer: "Issuer",
					VerifiablePresentation: &PresentationPayloadVP{
						ProcessHash:           "process-hash",
						ProcessUrl:            "process-url",
						ProcessDescription:    "process-description",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Audience is empty and it is mandatory",
		},
		{
			name: "test without VerifiablePresentation which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &PresentationPayload{
					Issuer:   "Issuer",
					Audience: "audience",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value VerifiablePresentation is empty and it is mandatory",
		},
		{
			name: "test without ProcessHash which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &PresentationPayload{
					Issuer:   "Issuer",
					Audience: "audience",
					VerifiablePresentation: &PresentationPayloadVP{
						ProcessUrl:            "process-url",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value ProcessHash is empty and it is mandatory",
		},
		{
			name: "test without ProcessUrl which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &PresentationPayload{
					Issuer:   "Issuer",
					Audience: "audience",
					VerifiablePresentation: &PresentationPayloadVP{
						ProcessHash:           "process-hash",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value ProcessUrl is empty and it is mandatory",
		},
		{
			name: "test without VerifiableCredentials which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &PresentationPayload{
					Issuer:   "Issuer",
					Audience: "audience",
					VerifiablePresentation: &PresentationPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value VerifiableCredentials is empty and it is mandatory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreatePresentation(tt.args.header, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePresentation() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
			if tt.args.payload.IssuedAt == 0 {
				tt.args.payload.IssuedAt = _CurrentTimestamp
			}
			if !reflect.DeepEqual(got, tt.want) {
				a, _ := json.Marshal(got)
				b, _ := json.Marshal(tt.want)
				t.Errorf("CreatePresentation() -> %s \n\tgot:  %s,\n\twant: %s", tt.name, string(a), string(b))
			}
		})
	}
}
