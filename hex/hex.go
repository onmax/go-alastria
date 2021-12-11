package hex

import (
	"encoding/hex"
	ehex "encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
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

func PublicKeyToAddress(publicKeyStr string) (common.Address, error) {
	publicKey, err := HexToByteArr(publicKeyStr) // TODO Sanitize removing 0x and 03 || 04
	if err != nil {
		return common.Address{}, err
	}
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[:])
	buf = hash.Sum(nil)
	publicAddress := hexutil.Encode(buf[12:])

	return common.HexToAddress(publicAddress), nil
}

func StringTo32ByteArray(input string) [32]byte {
	var output [32]byte
	copy(output[:], input)
	return output
}
