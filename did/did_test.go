package did

import (
	"reflect"
	"testing"
)

func TestNewDid(t *testing.T) {
	type args struct {
		network      string
		networkId    string
		proxyAddress string
	}
	tests := []struct {
		name string
		args args
		want *Did
	}{
		{
			name: "did ok",
			args: args{
				network:      "test",
				networkId:    "test",
				proxyAddress: "test",
			},
			want: &Did{
				Protocol:     "ala",
				Network:      "test",
				NetworkId:    "test",
				ProxyAddress: "test",
			},
		},
		{
			name: "did not ok",
			args: args{
				network:      "",
				networkId:    "",
				proxyAddress: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDid(tt.args.network, tt.args.networkId, tt.args.proxyAddress); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDid_String(t *testing.T) {
	type fields struct {
		Protocol     string
		Network      string
		NetworkId    string
		ProxyAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "did ok",
			fields: fields{
				Protocol:     "ala",
				Network:      "test",
				NetworkId:    "test",
				ProxyAddress: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Did{
				Protocol:     tt.fields.Protocol,
				Network:      tt.fields.Network,
				NetworkId:    tt.fields.NetworkId,
				ProxyAddress: tt.fields.ProxyAddress,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("Did.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
