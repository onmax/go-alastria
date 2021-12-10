package did

import (
	"errors"
	"strings"

	alaTypes "github.com/onmax/go-alastria/types"
)

func NewDid(network, networkId, proxyAddress string) *alaTypes.Did {
	// TODO return error
	if network == "" || networkId == "" || proxyAddress == "" {
		return nil
	}

	return &alaTypes.Did{
		Protocol:     "ala",
		Network:      network,
		NetworkId:    networkId,
		ProxyAddress: proxyAddress,
	}
}

func NewDidFromString(did string) (*alaTypes.Did, error) {
	parts := strings.Split(did, ":")
	if len(parts) != 5 {
		return nil, errors.New("invalid did")
	}

	return &alaTypes.Did{
		Protocol:     parts[1],
		Network:      parts[2],
		NetworkId:    parts[3],
		ProxyAddress: parts[4],
	}, nil
}
