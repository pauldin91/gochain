package internal

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"log"
)

type KeyPair struct {
	keyPair *ecdsa.PrivateKey
}

func NewKeyPair() KeyPair {
	pair, err := genKeyPair()
	if err != nil {
		log.Fatal("unable to generate Key pair")
	}
	return KeyPair{
		keyPair: pair,
	}
}

func (pair *KeyPair) GetPublicKey() string {
	pub, _ := x509.MarshalPKIXPublicKey(pair.keyPair)
	public := string(pub)
	return public
}

func (pair *KeyPair) Sign(hashedData string) string {

	res, err := pair.keyPair.Sign(rand.Reader, []byte(hashedData), crypto.SHA256)
	if err != nil {
		log.Fatal("could sign data")
	}
	return string(res)
}

func (pair *KeyPair) Verify() bool {

	return true
}

func genKeyPair() (private *ecdsa.PrivateKey, err error) {
	pair, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return pair, err
}

func Hash(data string) string {
	result := sha256.Sum256([]byte(data))
	return hex.EncodeToString(result[:])
}
