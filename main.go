package main

import (
	"log"

	"github.com/georgepadayatti/indy-sdk-go/wallet"
)

func main() {
	res := wallet.CreateWallet("abc", "abc")
	log.Printf("Status for indy wallet creation: %v\n", res.Err)
}
