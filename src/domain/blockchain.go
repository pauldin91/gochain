package domain

import (
	"encoding/json"

	"github.com/pauldin91/gochain/src/internal"
)

type Blockchain struct {
	Chain []Block
}

func Create() Blockchain {
	bc := Blockchain{}
	bc.Chain = append(bc.Chain, genesis())
	return bc
}

func (bc *Blockchain) AddBlock(data string, mineRate int64) Block {
	block := mineBlock(bc.Chain[len(bc.Chain)-1], data, mineRate)
	bc.Chain = append(bc.Chain, block)
	return block
}
func (b *Blockchain) IsValid(bc []Block) bool {

	jsonGenesis, _ := json.Marshal([]byte(bc[0].ToString()))
	gen := genesis()
	if string(jsonGenesis) != gen.ToString() {
		return false
	}
	for i := 1; i < len(bc); i++ {
		block := bc[i]
		lastBlock := bc[i-1]

		if block.LastHash != lastBlock.Hash ||
			block.Hash != internal.Hash(block.ToString()) {
			return false
		}
	}
	return true
}

func (bc *Blockchain) ReplaceChain(newChain []Block) bool {
	if len(newChain) <= len(newChain) || bc.IsValid(newChain) {
		return false
	}
	bc.Chain = newChain
	return true
}
