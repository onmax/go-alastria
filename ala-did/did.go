package did

import "fmt"

type Did struct {
	Protocol     string
	Network      string
	NetworkId    string
	ProxyAddress string
}

func NewDid(network, networkId, proxyAddress string) *Did {
	if network == "" || networkId == "" || proxyAddress == "" {
		return nil
	}

	return &Did{
		Protocol:     "ala",
		Network:      network,
		NetworkId:    networkId,
		ProxyAddress: proxyAddress,
	}
}

func (d *Did) String() string {
	return fmt.Sprintf("did:%s:%s:%s:%s", d.Protocol, d.Network, d.NetworkId, d.ProxyAddress)
}
