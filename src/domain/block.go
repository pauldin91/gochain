package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/pauldin91/gochain/src/utils"
)

var genesisLastHash = strings.Repeat("*", 32)

var cfg, _ = utils.LoadConfig(".")

type Block struct {
	Timestamp  time.Time `json:"timestamp"`
	LastHash   string    `json:"last_hash"`
	Hash       string    `json:"hash"`
	Data       string    `json:"data"`
	Nonce      int64     `json:"nonce"`
	Difficulty int64     `json:"difficulty"`
}

func genesis() Block {
	block := Block{
		LastHash: genesisLastHash,
		Nonce:    0,
	}
	block.Data = ""
	block.Hash = utils.Hash(block.ToString())
	return block
}

func (b *Block) ToString() string {
	return fmt.Sprintf("Timestamp: %s\nLastHash: %s\nHash: %s\nData: %s\nNonce: %d\nDifficulty: %d\n",
		b.Timestamp.Format(time.RFC3339), b.LastHash, b.Hash, b.Data, b.Nonce, b.Difficulty)
}

func adjustDifficulty(lastBlock Block, currentTime time.Time) int64 {
	diff := lastBlock.Difficulty
	var start time.Time
	if lastBlock.Timestamp.IsZero() {
		start = time.Now().UTC()
	} else {
		start = lastBlock.Timestamp
	}
	dur := start.UnixMilli() + int64(cfg.MineRate)

	if dur > currentTime.UnixMilli() {
		diff += 1
	} else {
		diff -= 1
		if diff <= 0 {
			diff = 1
		}
	}
	return diff
}

func mineBlock(lastBlock Block, data string) Block {

	var hash string
	var timestamp time.Time
	var nonce int64 = 0

	for {
		nonce++
		timestamp = time.Now().UTC()
		difficulty := adjustDifficulty(lastBlock, timestamp)
		pref := strings.Repeat("0", int(difficulty))
		copy := Block{
			Nonce:      nonce,
			Timestamp:  timestamp,
			Difficulty: difficulty,
			LastHash:   lastBlock.Hash,
			Data:       data,
		}
		hash = utils.Hash(copy.ToString())
		copy.Hash = hash
		if strings.HasPrefix(copy.Hash, pref) {
			return copy
		}
	}
}
