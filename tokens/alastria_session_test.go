package tokens

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCreateAlastriaSession(t *testing.T) {
	type args struct {
		header  *Header
		payload *ASPayload
	}
	tests := []struct {
		name             string
		args             args
		want             *AS
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
				payload: &ASPayload{
					IssuedAt:      _CurrentTimestamp,
					ExpiresAt:     TimestampInTheFuture,
					NotBefore:     _CurrentTimestamp,
					JSONTokenId:   "json-token-id",
					Issuer:        "issuer",
					AlastriaToken: "alastria-token",
					Contexts:      []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:         []string{"AlastriaSession", "US211"},
				},
			},
			want: &AS{
				Header: &Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &ASPayload{
					IssuedAt:      _CurrentTimestamp,
					ExpiresAt:     TimestampInTheFuture,
					NotBefore:     _CurrentTimestamp,
					JSONTokenId:   "json-token-id",
					Issuer:        "issuer",
					AlastriaToken: "alastria-token",
					Contexts:      []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:         []string{"AlastriaSession", "US211"},
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
				payload: &ASPayload{
					IssuedAt:      _CurrentTimestamp,
					Issuer:        "issuer",
					AlastriaToken: "alastria-token",
				},
			},
			want: &AS{
				Header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				Payload: &ASPayload{
					IssuedAt:      _CurrentTimestamp,
					Issuer:        "issuer",
					AlastriaToken: "alastria-token",
					Contexts:      []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:         []string{"AlastriaSession"},
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
				payload: &ASPayload{
					Issuer: "issuer",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value AlastriaToken is empty and it is mandatory",
		},
		{
			name: "test without Issuer which is mandatory",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &ASPayload{
					AlastriaToken: "alastria-token",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Issuer is empty and it is mandatory",
		},
		{
			name: "test invalid type",
			args: args{
				header: &Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &ASPayload{
					AlastriaToken: "alastria-token",
					Issuer:        "issuer",
					Types:         []string{"AlastriaSession", "US221", "invalid"},
				},
			},
			wantErr:          true,
			errorMsgContains: fmt.Sprintf("the value Types=invalid is invalid. Only the following values are accepted %s", validASTypes),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAlastriaSession(tt.args.header, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAlastriaSession() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
			if tt.args.payload.IssuedAt == 0 {
				tt.args.payload.IssuedAt = _CurrentTimestamp
			}
			if !reflect.DeepEqual(got, tt.want) {
				a, _ := json.Marshal(got)
				b, _ := json.Marshal(tt.want)
				t.Errorf("CreateAlastriaSession() -> %s \n\tgot:  %s,\n\twant: %s", tt.name, string(a), string(b))
			}
		})
	}
}
