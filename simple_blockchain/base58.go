package main

import (
	"bytes"
	"math/big"
)

var ALPHABET = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0)
	x.SetBytes(input)

	zero := big.NewInt(0)
	base := big.NewInt(int64(len(ALPHABET)))
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, ALPHABET[mod.Int64()])
	}

	result = ReverseBytes(result)
	for c := range input {
		if c == 0x00 {
			result = append([]byte{ALPHABET[0]}, result...)
		} else {
			break
		}
	}

	return result
}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for c := range input {
		if c == 0x00 {
			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _, c := range payload {
		charIndex := bytes.IndexByte(ALPHABET, c)
		result.Mul(result, big.NewInt(int64(len(ALPHABET))))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}
