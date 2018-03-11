package main

import (
    "os"
    "fmt"
    "flag"
)

type CLI struct {
    *BlockChain
}

func (cli *CLI) Run() {
    cli.validateArgs()

    addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

    addBlockData := addBlockCmd.String("data", "", "Block Data")

    switch os.Args[1] {
    case "addblock":
        err := addBlockCmd.Parse(os.Args[2:])
        handleError(err)
    case "printchain":
        err := printChainCmd.Parse(os.Args[2:])
        handleError(err)
    default:
        cli.printUsage()
        os.Exit(0)
    }

    if addBlockCmd.Parsed() {
        if *addBlockData == "" {
            addBlockCmd.Usage()
            os.Exit(0)
        }
        cli.AddBlock(*addBlockData)
    }

    if printChainCmd.Parsed() {
        cli.PrintChain()
    }
}

func (cli *CLI) printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
    fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
    if len(os.Args) < 2 {
        cli.printUsage()
        os.Exit(0)
    }
}
