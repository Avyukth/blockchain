package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func run() error {

	//ethereum crypto
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.PublicKey

	addreess := crypto.PubkeyToAddress(publicKey).String()

	fmt.Println(addreess)
	return nil

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
