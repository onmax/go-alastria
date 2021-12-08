package hex

import (
	"reflect"
	"testing"
)

func TestRemove0x(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no 0x",
			args: args{input: "1234"},
			want: "1234",
		},
		{
			name: "0x",
			args: args{input: "0x123"},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove0x(tt.args.input); got != tt.want {
				t.Errorf("Remove0x() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd0x(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no 0x",
			args: args{input: "1234"},
			want: "0x1234",
		},
		{
			name: "0x",
			args: args{input: "0x123"},
			want: "0x123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add0x(tt.args.input); got != tt.want {
				t.Errorf("Add0x() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteArrToHex(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteArrToHex(tt.args.input); got != tt.want {
				t.Errorf("ByteArrToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToByteArr(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToByteArr(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToByteArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexToByteArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
