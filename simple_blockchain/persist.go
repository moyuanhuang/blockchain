package main

import (
    "encoding/gob"
    "bytes"
)

func DeserializeBlock(data []byte) (block *Block) {
    dec := gob.NewDecoder(bytes.NewReader(data))

    err := dec.Decode(block)
    handleError(err)
    return
}
