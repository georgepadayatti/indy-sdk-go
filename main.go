package main

import (
	"log"

	"github.com/georgepadayatti/indy-wrapper/wallet"
)

func main() {
	res := wallet.CreateWallet("abc", "abc")
	log.Printf("Status for indy wallet creation: %v\n", res.Err)
}
