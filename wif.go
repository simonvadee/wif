package main 

import (
	"log"
	"fmt"
	"crypto/elliptic"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

func main() {
	sk, err := btcec.NewPrivateKey(elliptic.P256())
	if err != nil {
		log.Fatal(err)
	}
	wif, err := btcutil.NewWIF(sk, &chaincfg.MainNetParams, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WIF: ", wif)
}