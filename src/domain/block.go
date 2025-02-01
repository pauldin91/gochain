package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/pauldin91/gochain/src/internal"
)

type Block struct {
	Timestamp  time.Time `json:"timestamp"`
	LastHash   string    `json:"last_hash"`
	Hash       string    `json:"hash"`
	Data       []string  `json:"data"`
	Nonce      int64     `json:"nonce"`
	Difficulty int64     `json:"difficulty"`
}

func genesis() Block {
	block := Block{
		LastHash: "**************",
		Nonce:    0,
	}
	block.Data = make([]string, 0)
	block.Hash = internal.Hash(block.ToString())
	return block
}

func (b *Block) ToString() string {
	return fmt.Sprintf("Timestamp: %s", b.Timestamp.Format(time.RFC3339), b.LastHash, b.Hash, strings.Join(b.Data, ","), b.Nonce, b.Difficulty)
}

func adjustDifficulty(lastBlock Block, currentTime time.Time, mineRate int64) int64 {
	diff := lastBlock.Difficulty
	dur := lastBlock.Timestamp.UnixMilli() + mineRate

	if dur > currentTime.UnixMilli() {
		diff += 1
	} else {
		diff -= 1
	}
	return diff
}

func mineBlock(lastBlock Block, data string, mineRate int64) Block {
	copy := Block{
		Timestamp:  lastBlock.Timestamp,
		Hash:       lastBlock.Hash,
		LastHash:   lastBlock.LastHash,
		Data:       lastBlock.Data,
		Nonce:      0,
		Difficulty: lastBlock.Difficulty,
	}
	nonce := 0

	for {
		nonce++
		copy.Timestamp = time.Now().UTC()
		copy.Difficulty = adjustDifficulty(lastBlock, copy.Timestamp, mineRate)
		copy.Hash = internal.Hash(copy.ToString())
		pref := strings.Repeat("0", int(copy.Difficulty))
		if strings.HasPrefix(copy.Hash, pref) {
			return copy
		}

	}

}
