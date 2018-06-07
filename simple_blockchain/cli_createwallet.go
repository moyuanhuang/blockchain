package main

import (
	"fmt"
)

func (cli *CLI) createWallet() {
	wallets, err := NewWallets()
	handleError(err)
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new wallet address: %s\n", address)
}
