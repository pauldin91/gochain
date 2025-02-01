package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/pauldin91/gochain/src/internal"
)

type Block struct {
	timestamp  time.Time
	lastHash   string
	hash       string
	data       []string
	nonce      int64
	difficulty int64
}

func genesis() Block {
	block := Block{
		timestamp: time.Now().UTC(),
		lastHash:  "**************",
		nonce:     0,
	}
	block.data = make([]string, 0)
	block.hash = internal.Hash(block.ToString())
	return block
}

func (b *Block) ToString() string {
	return fmt.Sprintf("", b.timestamp, b.lastHash, b.hash, strings.Join(b.data, ","), b.nonce, b.difficulty)
}

func adjustDifficulty(lastBlock Block, currentTime time.Time, mineRate int64) int64 {
	diff := lastBlock.difficulty
	dur := lastBlock.timestamp.UnixMilli() + mineRate

	if dur > currentTime.UnixMilli() {
		diff += 1
	} else {
		diff -= 1
	}
	return diff
}

func mineBlock(lastBlock Block, data string, mineRate int64) Block {
	copy := Block{
		timestamp:  lastBlock.timestamp,
		hash:       lastBlock.hash,
		lastHash:   lastBlock.lastHash,
		data:       lastBlock.data,
		nonce:      0,
		difficulty: lastBlock.difficulty,
	}
	nonce := 0

	for {
		nonce++
		copy.timestamp = time.Now().UTC()
		copy.difficulty = adjustDifficulty(lastBlock, copy.timestamp, mineRate)
		copy.hash = internal.Hash(copy.ToString())
		pref := strings.Repeat("0", int(copy.difficulty))
		if strings.HasPrefix(copy.hash, pref) {
			return copy
		}

	}

}
