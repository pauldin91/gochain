package domain

import (
	"encoding/json"

	"github.com/pauldin91/gochain/src/internal"
)

type BlockchainDto struct {
	Chain []string `json:"chain"`
}

type Blockchain struct {
	Chain []Block
}

func (bc Blockchain) String() string {
	var bcs []string
	for _, b := range bc.Chain {
		js, _ := json.Marshal(b)
		bcs = append(bcs, string(js))
	}
	res, _ := json.Marshal(BlockchainDto{Chain: bcs})
	return string(res)

}

func Create() Blockchain {
	bc := Blockchain{}
	bc.Chain = append(bc.Chain, genesis())
	return bc
}

func (bc *Blockchain) AddBlock(data string) Block {
	block := mineBlock(bc.Chain[len(bc.Chain)-1], data)
	bc.Chain = append(bc.Chain, block)
	return block
}
func IsValid(bc []Block) bool {

	jsonGenesis, _ := json.Marshal(bc[0])
	gen, _ := json.Marshal(genesis())
	if string(jsonGenesis) != string(gen) {
		return false
	}
	for i := 1; i < len(bc); i++ {
		block := bc[i]
		lastBlock := bc[i-1]
		block.Hash = ""
		expectedHash := internal.Hash(block.ToString())
		block.Hash = bc[i].Hash
		if block.LastHash != lastBlock.Hash ||
			block.Hash != expectedHash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) ReplaceChain(newChain []Block) bool {

	isValid := IsValid(newChain)

	if len(newChain) <= len(bc.Chain) || !isValid {
		return false
	}
	bc.Chain = newChain
	return true
}
