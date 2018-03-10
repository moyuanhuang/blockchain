package main

import (
    "fmt"
    "github.com/boltdb/bolt"
    "log"
)

const dbFile = "blockchain.db"
const BlocksBucket = "blocks"

var lastHashKey = []byte("l")

type BlockChain struct {
    lastHash []byte
    db *bolt.DB
}

func NewBlockChain() *BlockChain {
    db, err := bolt.Open(dbFile, 0600, nil)
    handleError(err)

    var lastHash []byte

    err = db.Update(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BlocksBucket))
        if bucket == nil {
            fmt.Printf("Creating Genesis block...\n")
            gb := NewGenesisBlock()

            bucket, err = tx.CreateBucket([]byte(BlocksBucket))
            handleError(err)

            err = bucket.Put(lastHashKey, gb.Hash)
            handleError(err)

            err = bucket.Put(gb.Hash, gb.Serialize())
            handleError(err)
        } else {
            log.Panic("Genesis Block already exist!")
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
