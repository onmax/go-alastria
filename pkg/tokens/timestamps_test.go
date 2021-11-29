package tokens

import (
	"fmt"
	"strings"
	"testing"
)

func _MemoryAddress(x uint64) *uint64 { return &x }

var MsInADay uint64 = 3600 * 1000 * 24 // 1 day in milliseconds
var _CurrentTimestamp = CurrentTimestamp()
var TimestampInTheFuture = _CurrentTimestamp + MsInADay
var TimestampInThePast = _CurrentTimestamp - MsInADay

func TestValidateTimestamps(t *testing.T) {
	type args struct {
		iat *uint64
		exp *uint64
		nbf *uint64
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorMsgContains string
	}{
		{
			name:             "nbf is after exp",
			args:             args{iat: _MemoryAddress(TimestampInThePast), nbf: _MemoryAddress(TimestampInTheFuture + 100), exp: _MemoryAddress(TimestampInTheFuture)},
			wantErr:          true,
			errorMsgContains: fmt.Sprintf("the value NotBefore must be before ExpiresAt. ExpiresAt: %v. NotBefore: %v", TimestampInTheFuture, TimestampInTheFuture+100),
		},
		{
			name:             "all fields are ok",
			args:             args{iat: _MemoryAddress(TimestampInThePast), exp: _MemoryAddress(TimestampInTheFuture), nbf: _MemoryAddress(TimestampInThePast)},
			wantErr:          false,
			errorMsgContains: "",
		},
		{
			name:             "without any fields",
			args:             args{iat: _MemoryAddress(0), exp: _MemoryAddress(0), nbf: _MemoryAddress(0)},
			wantErr:          false,
			errorMsgContains: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTimestamps(tt.args.iat, tt.args.exp, tt.args.nbf); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTimestamps() error = %v", err)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
		})
	}
}

func TestValidateIssuedAt(t *testing.T) {
	type args struct {
		iat *uint64
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorMsgContains string
	}{
		{
			name:             "iat is nil",
			args:             args{iat: nil},
			wantErr:          true,
			errorMsgContains: "the value IssuedAt is empty and it is mandatory",
		},
		{
			name:             "iat is zero",
			args:             args{iat: _MemoryAddress(0)},
			wantErr:          false,
			errorMsgContains: "",
		},
		{
			name:             "iat is not in the future",
			args:             args{iat: _MemoryAddress(TimestampInTheFuture)},
			wantErr:          true,
			errorMsgContains: "the value IssuedAt is greater than the current timestamp",
		},
		{
			name:    "iat is ok",
			args:    args{iat: _MemoryAddress(CurrentTimestamp())},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shouldChange := tt.args.iat != nil && *tt.args.iat == 0

			if err := ValidateIssuedAt(tt.args.iat); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTimestamps() error = %v", err)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}

			if shouldChange && *tt.args.iat == 0 {
				t.Error("ValidateIssuedAt() should set a value when iat is 0")
			}
		})
	}
}

func TestValidateExpiresAt(t *testing.T) {
	type args struct {
		exp *uint64
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorMsgContains string
	}{
		{
			name:             "exp is nil",
			args:             args{exp: nil},
			wantErr:          true,
			errorMsgContains: "the value ExpiresAt is empty and it is mandatory",
		},
		{
			name:             "exp is zero",
			args:             args{exp: _MemoryAddress(0)},
			wantErr:          false,
			errorMsgContains: "",
		},
		{
			name:             "exp is in the past",
			args:             args{exp: _MemoryAddress(TimestampInThePast)},
			wantErr:          true,
			errorMsgContains: "the value ExpiresAt is less than the current timestamp",
		},
		{
			name:    "exp is in the future",
			args:    args{exp: _MemoryAddress(TimestampInTheFuture)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateExpiresAt(tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTimestamps() error = %v", err)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
		})
	}
}

func TestValidateNotBefore(t *testing.T) {
	type args struct {
		nbf *uint64
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorMsgContains string
	}{
		{
			name:             "nbf is nil",
			args:             args{nbf: nil},
			wantErr:          true,
			errorMsgContains: "the value NotBefore is empty and it is mandatory",
		},
		{
			name:    "nbf is zero",
			args:    args{nbf: _MemoryAddress(0)},
			wantErr: false,
		},
		{
			name:             "nbf is in the past",
			args:             args{nbf: _MemoryAddress(TimestampInThePast)},
			wantErr:          false,
			errorMsgContains: "",
		},
		{
			name:             "nbf is in the future",
			args:             args{nbf: _MemoryAddress(TimestampInTheFuture)},
			wantErr:          false,
			errorMsgContains: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNotBefore(tt.args.nbf); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTimestamps() error = %v", err)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
		})
	}
}
