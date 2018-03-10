package main

import (
    "encoding/gob"
    "time"
    "bytes"
    "log"
)

type Block struct {
    PrevHash []byte
    Hash []byte
    Data []byte
    Timestamp int64
    Nonce uint
}

func NewBlock(data string, prevHash []byte) *Block {
    start := time.Now()
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

    log.Printf("New block added. used %v\n", time.Now().Sub(start))
    return newBlock
}

func NewGenesisBlock() *Block {
    return NewBlock("Genesis Block", []byte{})
}

func (b *Block) Serialize() []byte {
    var result bytes.Buffer
    enc := gob.NewEncoder(&result)

    err := enc.Encode(b)
    handleError(err)

    return result.Bytes()
}
