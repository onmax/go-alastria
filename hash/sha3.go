package hash

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

func Sha3(input string) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(input))
	buf := hash.Sum(nil)
	return hexutil.Encode(buf)
}
