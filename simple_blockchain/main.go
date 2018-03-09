package main

import (
    "fmt"
)

func main() {
    blockChain := NewBlockChain()

    blockChain.AddBlock("This is the first block")
    blockChain.AddBlock("This is the second block")

    for _, b := range blockChain.blocks {
        fmt.Printf("%s\n%x\n%x\n\n", b.Data, b.PrevHash, b.Hash)
    }
}
