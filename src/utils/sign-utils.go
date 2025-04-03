package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/big"
)

type KeyPair struct {
	keyPair *ecdsa.PrivateKey
	private string
	public  string
}

// NewKeyPair generates a new key pair safely
func NewKeyPair() *KeyPair {
	var newKeyPair = KeyPair{}
	pair := genKeyPair(&newKeyPair.private, &newKeyPair.public)
	newKeyPair.keyPair = pair
	return &newKeyPair
}

// GetPublicKey returns the base64-encoded public key
func (pair *KeyPair) GetPublicKey() string {
	return pair.public
}

func (pair *KeyPair) Sign(hashedData string) string {

	res, err := pair.keyPair.Sign(rand.Reader, []byte(hashedData), crypto.SHA256)
	if err != nil {
		log.Fatal("could sign data")
	}
	return string(res)
}

func VerifySignature(publicKey string, hashed []byte, sig []byte) bool {
	// Decode the base64-encoded public key
	keyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false
	}

	// Ensure keyBytes length is even and sufficient for P-256 (32 bytes each for X and Y)
	if len(keyBytes) != 64 {
		return false
	}

	// Split keyBytes into X and Y coordinates
	x := new(big.Int).SetBytes(keyBytes[:32])
	y := new(big.Int).SetBytes(keyBytes[32:])

	// Construct ECDSA public key
	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(), // Ensure curve is correct
		X:     x,
		Y:     y,
	}

	// Verify the signature
	valid := ecdsa.VerifyASN1(pubKey, hashed, sig)
	return valid
}

func genKeyPair(priv, pub *string) *ecdsa.PrivateKey {
	// Generate a new private key
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// Convert private key to base64
	privKeyBytes := privKey.D.Bytes()
	*priv = base64.StdEncoding.EncodeToString(privKeyBytes)

	// Convert public key (X || Y) to base64
	pubKeyBytes := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
	*pub = base64.StdEncoding.EncodeToString(pubKeyBytes)

	return privKey
}

func Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
