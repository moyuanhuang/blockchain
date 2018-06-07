package main

import (
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("Invalid from address!")
	}

	if !ValidateAddress(to) {
		log.Panic("Invalid to address!")
	}

	bc := NewBlockChain()
	tx := NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}
