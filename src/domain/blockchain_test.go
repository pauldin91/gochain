package domain

import (
	"encoding/json"
	"testing"
	"time"
)

var gen Block = genesis()

var msg = Input{
	Timestamp: time.Now().UTC(),
	Address:   "r3ciP13nT",
	Amount:    50.44,
}

func TestCreate(t *testing.T) {
	e := Create()
	jsonGen, _ := json.Marshal(gen)
	jsonFirst, _ := json.Marshal(e.Chain[0])
	if string(jsonFirst) != string(jsonGen) {
		t.Error("First block in chain must be genesis")
	}
}

func TestAddBlock(t *testing.T) {
	e := Create()
	jsonMsg, _ := json.Marshal(msg)
	e.AddBlock(string(jsonMsg))

	if len(e.Chain) != 2 {
		t.Error("invalid chain length")
	}
}

func TestReplaceChain(t *testing.T) {
	e := Create()
	jsonMsg, _ := json.Marshal(msg)
	e.AddBlock(string(jsonMsg))

	b := Create()
	res := e.ReplaceChain(b.Chain)
	if res {
		t.Error("longest chain must not be replaced by smaller ones")
	}
	res = b.ReplaceChain(e.Chain)
	if !res {
		t.Error("smaller chain must be replaced by longer one")
	}

}
