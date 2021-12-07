package tokens

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	alaTypes "github.com/onmax/go-alastria/types"
)

func TestCreateAlastriaIdentityCreation(t *testing.T) {
	type args struct {
		header  *alaTypes.Header
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
				header: &alaTypes.Header{
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
				Header: &alaTypes.Header{
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
				header: &alaTypes.Header{
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
				Header: &alaTypes.Header{
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
				header: &alaTypes.Header{
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
				header: &alaTypes.Header{
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
				header: &alaTypes.Header{
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
				header: &alaTypes.Header{
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

func TestDecodeAIC(t *testing.T) {
	type args struct {
		signedAIC string
	}
	tests := []struct {
		name    string
		args    args
		want    *AIC
		wantErr bool
	}{
		{
			name: "Should decode an AIC successfully",
			args: args{
				signedAIC: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsInB1YmxpY0tleSI6InB1YmxpYy1rZXkiLCJqdGkiOiJqc29uLXRva2VuLWlkIiwiY3JlYXRlQWxhc3RyaWFUWCI6ImNyZWF0ZS1hbGFzdHJpYS10eCIsImFsYXN0cmlhVG9rZW4iOiJhbGFzdHJpYS10b2tlbiIsIkBjb250ZXh0IjpbImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2FydGlmYWN0cy92MSJdLCJ0eXBlIjpbIkFsYXN0cmlhSWRlbnRpdHlDcmVhdGlvbiIsIlVTMTIiXX0.Fjjy_24sitUYVpLnsQkqet7nl6zb70wg7A2QvZAkHO0cCZGQzoTOA4CAvTNn2-b0QPR_vgoNJ--PuZlxBEBBBg",
			},
			want: &AIC{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &AICPayload{
					IssuedAt:         1,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeAIC(tt.args.signedAIC)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeAIC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeAIC() = %v, want %v", got, tt.want)
			}
		})
	}
}
