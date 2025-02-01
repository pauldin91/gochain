package domain

import (
	"fmt"

	"github.com/pauldin91/gochain/src/internal"
)

type Wallet struct {
	balance   float64
	keyPair   internal.KeyPair
	publicKey string
}

func (w Wallet) ToString() string {
	return fmt.Sprintf("Wallet - \npublicKey\t: %s\nbalance: %.8f\n", w.publicKey, w.balance)
}
