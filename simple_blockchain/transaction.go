package main

import (
    "fmt"
    "bytes"
    "log"
    "os"
    "encoding/gob"
    "encoding/hex"
    "crypto/sha256"
)

type Transaction struct {
    ID []byte
    Inputs []TXInput
    Outputs []TXOutput
}

type TXOutput struct {
    Value int
    ScriptPubKey string
}

type TXInput struct {
    // these two is used to index a UTXO
    Txid []byte
    Vout int
    ScriptSig string
}

const SUBSIDY = 10

func NewCoinbaseTX(to, data string) *Transaction {
    if data == "" {
        data = fmt.Sprintf("Reward to %s", to)
    }

    inputs := TXInput{[]byte{}, -1, data}
    outputs := TXOutput{SUBSIDY, to}
    tx := Transaction{nil, []TXInput{inputs}, []TXOutput{outputs}}
    tx.SetID()
    return &tx
}

func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
    var inputs []TXInput
    var outputs []TXOutput

    acc, validOutputs := bc.FindSpendableOutputs(from, amount)

    if acc < amount {
        log.Panic("Not enough funds!")
        os.Exit(1)
    }

    for txid, outIndexes := range validOutputs {
        txID, err := hex.DecodeString(txid)
        handleError(err)

        for _, outIndex := range outIndexes {
            input := TXInput{txID, outIndex, from}
            inputs = append(inputs, input)
        }
    }

    outputs = append(outputs, TXOutput{amount, to})
    if acc > amount {
        outputs = append(outputs, TXOutput{acc - amount, from})
    }

    tx := Transaction{nil, inputs, outputs}
    tx.SetID()

    return &tx
}

// ********************
// methods of Transaction
// ********************
func (tx *Transaction) IsCoinbaseTx() bool {
    return len(tx.Inputs) == 1 && len(tx.Inputs[0].Txid) == 0 && tx.Inputs[0].Vout == -1
}

func (tx *Transaction) SetID() {
    var encoded bytes.Buffer
    var hash [32]byte

    enc := gob.NewEncoder(&encoded)
    err := enc.Encode(tx)
    handleError(err)
    hash = sha256.Sum256(encoded.Bytes())
    tx.ID = hash[:]
}

// ********************
// methods of TXInput
// ********************
func (input *TXInput) CanUnlockOutputWith(data string) bool {
    return input.ScriptSig == data
}


// ********************
// methods of TXOnput
// ********************
func (output *TXOutput) CanBeUnlockedWith(data string) bool {
    return output.ScriptPubKey == data
}
