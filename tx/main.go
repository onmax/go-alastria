package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/owenyuwono/eth-contract/api"
)

func main() {

	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(header.Number.String())
	}
	b, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(b)
	}

	conn, err := api.NewApi(common.HexToAddress("CONTRACT_ADDRESS"), client)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(conn)
	}
}
