package domain

import (
	"fmt"

	"github.com/pauldin91/gochain/src/internal"
)

type Wallet struct {
	balance float64
	keyPair internal.KeyPair
	address string
}

func (w Wallet) ToString() string {
	return fmt.Sprintf("Wallet - \npublicKey\t: %s\nbalance: %.8f\n", w.keyPair.GetPublicKey(), w.balance)
}

func (w Wallet) CalculateBalance(chain Blockchain) {

}

func blockchainWallet() Wallet {
	return Wallet{
		balance: 0.0,
		keyPair: internal.NewKeyPair(),
		address: "hakuna-mattata",
	}
}
