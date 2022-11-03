package model

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	Sender   string `json:"sender"`
	Reciever string `json:"reciever"`
	Amount   int    `json:"amount"`
}

type Block struct {
	Index         int           `json:"index"`
	DateTime      string        `json:"dateTime"`
	Proof         int           `json:"proof"`
	Previous_hash string        `json:"previous_hash"`
	Transactions  []Transaction `json:"transactions"`
}

type ChainModel struct {
	Chain  []Block `json:"chain"`
	Length int     `json:"length"`
}

func (c ChainModel) TextOutput() string {
	b, err := json.Marshal(c.Chain[0].Transactions)
	if err != nil {
		fmt.Println(err)
	}

	p := fmt.Sprintf(
		"Index: %d\nDateTime : %s\nProof: %d\nPrevious_hash: %s\nTransactions:%s\n ChainLength:%d\n",
		c.Chain[0].Index, c.Chain[0].DateTime, c.Chain[0].Proof, c.Chain[0].Previous_hash, string(b), c.Length)
	return p
}
