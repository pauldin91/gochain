package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pauldin91/gochain/src/domain"
)

type BalanceResponse struct {
	Message string `json:"message"`
}

type BlockchainDto struct {
	Chain string `json:"chain"`
}
type BlockRequestDto struct {
	Data string `json:"data"`
}

type WalletDto struct {
	Address string `json:"address"`
}

type PoolDto struct {
	Dtos []TransactionResponseDto `json:"dtos"`
}

type TransactionRequestDto struct {
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

type TransactionResponseDto struct {
	Data string `json:"data"`
}

func (t *TransactionResponseDto) Map(tr *domain.Transaction) {
	t.Data = tr.String()
}

func (pool *PoolDto) Map(tp *domain.TransactionPool) {
	for _, t := range tp.Transactions {
		tr := &TransactionResponseDto{
			Data: t.String(),
		}
		pool.Dtos = append(pool.Dtos, *tr)
	}
}

type BlockResponseDto struct {
	Timestamp  time.Time `json:"timestamp"`
	LastHash   string    `json:"last_hash"`
	Hash       string    `json:"hash"`
	Data       string    `json:"data"`
	Nonce      int64     `json:"nonce"`
	Difficulty int64     `json:"difficulty"`
}

func (dto *BlockResponseDto) Map(block domain.Block) {
	dto.Data = block.Data
	dto.LastHash = block.LastHash
	dto.Hash = block.Hash
	dto.Timestamp = block.Timestamp
	dto.Nonce = block.Nonce
	dto.Difficulty = block.Difficulty
}

func (dto *BlockchainDto) Map(ch domain.Blockchain) {
	dto.Chain = ch.String()
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: errMsg,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
