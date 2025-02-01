package internal

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"

	"github.com/google/uuid"
)

func id() string {
	return uuid.NewString()
}

type KeyPair struct {
	Private string
	Public  string
}

func Hash(data string) string {
	result := sha256.Sum256([]byte(data))
	return hex.EncodeToString(result[:])
}

func genKeyPair() (private string, public string, err error) {
	pair, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	byts, _ := x509.MarshalECPrivateKey(pair)
	pub, err := x509.MarshalPKIXPublicKey(pair.PublicKey)
	private = string(byts)
	public = string(pub)
	return private, public, err
}
