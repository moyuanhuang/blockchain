package main

type BlockChain struct {
    blocks []*Block
}

func NewBlockChain() *BlockChain {
    return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string) {
    prevBlock := bc.blocks[len(bc.blocks) - 1]
    newblock := NewBlock(data, prevBlock.Hash)
    bc.blocks = append(bc.blocks, newblock)
}
