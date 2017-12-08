package main

import (
	"encoding/json"
)

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	Proof        int
	PreviousHash string
}

func (b Block) hash() string {
	by, _ := json.Marshal(b)

	return ComputeHashSha256(by)
}
