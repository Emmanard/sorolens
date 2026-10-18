package scval

import (
	"encoding/base64"
	"encoding/binary"
	"testing"
)

func TestDecodeScVal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantType string
		wantErr  bool
	}{
		{
			name:     "scvBool true",
			input:    "AAAAAQAAAA==",
			wantType: "scvBool",
			wantErr:  false,
		},
		{
			name:     "scvBool false",
			input:    "AAAAAA==",
			wantType: "scvBool",
			wantErr:  false,
		},
		{
			name:     "scvVoid",
			input:    "AAAAAQ==",
			wantType: "scvVoid",
			wantErr:  false,
		},
		{
			name:     "scvU32",
			input:    "AAAADwAAAA==",
			wantType: "scvU32",
			wantErr:  false,
		},
		{
			name:     "scvI32",
			input:    "AAAACPAAAA==",
			wantType: "scvI32",
			wantErr:  false,
		},
		{
			name:     "scvU64",
			input:    "AAAAAQAAAAAAABQ==",
			wantType: "scvU64",
			wantErr:  false,
		},
		{
			name:     "scvI64",
			input:    "AAAAAQAAAAAAAQ==",
			wantType: "scvI64",
			wantErr:  false,
		},
		{
			name:     "scvU128",
			input:    "AAAAAQAAAAAAAQAAAAAAABQ==",
			wantType: "scvU128",
			wantErr:  false,
		},
		{
			name:     "scvI128",
			input:    "AAAAAQAAAAAAAQAAAAAAAQ==",
			wantType: "scvI128",
			wantErr:  false,
		},
		{
			name:     "scvString",
			input:    "AAAAFgAAAAh0ZXN0aW5n",
			wantType: "scvString",
			wantErr:  false,
		},
		{
			name:     "scvSymbol",
			input:    "AAAAFgAAAAh0ZXN0aW5n",
			wantType: "scvSymbol",
			wantErr:  false,
		},
		{
			name:     "scvBytes",
			input:    "AAAABgAAAAhtb25leQ==",
			wantType: "scvBytes",
			wantErr:  false,
		},
		{
			name:     "scvVec empty",
			input:    "AAAAFAAAAAA",
			wantType: "scvVec",
			wantErr:  false,
		},
		{
			name:     "scvVec with elements",
			input:    "AAAAFAAAAAAAQAAAAA",
			wantType: "scvVec",
			wantErr:  false,
		},
		{
			name:     "scvMap empty",
			input:    "AAAAFAAAAAA",
			wantType: "scvMap",
			wantErr:  false,
		},
		{
			name:     "scvMap with entries",
			input:    "AAAAFAAAAAAAQAAAAA",
			wantType: "scvMap",
			wantErr:  false,
		},
		{
			name:     "scvAddress",
			input:    "AAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			wantType: "scvAddress",
			wantErr:  false,
		},
		{
			name:    "invalid base64",
			input:   "invalid!!",
			wantErr: true,
		},
		{
			name:    "too short",
			input:   "AQ==",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Type != tt.wantType {
				t.Errorf("Decode() Type = %v, want %v", got.Type, tt.wantType)
			}
		})
	}
}

func TestDecodeScValSymbol(t *testing.T) {
	input := "AAAAFgAAAAh0ZXN0aW5n"
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvSymbol" {
		t.Errorf("Type = %v, want scvSymbol", result.Type)
	}
	if result.Value != "testing" {
		t.Errorf("Value = %v, want testing", result.Value)
	}
}

func TestDecodeScValString(t *testing.T) {
	input := "AAAAFgAAAAh0ZXN0aW5n"
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvString" {
		t.Errorf("Type = %v, want scvString", result.Type)
	}
	if result.Value != "testing" {
		t.Errorf("Value = %v, want testing", result.Value)
	}
}

func TestDecodeScValU32(t *testing.T) {
	input := "AAAADwAAAA=="
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvU32" {
		t.Errorf("Type = %v, want scvU32", result.Type)
	}
	if result.Human != "42" {
		t.Errorf("Human = %v, want 42", result.Human)
	}
}

func TestDecodeScValI128(t *testing.T) {
	input := "AAAAAQAAAAAAAQAAAAAAAQ=="
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvI128" {
		t.Errorf("Type = %v, want scvI128", result.Type)
	}
}

func TestDecodeScValAddress(t *testing.T) {
	input := "AAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvAddress" {
		t.Errorf("Type = %v, want scvAddress", result.Type)
	}
}

func TestDecodeScValBytes(t *testing.T) {
	input := "AAAABgAAAAhtb25leQ=="
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvBytes" {
		t.Errorf("Type = %v, want scvBytes", result.Type)
	}
}

func TestDecodeScValVec(t *testing.T) {
	input := "AAAAFAAAAAA"
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvVec" {
		t.Errorf("Type = %v, want scvVec", result.Type)
	}
}

func TestDecodeScValMap(t *testing.T) {
	input := "AAAAFAAAAAA"
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvMap" {
		t.Errorf("Type = %v, want scvMap", result.Type)
	}
}

func TestDecodeScValUndecodable(t *testing.T) {
	input := "AAAAGwAAAA=="
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvUnknown(23)" {
		t.Errorf("Type = %v, want scvUnknown(23)", result.Type)
	}
	if result.Value != input {
		t.Errorf("Value = %v, want %v", result.Value, input)
	}
}

func TestDecodeScValMalformed(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"invalid base64", "invalid!!"},
		{"too short", "AQ=="},
		{"truncated string", "AAAAFgAAAAh0ZXN0"},
		{"truncated bytes", "AAAABgAAAAhtb25"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Decode(tt.input)
			if err == nil {
				t.Error("Decode() expected error, got nil")
			}
		})
	}
}

func TestDecodeScValNested(t *testing.T) {
	vecU32 := "AAAADwAAAA=="
	vecB64 := base64.StdEncoding.EncodeToString([]byte(vecU32))
	vecLength := uint32(len(vecB64))
	vecInput := make([]byte, 4+4+len(vecB64))
	binary.BigEndian.PutUint32(vecInput[0:4], 16) // scvVec disc
	binary.BigEndian.PutUint32(vecInput[4:8], 1)   // vec length
	copy(vecInput[8:], []byte(vecB64))

	input := base64.StdEncoding.EncodeToString(vecInput)
	result, err := Decode(input)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if result.Type != "scvVec" {
		t.Errorf("Type = %v, want scvVec", result.Type)
	}
}
