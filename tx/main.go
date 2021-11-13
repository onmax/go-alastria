package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ecr "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	"github.com/onmax/go-alastria/crypto"
	"github.com/onmax/go-alastria/keystore"
	"github.com/onmax/go-alastria/tokens"
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
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(chainId)
	}
	// _, pr, pubKeyEntity1, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
	// fmt.Println(pubKeyEntity1, pr)
	fmt.Sprintf("calling createEntity2AlastriaID\n")
	createEntity2AlastriaID(client, chainId)
}

func createEntity2AlastriaID(client bind.ContractBackend, chainId *big.Int) {
	conn, err := identity.NewAlastriaContracts(common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"), client)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(conn.Version(&bind.CallOpts{}))
	}

	k, privKeyEntity1, pubKeyEntity1, err := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("privKeyEntity1: %s, pubKeyEntity1: %s\n", privKeyEntity1, pubKeyEntity1)

	didEntity1 := "did:ala:quor:redT:d2f868f056ef3a48bbc8d446dfed411e9bf93ab0"
	providerURL := "https://regular.telsius.blockchainbyeveris.io:2000"
	callbackUrl := "https://serviceprovider.alastria.blockchainbyeveris.io/api/login/"
	networkId := "Alastria network"
	exp := uint64(2036798552159)
	kidCredential := "did:ala:quor:redt:12eeaCCA9eEbB78eB97d7cac6b#keys-1"
	iat := uint64(1636798552159)
	jti := "ze298y42sba"
	at := tokens.AT{
		Header: &tokens.Header{
			Algorithm:    "ES256K",
			Type:         "JWT",
			KeyID:        kidCredential,
			JSONWebToken: pubKeyEntity1,
		},
		Payload: &tokens.ATPayload{
			IssuedAt:          iat,
			ExpiresAt:         exp,
			Issuer:            didEntity1,
			AlastriaNetworkId: networkId,
			CallbackURL:       callbackUrl,
			GatewayURL:        providerURL,
			JSONTokenId:       jti,
			Types:             []string{"AlastriaToken"},
		},
	}

	signed, err := crypto.Sign(at, privKeyEntity1)

	fmt.Sprintf("signed %s\n", signed)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nSigned AT %s\n\n", signed)
	}

	fmt.Printf("--->\n")
	_, _, pubKey, err := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")
	fmt.Printf("public key of entity 2 %s\n", pubKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("creating new widentity \n")

	tx, err := conn.CreateAlastriaIdentity(
		&bind.TransactOpts{
			From:     k.Address,
			GasLimit: 600000,
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				fmt.Printf("signing with addresses \n%s\n%s  \n", address.Hex(), k.Address.Hex())
				if address != k.Address {
					return nil, errors.New("not authorized to sign this account")
				}
				fmt.Printf("address signing %s\n", address)

				signer := types.HomesteadSigner{}
				signature, err := ecr.Sign(signer.Hash(tx).Bytes(), k.PrivateKey)
				fmt.Printf("signature %x\n", signature)
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
		}, []byte(pubKey))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tx", tx)
	// r, err := conn.IsIdentityIssuer(&bind.CallOpts{}, k.Address)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Identity 1 is identity issuer", r)
}

// func createSubject() {
// 	account, err := keystore.CreateKs("./", "password")
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("Subject address", account.Address)
// 	}

// 	aic := tokens.AIC{
// 		Header: &tokens.Header{
// 			Algorithm:    "ES256K",
// 			Type:         "JWT",
// 			KeyID:        "key-id",
// 			JSONWebToken: "json-web-token",
// 		},
// 		Payload: &tokens.AICPayload{
// 			IssuedAt:         1,
// 			PublicKey:        "public-key",
// 			JSONTokenId:      "json-token-id",
// 			CreateAlastriaTX: "create-alastria-tx",
// 			AlastriaToken:    "alastria-token",
// 			Contexts:         []string{"https://alastria.github.io/identity/artifacts/v1"},
// 			Types:            []string{"AlastriaIdentityCreation", "US12"},
// 		},
// 	}
// }
