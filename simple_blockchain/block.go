package main

import (
    "time"
)

type Block struct {
    PrevHash []byte
    Hash []byte
    Data []byte
    Timestamp int64
    Nonce uint
}

func NewBlock(data string, prevHash []byte) *Block {
    newBlock := &Block{
        PrevHash: prevHash,
        Data: []byte(data),
        Timestamp: time.Now().Unix(),
        Nonce: 0,
    }

    pow := NewProofOfWork(newBlock)
    nonce, hash := pow.Run()

    newBlock.Hash = hash[:]
    newBlock.Nonce = nonce

    return newBlock
}

func NewGenesisBlock() *Block {
    return NewBlock("Genesis Block", []byte{})
}
