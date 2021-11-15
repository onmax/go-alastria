package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	identity "github.com/onmax/go-alastria/contracts/alastria-identity-manager"
	pkr "github.com/onmax/go-alastria/contracts/alastria-public-key-registry"
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
	// fmt.Sprintf("calling createEntity2AlastriaID\n")
	// createEntity2AlastriaID(client, chainId)
	tx := CreateAlastriaIdentity(client)
	aic := signAIC(*tx)
	tx2 := PrepareAlastriaId(client)

	if tx == nil {
		fmt.Println("No tx")
	}
	if tx2 == nil {
		fmt.Println("No tx2")
	}

	fmt.Println("Sending tx2")
	err = client.SendTransaction(context.Background(), tx2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sending tx")
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("---------\ntx2 %v\n---------\n", tx2.Hash().Hex())
	fmt.Printf("---------\ntx %v\n---------\n", tx.Hash().Hex())
	fmt.Printf("---------\naic %v\n---------\n", aic)
}

func signAIC(tx types.Transaction) string {
	_, privKeyEntity1, pubKeyEntity1, err := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
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

	signedAT, err := crypto.Sign(at, privKeyEntity1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nSigned AT %s\n\n", signedAT)
	}

	_, pk2, pubk2, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")

	b, _ := tx.MarshalBinary()
	createAlastriaIdTx := "0x" + hex.EncodeToString(b)

	aic := tokens.AIC{
		Header: &tokens.Header{
			Algorithm:    "ES256K",
			Type:         "JWT",
			KeyID:        "key-id",
			JSONWebToken: "json-web-token",
		},
		Payload: &tokens.AICPayload{
			IssuedAt:         1,
			PublicKey:        pubk2,
			JSONTokenId:      "json-token-id",
			CreateAlastriaTX: createAlastriaIdTx,
			AlastriaToken:    signedAT,
			Contexts:         []string{"https://alastria.github.io/identity/artifacts/v1"},
			Types:            []string{"AlastriaIdentityCreation", "US12"},
		},
	}
	aicSigned, _ := crypto.Sign(aic, pk2)
	return aicSigned
}

// func CreateAlastriaIdentity(client bind.ContractBackend) *types.Transaction {
// 	contract, err := pkr.NewAlastriaContracts(common.HexToAddress("0x017B0F93d1F31666Ae0bfb35f2e35bC1b0AC43C1"), client)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("Connected to public Key Registry")
// 	}
// 	k, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
// 	opts, _ := getTxOpt(k.PrivateKey)
// 	opts.NoSend = true

// 	_, _, pubKey, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")
// 	tx, err := contract.AddKey(opts, pubKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Tx", tx)

// 	_identity, _ := identity.NewAlastriaContracts(common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"), client)
// 	t2, _ := _identity.CreateAlastriaIdentity(opts, tx.Data())
// 	return t2
// }

// func PrepareAlastriaId(client bind.ContractBackend) *types.Transaction {
// 	conn, err := identity.NewAlastriaContracts(common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"), client)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println(conn.Version(&bind.CallOpts{}))
// 	}
// 	_, _, pubk2, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")
// 	k, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
// 	opts, _ := getTxOpt(k.PrivateKey)
// 	opts.NoSend = true
// 	tx, _ := conn.CreateAlastriaIdentity(opts, []byte(pubk2))

// 	return tx
// }

func CreateAlastriaIdentity(client *ethclient.Client) *types.Transaction {
	contract, err := pkr.NewAlastriaContracts(common.HexToAddress("0x017B0F93d1F31666Ae0bfb35f2e35bC1b0AC43C1"), client)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to public Key Registry")
	}
	k, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
	opts, _ := getTxOpt(k.PrivateKey)
	opts.NoSend = true

	k2, _, pubKey, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")
	ttx, err := contract.AddKey(opts, pubKey)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	ttx_, err := types.SignTx(ttx, types.NewEIP155Signer(chainID), k2.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tx", ttx_)

	_identity, _ := identity.NewAlastriaContracts(common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"), client)
	opts, _ = getTxOpt(k.PrivateKey)
	opts.NoSend = true
	t2, _ := _identity.CreateAlastriaIdentity(opts, ttx_.Data())
	chainID, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedt2, err := types.SignTx(t2, types.NewEIP155Signer(chainID), k2.PrivateKey)
	return signedt2
}

func PrepareAlastriaId(client *ethclient.Client) *types.Transaction {
	conn, err := identity.NewAlastriaContracts(common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"), client)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(conn.Version(&bind.CallOpts{}))
	}
	_, _, pubk2, _ := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")
	k, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
	opts, _ := getTxOpt(k.PrivateKey)
	opts.NoSend = true
	tx, _ := conn.CreateAlastriaIdentity(opts, []byte(pubk2))
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), k.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx
}

// func AddKeys(client bind.ContractBackend) {
// 	_identity, _ := identity.NewAlastriaContracts(common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"), client)
// 	// _, _, _, _ := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
// 	opts := &bind.CallOpts{
// 		Pending: true,
// 		Context: context.Background(),
// 	}
// 	addrs, _ := _identity.IdentityKeys(opts, common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"))
// 	s, _ := addrs.MarshalText()
// 	fmt.Printf("-------- %s\n\n", addrs)
// 	fmt.Printf("normal address %v\n\n", s)

// }

// func createEntity2AlastriaID(client bind.ContractBackend, chainId *big.Int) {
// 	conn, err := identity.NewAlastriaContracts(common.HexToAddress("0x3bFe9aAFc360a31D4060ef3E8cb5013Da197015a"), client)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println(conn.Version(&bind.CallOpts{}))
// 	}

// 	k, privKeyEntity1, pubKeyEntity1, err := keystore.ImportKs("../assets/keystores/entity1-a9728125c573924b2b1ad6a8a8cd9bf6858ced49.json", "Passw0rd")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("privKeyEntity1: %s, pubKeyEntity1: %s\n", privKeyEntity1, pubKeyEntity1)

// 	fmt.Printf("--->\n")
// 	_, _, pubKey, err := keystore.ImportKs("../assets/keystores/entity2-ad88f1a89cf02a32010b971d8c8af3a2c7b3bd94.json", "Passw0rd")
// 	fmt.Printf("public key of entity 2 %s\n", pubKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("creating new widentity \n")

// 	tx, err := conn.CreateAlastriaIdentity(
// 		&bind.TransactOpts{
// 			From:     k.Address,
// 			GasLimit: 600000,
// 			GasPrice: big.NewInt(0),
// 			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
// 				fmt.Printf("signing with addresses \n%s\n%s  \n", address.Hex(), k.Address.Hex())
// 				if address != k.Address {
// 					return nil, errors.New("not authorized to sign this account")
// 				}
// 				fmt.Printf("address signing %s\n", address)

// 				signer := types.HomesteadSigner{}
// 				signature, err := ecr.Sign(signer.Hash(tx).Bytes(), k.PrivateKey)
// 				fmt.Printf("signature %x\n", signature)
// 				if err != nil {
// 					return nil, err
// 				}
// 				ttx, err := tx.WithSignature(signer, signature)
// 				fmt.Printf("ttx %s\n", ttx)
// 				return ttx, err
// 			},
// 		}, []byte(pubKey))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Tx", tx)
// 	r, err := conn.IsIdentityIssuer(&bind.CallOpts{}, k.Address)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Identity 1 is identity issuer", r)
// }

func getTxOpt(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return nil, err
	}
	opts.Value = big.NewInt(0)
	opts.GasLimit = 600000
	opts.GasPrice = big.NewInt(0)
	fmt.Println("opts %s", opts)
	return opts, nil
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
