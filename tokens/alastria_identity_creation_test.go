package tokens

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCreateAlastriaIdentityCreation(t *testing.T) {
	type args struct {
		header  *Header
		payload *AICPayload
	}
	tests := []struct {
		name             string
		args             args
		want             *AIC
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
				payload: &AICPayload{
					IssuedAt:         _CurrentTimestamp,
					ExpiresAt:        TimestampInTheFuture,
					NotBefore:        _CurrentTimestamp,
					PublicKey:        "public-key",
					JSONTokenId:      "json-token-id",
					CreateAlastriaTX: "create-alastria-tx",
					AlastriaToken:    "alastria-token",
					Contexts:         []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:            []string{"AlastriaIdentityCreation", "US12"},
				},
			},
			want: &AIC{
				Header: &Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &AICPayload{
					IssuedAt:         _CurrentTimestamp,
					ExpiresAt:        TimestampInTheFuture,
					NotBefore:        _CurrentTimestamp,
					PublicKey:        "public-key",
					JSONTokenId:      "json-token-id",
					CreateAlastriaTX: "create-alastria-tx",
					AlastriaToken:    "alastria-token",
					Contexts:         []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:            []string{"AlastriaIdentityCreation", "US12"},
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
				payload: &AICPayload{
					IssuedAt:         _CurrentTimestamp,
					PublicKey:        "public-key",
					CreateAlastriaTX: "create-alastria-tx",
					AlastriaToken:    "alastria-token",
				},
			},
			want: &AIC{
				Header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				Payload: &AICPayload{
					IssuedAt:         _CurrentTimestamp,
					PublicKey:        "public-key",
					CreateAlastriaTX: "create-alastria-tx",
					AlastriaToken:    "alastria-token",
					Contexts:         []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:            []string{"AlastriaIdentityCreation"},
				},
			},
			wantErr: false,
		},
		{
			name: "test without AlastriaToken which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &AICPayload{
					IssuedAt:         _CurrentTimestamp,
					PublicKey:        "public-key",
					CreateAlastriaTX: "create-alastria-tx",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value AlastriaToken is empty and it is mandatory",
		},
		{
			name: "test without PublicKey which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &AICPayload{
					AlastriaToken:    "alastria-token",
					CreateAlastriaTX: "create-alastria-tx",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value PublicKey is empty and it is mandatory",
		},
		{
			name: "test without CreateAlastriaTX which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &AICPayload{
					AlastriaToken: "alastria-token",
					PublicKey:     "public-key",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value CreateAlastriaTX is empty and it is mandatory",
		},
		{
			name: "test invalid type",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &AICPayload{
					AlastriaToken:    "alastria-token",
					PublicKey:        "public-key",
					CreateAlastriaTX: "create-alastria-tx",
					Types:            []string{"AlastriaIdentityCreation", "US12", "invalid"},
				},
			},
			wantErr:          true,
			errorMsgContains: fmt.Sprintf("the value Types=invalid is invalid. Only the following values are accepted %s", validAICTypes),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAlastriaIdentityCreation(tt.args.header, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAlastriaIdentityCreation() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
			if tt.args.payload.IssuedAt == 0 {
				tt.args.payload.IssuedAt = _CurrentTimestamp
			}
			if !reflect.DeepEqual(got, tt.want) {
				a, _ := json.Marshal(got)
				b, _ := json.Marshal(tt.want)
				t.Errorf("CreateAlastriaIdentityCreation() -> %s \n\tgot:  %s,\n\twant: %s", tt.name, string(a), string(b))
			}
		})
	}
}
