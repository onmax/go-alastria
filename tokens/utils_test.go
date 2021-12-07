package tokens

import (
	"reflect"
	"strings"
	"testing"

	alaTypes "github.com/onmax/go-alastria/types"
)

func TestStringInSlice(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{a: "test", list: []string{"test", "test1"}},
			want: true,
		},
		{
			name: "not in the array",
			args: args{a: "test", list: []string{"test1"}},
			want: false,
		},
		{
			name: "with empty array",
			args: args{a: "test", list: []string{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateHeader(t *testing.T) {
	type args struct {
		header *alaTypes.Header
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		expectedHeader   *alaTypes.Header
		errorMsgContains string
	}{
		{
			name:             "header is nil",
			args:             args{header: nil},
			wantErr:          true,
			errorMsgContains: "header is nil",
		},
		{
			name: "without default types",
			args: args{
				header: &alaTypes.Header{},
			},
			wantErr: false,
			expectedHeader: &alaTypes.Header{
				Algorithm: "ES256K",
				Type:      "JWT",
			},
		},
		{
			name: "invalid type",
			args: args{
				header: &alaTypes.Header{
					Type: "invalid",
				},
			},
			wantErr:          true,
			errorMsgContains: "invalid Type equals to invalid in header. Use: JWT",
		},
		{
			name: "invalid algorithm",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "invalid",
				},
			},
			wantErr:          true,
			errorMsgContains: "invalid Algorithm equals to invalid in header. Use: ES256K",
		},
		{
			name: "valid",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
			},
			wantErr: false,
			expectedHeader: &alaTypes.Header{
				Algorithm: "ES256K",
				Type:      "JWT",
			},
		},
		{
			name: "valid with all values",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
			},
			wantErr: false,
			expectedHeader: &alaTypes.Header{
				Algorithm:    "ES256K",
				Type:         "JWT",
				KeyID:        "key-id",
				JSONWebToken: "json-web-token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHeader(tt.args.header); (err != nil) != tt.wantErr {
				t.Errorf("ValidateHeader() error = %v", err)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateHeader() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}

			if tt.expectedHeader != nil && !reflect.DeepEqual(tt.args.header, tt.expectedHeader) {
				t.Errorf("ValidateHeader() = %v, want %v", tt.args.header, tt.expectedHeader)
			}
		})
	}
}

func TestCheckMandatoryStringFieldsNotEmpty(t *testing.T) {
	type args struct {
		values map[string]string
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorMsgContains string
	}{
		{
			name: "empty values",
			args: args{
				values: map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "missing values",
			args: args{
				values: map[string]string{
					"field_1": "",
					"field_2": "",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value field_1 is empty and it is mandatory",
		},
		{
			name: "missing values 2",
			args: args{
				values: map[string]string{
					"field_1": "some value",
					"field_2": "",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value field_2 is empty and it is mandatory",
		},
		{
			name: "valid",
			args: args{
				values: map[string]string{
					"field_1": "some value",
					"field_2": "some value",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkMandatoryStringFieldsAreNotEmpty(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTimestamps() error = %v", err)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
		})
	}
}

func TestCheckMandatoryStringArrayFieldsNotEmpty(t *testing.T) {
	type args struct {
		values map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkMandatoryStringArrayFieldsNotEmpty(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("CheckMandatoryStringArrayFieldsNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddDefaultValues(t *testing.T) {
	type args struct {
		values        []string
		defaultValues []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDefaultValues(tt.args.values, tt.args.defaultValues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDefaultValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateEnum(t *testing.T) {
	type args struct {
		values      []string
		validValues []string
		field       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateEnum(tt.args.values, tt.args.validValues, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("ValidateEnum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckMandatoryStringFieldsAreNotEmpty(t *testing.T) {
	type args struct {
		values map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkMandatoryStringFieldsAreNotEmpty(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("CheckMandatoryStringFieldsAreNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckMandatoryStructFieldsAreNotEmpty(t *testing.T) {
	type args struct {
		values map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkMandatoryStructFieldsAreNotEmpty(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("CheckMandatoryStructFieldsAreNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckLevelOfAssurance(t *testing.T) {
	type args struct {
		data *map[string]interface{}
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		errorMsgContains string
	}{
		{
			name:             "arg is nil",
			args:             args{},
			wantErr:          true,
			errorMsgContains: "arg is nil",
		},
		{
			name:             "levelOfAssurance is empty",
			args:             args{data: &map[string]interface{}{"other-key": 10}},
			wantErr:          true,
			errorMsgContains: "levelOfAssurance is empty",
		},
		{
			name:             "levelOfAssurance is negative",
			args:             args{data: &map[string]interface{}{"levelOfAssurance": -1}},
			wantErr:          true,
			errorMsgContains: "levelOfAssurance is not a uint8",
		},
		{
			name:             "levelOfAssurance is not a uint8",
			args:             args{data: &map[string]interface{}{"levelOfAssurance": 10}},
			wantErr:          true,
			errorMsgContains: "levelOfAssurance is not a uint8",
		},
		{
			name:             "levelOfAssurance is not in the valid range",
			args:             args{data: &map[string]interface{}{"levelOfAssurance": 4}},
			wantErr:          true,
			errorMsgContains: "levelOfAssurance is invalid. Only 0, 1, 2, 3 are valid",
		},
		{
			name:             "levelOfAssurance is valid",
			args:             args{data: &map[string]interface{}{"levelOfAssurance": 0}},
			wantErr:          false,
			errorMsgContains: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkLevelOfAssurance(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CheckLevelOfAssurance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
