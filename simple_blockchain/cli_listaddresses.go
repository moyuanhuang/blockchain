package main

import (
	"fmt"
)

func (cli *CLI) listAddresses() {
	wls, err := NewWallets()
	handleError(err)

	addresses := wls.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
