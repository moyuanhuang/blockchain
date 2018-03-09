package main

import (
    "fmt"
    "math"
    "math/big"
    "bytes"
    "crypto/sha256"
    "log"
)

// how many bit we require to be zero
const (
    TargetBits = 24
    MaxNonce = math.MaxInt64
)

type ProofofWork struct {
    block *Block
    target *big.Int
}

func NewProofOfWork(block *Block) *ProofofWork {
    target := big.NewInt(1)
    target.Lsh(target, 256 - TargetBits)
    return &ProofofWork{block, target}
}

func (pow *ProofofWork) prepareData(nonce int) []byte {
    data := bytes.Join([][]byte{
        pow.block.PrevHash,
        pow.block.Data,
        IntToHex(pow.block.Timestamp),
        IntToHex(int64(TargetBits)),
        IntToHex(int64(nonce)),
    }, []byte{})

    return data
}
func (pow *ProofofWork) Run() (nonce uint, hash [32]byte) {
    fmt.Printf("Mining a new block...\n")
    var hashInt big.Int

    for nonce := 0; nonce < MaxNonce; nonce++ {
        data := pow.prepareData(nonce)
        hash = sha256.Sum256(data)
        hashInt.SetBytes(hash[:])

        if hashInt.Cmp(pow.target) == -1 {
            fmt.Printf("%x\n", hash)
            break
        }
    }

    if nonce == MaxNonce {
        log.Panic("Failed to provide POW.\n")
    }
    return
}
