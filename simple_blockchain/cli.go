package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct{}

func (cli *CLI) Run() {
	cli.validateArgs()

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createWallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	createBlockchainAddr := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	getBalanceAddr := getBalanceCmd.String("address", "", "The address to get balance for")

	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendAmount := sendCmd.Int("amount", 0, "The amount of coin to send")

	switch os.Args[1] {
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		handleError(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		handleError(err)
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		handleError(err)
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		handleError(err)
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		handleError(err)
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		handleError(err)
	default:
		cli.printUsage()
		os.Exit(0)
	}

	if sendCmd.Parsed() {
		if *sendTo == "" || *sendFrom == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddr == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddr)
	}

	if createWalletCmd.Parsed() {
		cli.createWallet()
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses()
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddr == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddr)
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("   createblockchain --address ADDRESS - add a block to the blockchain, send the reward to ADDRESS")
	fmt.Println("   printchain - print all the blocks of the blockchain")
	fmt.Println("  	createwallet - generate a new key pair and save it to wallet file")
	fmt.Println("   listaddresses - list all wallet addresses")
	fmt.Println("   send --from FROM --to TO --amount AMOUNT - send AMOUNT of coins from FROM to TO")
	fmt.Println("   getbalance --address ADDRESS - get the balance of ADDRESS")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(0)
	}
}
