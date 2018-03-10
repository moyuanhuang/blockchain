package main

import (

)

func main() {
    blockChain := NewBlockChain()

    blockChain.AddBlock("This is the first block")
    blockChain.AddBlock("This is the second block")
}
