package domain

import (
	"strings"
	"testing"
	"time"

	"github.com/pauldin91/gochain/src/utils"
)

var genesisBlock = genesis()

func TestGenesis(t *testing.T) {

	if len(genesisBlock.Data) != 0 ||
		genesisBlock.LastHash != genesisLastHash {
		t.Error("Data and lasthash should be empty for genesis")
	}
	if genesisBlock.Difficulty != 0 ||
		genesisBlock.Nonce != 0 {
		t.Error("Difficulty and nonce should be 0 for genesis")
	}
	block := Block{
		LastHash: genesisLastHash,
		Nonce:    0,
	}
	block.Data = ""
	block.Hash = utils.Hash(block.ToString())
	if genesisBlock.Hash != block.Hash {
		t.Error("Hashes missmatch")
	}

	if !genesisBlock.Timestamp.IsZero() {
		t.Error("Genesis time is not zero")
	}
}

func TestAdjustDifficulty(t *testing.T) {
	diff := adjustDifficulty(genesisBlock, time.Now().UTC())
	if diff != 1 {
		t.Errorf("Difficulty should be %d\n", diff)
	}
	genesisBlock.Difficulty = 5
	diff = adjustDifficulty(genesisBlock, time.Now().UTC().Add(time.Duration(time.Second*4)))
	if diff != 4 {
		t.Errorf("Difficulty should be %d\n", diff)
	}
	genesisBlock.Difficulty = 0
}

func TestMineBlock(t *testing.T) {
	mined := mineBlock(genesisBlock, "")
	if !strings.HasPrefix(mined.Hash, strings.Repeat("0", int(genesisBlock.Difficulty))) {
		t.Errorf("Difficulty was %d while output was %s", genesisBlock.Difficulty, mined.Hash)
	}
	genesisBlock.Difficulty = 5
	time.Sleep(time.Second * 4)
	mined = mineBlock(genesisBlock, "")
	if !strings.HasPrefix(mined.Hash, strings.Repeat("0", int(genesisBlock.Difficulty-1))) {
		t.Errorf("Difficulty was %d while output was %s", genesisBlock.Difficulty, mined.Hash)
	}
	genesisBlock.Difficulty = 0

}
