package main

import (
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address string) {
	if !ValidateAddress(address) {
		log.Panic("Invalid wallet address!")
	}

	bc := NewBlockChain()

	balance := 0
	UTXOs := bc.FindUTXO(address)
	for _, output := range UTXOs {
		balance += output.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}
