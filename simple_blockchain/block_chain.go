package main

import (
    "fmt"

    "github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const BlocksBucket = "blocks"

var lastHashKey = []byte("l")

type BlockChain struct {
    lastHash []byte
    db *bolt.DB
}

type BlockChainIterator struct {
    curHash []byte
    db *bolt.DB
}

func NewBlockChain() *BlockChain {
    db, err := bolt.Open(dbFile, 0600, nil)
    handleError(err)

    var lastHash []byte

    err = db.Update(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BlocksBucket))
        if bucket == nil {
            fmt.Printf("No existing block chain, creating a new one with Genesis block...\n")
            gb := NewGenesisBlock()

            bucket, err = tx.CreateBucket([]byte(BlocksBucket))
            handleError(err)

            err = bucket.Put(lastHashKey, gb.Hash)
            handleError(err)

            err = bucket.Put(gb.Hash, gb.Serialize())
            handleError(err)
        } else {
            lastHash = bucket.Get(lastHashKey)
        }
        return nil
    })
    handleError(err)

    return &BlockChain{lastHash, db}
}

func (bc *BlockChain) AddBlock(data string) {
    var lastHash []byte

    err := bc.db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BlocksBucket))
        lastHash = bucket.Get(lastHashKey)

        return nil
    })
    handleError(err)

    newBlock := NewBlock(data, lastHash)

    err = bc.db.Update(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BlocksBucket))
        err := bucket.Put(newBlock.Hash, newBlock.Serialize())
        handleError(err)

        err = bucket.Put(lastHashKey, newBlock.Hash)
        handleError(err)

        bc.lastHash = newBlock.Hash
        return nil
    })

}

func (bc *BlockChain) Iterator() *BlockChainIterator {
    return &BlockChainIterator{bc.lastHash, bc.db}
}

func (it *BlockChainIterator) Next() *Block {
    var block *Block
    block = nil

    err := it.db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BlocksBucket))
        data := bucket.Get(it.curHash)
        if data != nil {
            block = DeserializeBlock(data)
        }

        return nil
    })
    handleError(err)

    if block != nil {
        it.curHash = block.PrevHash
    }

    return block
}

func (bc *BlockChain) PrintChain() {
    it := bc.Iterator()

    for {
        block := it.Next()
        if block == nil {
            break
        }

        fmt.Printf("%s\n%x\n%x\n\n", block.Data, block.Hash, block.PrevHash)
    }
}
