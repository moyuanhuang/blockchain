package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"crypto/elliptic"
	"encoding/gob"
)

// WALLETFILE is the file that stores the wallets data
const WALLETFILE = "wallet.dat"

// Wallets stores a collection of wallets
type Wallets struct {
	Wallets map[string]*Wallet
}

// NewWallets creates Wallets and fills it from a file if it exists
func NewWallets() (*Wallets, error) {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()
	return wallets, err
}

// CreateWallet adds a Wallet to Wallets
func (wls *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	wls.Wallets[address] = wallet
	return address
}

// GetAddresses returns an array of addresses stored in the wallet file
func (wls *Wallets) GetAddresses() []string {
	addrs := make([]string, 0)
	for address := range wls.Wallets {
		addrs = append(addrs, address)
	}
	return addrs
}

// GetWallet returns a Wallet by its address
func (wls *Wallets) GetWallet(addr string) *Wallet {
	return wls.Wallets[addr]
}

// LoadFromFile loads wallets from the file
func (wls *Wallets) LoadFromFile() error {
	if _, err := os.Stat(WALLETFILE); os.IsNotExist(err) {
		return nil
	}

	fileContent, err := ioutil.ReadFile(WALLETFILE)
	handleError(err)

	var wallets Wallets
	// need to Register the possible value of any interface in the decoded object
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	handleError(err)

	wls.Wallets = wallets.Wallets
	return nil
}

// SaveToFile saves wallets to a file
func (wls *Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())
	// encoder initialization takes in "Where to dump", and then during encoding
	// takes in "What to encode". HOWEVER, decoder initialization takes in "What
	// to decode", and during decoding takes in "Where to dump"
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(wls)
	handleError(err)

	err = ioutil.WriteFile(WALLETFILE, content.Bytes(), 0644)
	handleError(err)
}
