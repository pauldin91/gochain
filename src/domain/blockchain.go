package domain

import (
	"encoding/json"

	"github.com/pauldin91/gochain/src/internal"
)

type Blockchain struct {
	Chain []Block
}

func Create() {
	bc := Blockchain{}
	bc.Chain = append(bc.Chain, genesis())
}

func (bc *Blockchain) addBlock(data string, mineRate int64) Block {
	block := mineBlock(bc.Chain[len(bc.Chain)-1], data, mineRate)
	bc.Chain = append(bc.Chain, block)
	return block
}
func (b *Blockchain) isValid(bc []Block) bool {

	jsonGenesis, _ := json.Marshal([]byte(bc[0].ToString()))
	gen := genesis()
	if string(jsonGenesis) != gen.ToString() {
		return false
	}
	for i := 1; i < len(bc); i++ {
		block := bc[i]
		lastBlock := bc[i-1]

		if block.lastHash != lastBlock.hash ||
			block.hash != internal.Hash(block.ToString()) {
			return false
		}
	}
	return true
}

func (bc *Blockchain) replaceChain(newChain []Block) bool {
	if len(newChain) <= len(newChain) || bc.isValid(newChain) {
		return false
	}
	bc.Chain = newChain
	return true
}
