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
var testnet bool
func init() {
	flag.StringVar(&existing, "decode", "", "Existing WIF")
	flag.BoolVar(&testnet, "testnet", false, "Use testnet configuration")
	flag.Parse()
}

func main() {
	var sk *btcec.PrivateKey
	var wif *btcutil.WIF
	var err error
	var netParams = &chaincfg.MainNetParams

	if testnet {
		netParams = &chaincfg.TestNet3Params
	}

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
		wif, err = btcutil.NewWIF(sk, netParams, true)
		if err != nil {
			log.Fatal(err)
		}
	}

	var pubKey []byte
	switch wif.CompressPubKey {
	case true:
		pubKey = sk.PubKey().SerializeCompressed()
	case false:
		pubKey = sk.PubKey().SerializeUncompressed()
	}
	addr, err := btcutil.NewAddressPubKey(pubKey, netParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WIF: ", wif)
	fmt.Println("ADDRESS: ", addr.EncodeAddress())
}