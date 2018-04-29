package main

import (
    "fmt"
    "os"
    "encoding/hex"

    "github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const BlocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

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
    if !dbExists() {
        fmt.Println("No existing blockchain found. Create one first.")
        os.Exit(1)
    }

    db, err := bolt.Open(dbFile, 0600, nil)
    handleError(err)

    var lastHash []byte

    err = db.Update(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BlocksBucket))
        lastHash = bucket.Get([]byte("l"))

        return nil
    })
    handleError(err)

    bc := BlockChain{lastHash, db}

    return &bc
}

func CreateBlockchain(address string) *BlockChain {
    if dbExists() {
        fmt.Println("Blockchain already exists.")
        os.Exit(1)
    }

    var lastHash []byte

    db, err := bolt.Open(dbFile, 0600, nil)
    handleError(err)

    defer db.Close()

    err = db.Update(func(tx *bolt.Tx) error {
        cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
        genesis := NewGenesisBlock(cbtx)

        b, err := tx.CreateBucket([]byte(BlocksBucket))
        handleError(err)

        err = b.Put(genesis.Hash, genesis.Serialize())
        handleError(err)

        err = b.Put([]byte("l"), genesis.Hash)
        handleError(err)

        lastHash = genesis.Hash
        return nil
    })

    handleError(err)
    return &BlockChain{lastHash, db}
}

// create a new block that contains txs
func (bc *BlockChain) MineBlock(txs []*Transaction) {
    var lastHash []byte

    err := bc.db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BlocksBucket))
        lastHash = bucket.Get(lastHashKey)

        return nil
    })
    handleError(err)

    newBlock := NewBlock(txs, lastHash)

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

        fmt.Printf("Hash: %x\n", block.Hash)
        fmt.Printf("Prev Hash: %x\n", block.PrevHash)
        fmt.Printf("Number of transactions: %d\n\n", len(block.Transactions))

    }
}

func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction{
    var unspentTXs []Transaction
    spentTx := make(map[string][]int)
    iter := bc.Iterator()

    for {
        // notice that the iterator's order is from back to front
        block := iter.Next()
        for _, tx := range block.Transactions {
            txID := hex.EncodeToString(tx.ID)
        outputs:
            for outIdx, out := range tx.Outputs {
                if !out.CanBeUnlockedWith(address) {
                    continue
                }
                if spentTx[txID] != nil {
                    for _, spentIdx := range spentTx[txID] {
                        if outIdx == spentIdx {
                            continue outputs
                        }
                    }
                }
                unspentTXs = append(unspentTXs, *tx)
            }

            if !tx.IsCoinbaseTx() {
                for _, input := range tx.Inputs {
                    if input.CanUnlockOutputWith(address) {
                        spentTxId := hex.EncodeToString(input.Txid)
                        spentTx[spentTxId] = append(spentTx[spentTxId], input.Vout)
                    }
                }
            }
        }
        if block.IsGenesisBlock() {
            break
        }
    }
    return unspentTXs
}

// find outputs that is larger than amount. The return values are
// total spendable amount found, and a map that indexes the TxID
// and Vout of a spendable UTXO
func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
    accumulated := 0
    spendableOutputMap := make(map[string][]int)
    unspentTxs := bc.FindUnspentTransactions(address)

    find:
    for _, tx := range unspentTxs {
        txID := hex.EncodeToString(tx.ID)
        for outIdx, output := range tx.Outputs {
            if output.CanBeUnlockedWith(address) {
                spendableOutputMap[txID] = append(spendableOutputMap[txID], outIdx)
                accumulated += output.Value

                if accumulated >= amount {
                    break find
                }
            }
        }
    }
    return accumulated, spendableOutputMap
}

func (bc *BlockChain) FindUTXO(address string) []TXOutput{
    unspentTXs := bc.FindUnspentTransactions(address)
    UTXOs := make([]TXOutput, 0)

    for _, tx := range unspentTXs {
        for _, out := range tx.Outputs {
            if out.CanBeUnlockedWith(address) {
                UTXOs = append(UTXOs, out)
            }
        }
    }

    return UTXOs
}

func dbExists() bool {
    if _, err := os.Stat(dbFile); os.IsNotExist(err) {
        return false
    }
    return true
}
