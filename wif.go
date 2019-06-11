package main 

import (
	"log"
	"fmt"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

func main() {
	sk, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}
	wif, err := btcutil.NewWIF(sk, &chaincfg.MainNetParams, false)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := btcutil.NewAddressPubKey(sk.PubKey().SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WIF: ", wif)
	fmt.Println("ADDRESS: ", addr.EncodeAddress())
}