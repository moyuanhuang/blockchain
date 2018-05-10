package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

const VERSION = byte(0x00)
const ADDRESSCHECKSUMLEN = 4

// Wallet is just nothing but a key pair
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	return &Wallet{private, public}
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	// rand.Reader is a shared instance o f a cryptographically strong
	// pseudo-random generator.
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	handleError(err)
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, public
}

func (w *Wallet) GetAddress() []byte {
	hashedPubKey := hashPubKey(w.PublicKey)
	versionedPayload := append([]byte{VERSION}, hashedPubKey...)

	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)

	return Base58Encode(fullPayload)
}

func ValidateAddress(address string) bool {
	fullPayload := Base58Decode([]byte(address))
	version := fullPayload[0]
	pubKeyHash := fullPayload[1 : len(fullPayload)-ADDRESSCHECKSUMLEN]
	expected := fullPayload[len(fullPayload)-ADDRESSCHECKSUMLEN:]

	actual := checksum(append([]byte{version}, pubKeyHash...))
	return bytes.Compare(actual, expected) == 0
}

func hashPubKey(key []byte) []byte {
	publicSHA256 := sha256.Sum256(key)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	handleError(err)

	return RIPEMD160Hasher.Sum(nil)
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:ADDRESSCHECKSUMLEN]
}
