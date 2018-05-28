package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		log.Panic("Address is not valid!")
	}
	CreateBlockchain(address)
	fmt.Println("Done!")
}
