package main

import (
    "encoding/gob"
    "crypto/sha256"
    "time"
    "bytes"
    "log"
)

type Block struct {
    PrevHash []byte
    Hash []byte
    Transactions []*Transaction
    Timestamp int64
    Nonce int
}

func NewBlock(txs []*Transaction, prevHash []byte) *Block {
    start := time.Now()
    newBlock := &Block{
        PrevHash: prevHash,
        Transactions: txs,
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

func NewGenesisBlock(coinbase *Transaction) *Block {
    return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
    var result bytes.Buffer
    enc := gob.NewEncoder(&result)

    err := enc.Encode(b)
    handleError(err)

    return result.Bytes()
}

func DeserializeBlock(data []byte) *Block {
    var block Block
    dec := gob.NewDecoder(bytes.NewReader(data))

    err := dec.Decode(&block)
    handleError(err)
    return &block
}

func (b *Block) HashTransactions() []byte {
    var txHashes [][]byte
    var txHash [32]byte

    for _, tx := range b.Transactions {
        txHashes = append(txHashes, tx.ID)
    }

    txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

    return txHash[:]
}

func (b *Block) IsGenesisBlock() bool{
    return len(b.PrevHash) == 0
}
