package main 

import (
	"log"
	"fmt"
	"flag"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

var existing string
func init() {
	flag.StringVar(&existing, "decode", "", "Existing WIF")
	flag.Parse()
}

func main() {

	var sk *btcec.PrivateKey
	var wif *btcutil.WIF
	var err error

	if existing != "" {
		wif, err = btcutil.DecodeWIF(existing)
		if err != nil {
			log.Fatal(err)
		}
		sk = wif.PrivKey
	} else {
		sk, err = btcec.NewPrivateKey(btcec.S256())
		if err != nil {
			log.Fatal(err)
		}
		wif, err = btcutil.NewWIF(sk, &chaincfg.MainNetParams, false)
		if err != nil {
			log.Fatal(err)
		}
	}

	addr, err := btcutil.NewAddressPubKey(sk.PubKey().SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WIF: ", wif)
	fmt.Println("ADDRESS: ", addr.EncodeAddress())
}