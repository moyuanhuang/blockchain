package main

func (cli *CLI) printChain() {
	bc := NewBlockChain()
	bc.PrintChain()
}
