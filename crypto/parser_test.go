package crypto

import (
	"crypto/ecdsa"
	"reflect"
	"testing"
)

func TestHexToECDSAPub(t *testing.T) {
	type args struct {
		_pub string
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PublicKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToECDSAPub(tt.args._pub)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToECDSAPub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexToECDSAPub() = %v, want %v", got, tt.want)
			}
		})
	}
}
