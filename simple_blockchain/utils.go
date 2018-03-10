package main

import (
    "bytes"
    "encoding/binary"
    "log"
)

const Debug = 1

func IntToHex(i int64) []byte {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, i)
    if err != nil {
        log.Panic(err)
    }
    return buff.Bytes()
}

func DPrint(fmt string, values ...interface{}) {
    if Debug > 0 {
        log.Printf(fmt, values...)
    }
    return
}

func handleError(err error) {
    if err != nil {
        log.Panic(err)
    }
}
