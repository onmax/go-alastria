package hex

import (
	"encoding/hex"
	ehex "encoding/hex"
)

func Remove0x(input string) string {
	if len(input) >= 2 && input[0:2] == "0x" {
		return input[2:]
	}
	return input
}

func Add0x(input string) string {
	if len(input) > 0 && input[:2] != "0x" {
		return "0x" + input
	}
	return input
}

// Converts byte array to hex string
func ByteArrToHex(input []byte) string {
	return ehex.EncodeToString(input)
}

// Converts input hex string to byte array
func HexToByteArr(input string) ([]byte, error) {
	decodedByteArray, err := hex.DecodeString(Remove0x(input))
	if err != nil {
		return nil, err
	}
	return decodedByteArray, nil
}
