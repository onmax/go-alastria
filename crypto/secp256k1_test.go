package tokens

import (
	"fmt"
	"testing"
)

func Test_sign(t *testing.T) {
	type args struct {
		jwt        AlastriaIdentityCreation
		privateKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Should sign an AIC successfully",
			args:    args{jwt: AIC{&AICHeader{Algorithm: "ES256K", Type: "JWT", JSONWebToken: "jwk", KeyID: "kid"}, &AICPayload{Context: []string{"context"}, Types: []string{"types"}, CreateAlastriaTX: "createAlastriaTX", AlastriaToken: "alastriaToken", PublicKey: "publicKey", JSONTokenId: "jti", IssuedAt: 0, ExpiresAt: 0, NotBefore: 0}}, privateKey: "8dd5e4759a2d8a0ebeb1a1c53b113dad453daaf140b47c419eeef11a46d656b3"},
			want:    "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqd2siLCJraWQiOiJraWQifQ.eyJAY29udGV4dCI6WyJjb250ZXh0Il0sInR5cGUiOlsidHlwZXMiXSwiY3JlYXRlQWxhc3RyaWFUWCI6ImNyZWF0ZUFsYXN0cmlhVFgiLCJhbGFzdHJpYVRva2VuIjoiYWxhc3RyaWFUb2tlbiIsInB1YmxpY0tleSI6InB1YmxpY0tleSIsImp0aSI6Imp0aSIsImlhdCI6MCwiZXhwIjowLCJuYmYiOjB9.T7b5cXzLo71_oux43s0uKK2yUPsr7GCB7BpNEvL6q9gN72F_mOtnOjzk7NroeD4K61DCkSgp8-F_d1xL5NKDsA",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sign(tt.args.jwt, tt.args.privateKey)
			fmt.Printf("%v\n", got)

			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		signed string
		_pk    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Should verify a JWT successfully",
			args:    args{signed: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqd2siLCJraWQiOiJraWQifQ.eyJAY29udGV4dCI6WyJjb250ZXh0Il0sInR5cGUiOlsidHlwZXMiXSwiY3JlYXRlQWxhc3RyaWFUWCI6ImNyZWF0ZUFsYXN0cmlhVFgiLCJhbGFzdHJpYVRva2VuIjoiYWxhc3RyaWFUb2tlbiIsInB1YmxpY0tleSI6InB1YmxpY0tleSIsImp0aSI6Imp0aSIsImlhdCI6MCwiZXhwIjowLCJuYmYiOjB9.T7b5cXzLo71_oux43s0uKK2yUPsr7GCB7BpNEvL6q9gN72F_mOtnOjzk7NroeD4K61DCkSgp8-F_d1xL5NKDsA", _pk: "61a909c125d095e3a4f67125144427312ea6a42fb6321e94f41b25921e106bb8a5db57cc767c179dfe14fa959b615b2cf60852430867ce4675d18a1b3aa032b8"},
			wantErr: false,
		},
		{
			name:    "Should verify a JWT successfully signed from Java",
			args:    args{signed: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqd2siLCJraWQiOiJraWQifQ.eyJleHAiOjAsIm5iZiI6MCwianRpIjoianRpIiwiaWF0IjowLCJ0eXBlIjpbInR5cGVzIl0sImNyZWF0ZUFsYXN0cmlhVFgiOiJjcmVhdGVBbGFzdHJpYVRYIiwiYWxhc3RyaWFUb2tlbiI6ImFsYXN0cmlhVG9rZW4iLCJwdWJsaWNLZXkiOiJwdWJsaWNLZXkiLCJAY29udGV4dCI6WyJjb250ZXh0Il19.KTmgyC3rhqoOk9QOTjQsVLku-r6f5JqDF1UAldu813RdJ63UOPc5f2BKryMB-mdxxuh1WbspsB4zRYTb_ALzVw", _pk: "61a909c125d095e3a4f67125144427312ea6a42fb6321e94f41b25921e106bb8a5db57cc767c179dfe14fa959b615b2cf60852430867ce4675d18a1b3aa032b8"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Verify(tt.args.signed, tt.args._pk); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
