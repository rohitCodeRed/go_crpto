package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rohitCodeRed/go_crypto/model"
)

// type transaction struct {
// 	sender   string
// 	reciever string
// 	amount   int
// }

// type block struct {
// 	index         int
// 	dateTime      string
// 	proof         int
// 	previous_hash string
// 	TRANSACTIONS  []transaction
// }

var aDDRESS string

type node struct {
	url string
}

var CHAIN []model.Block
var TRANSACTIONS []model.Transaction
var nodes []node

func New() string {
	CHAIN = []model.Block{}
	TRANSACTIONS = []model.Transaction{}
	Create_block(1, "0")
	aDDRESS := uuid.New().String()
	return aDDRESS
}

func Create_block(proof int, previous_hash string) {

	newBlock := model.Block{Index: len(CHAIN),
		Proof:         proof,
		DateTime:      time.Now().String(),
		Previous_hash: previous_hash,
		Transactions:  TRANSACTIONS}

	TRANSACTIONS = []model.Transaction{}

	CHAIN = append(CHAIN, newBlock)

}

func Get_previous_block() model.Block {
	return CHAIN[len(CHAIN)-1]
}

func Proof_of_work(previous_proof int) int {
	new_proof := 1
	checkProof := false

	for checkProof == false {
		hash256 := sha256.New()
		input_proof := ((new_proof * new_proof) - (previous_proof * previous_proof))

		hash256.Write([]byte(strconv.Itoa(input_proof)))
		newHash := hex.EncodeToString(hash256.Sum(nil))

		if newHash[:4] == "0000" {
			checkProof = true
			break
		}
		new_proof++

	}
	return new_proof
}

func Hash(pBlock model.Block) string {
	b, err := json.Marshal(pBlock)

	if err != nil {
		fmt.Println(err)
	}
	hash256 := sha256.New()
	hash256.Write(b)
	// blockHash := sha256.Sum256([]byte(b))
	newHash := hex.EncodeToString(hash256.Sum(nil))

	return newHash

}

func is_chain_valid(pChain []model.Block) bool {
	previous_block := CHAIN[0]
	block_index := 1
	for block_index > len(CHAIN) {
		block := CHAIN[block_index]
		if block.Previous_hash != Hash(previous_block) {
			return false
		}
		previous_proof := previous_block.Proof
		proof := block.Proof

		hash256 := sha256.New()
		input_proof := ((proof * proof) - (previous_proof * previous_proof))

		hash256.Write([]byte(strconv.Itoa(input_proof)))
		newHash := hex.EncodeToString(hash256.Sum(nil))

		if newHash[:4] == "0000" {
			return false
		}

		previous_block = block
		block_index++
	}
	return true
}

func Add_transaction(sender string, reciever string, amount int) int {
	pTransaction := model.Transaction{Sender: sender, Reciever: reciever, Amount: amount}
	TRANSACTIONS = append(TRANSACTIONS, pTransaction)
	prevBlock := Get_previous_block()

	return prevBlock.Index
}

func add_node(address string) {
	pNode := node{url: address}
	nodes = append(nodes, pNode)
}

func replace_node_chain() (bool, error) {
	//networks := nodes
	longest_chain := []model.Block{}
	max_length := len(CHAIN)

	for _, node := range nodes {
		url := node.url
		resp, err := http.Get("http://" + url + "/get_chain")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		var cResp model.ChainModel
		//Decode the data
		if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
			log.Fatal("ooopsss! an error occurred, please try again")
		}
		length := cResp.Length
		pChain := cResp.Chain

		//create
		if length > max_length && is_chain_valid(pChain) {
			max_length = length
			longest_chain = pChain
		}

	}

	if len(longest_chain) > 0 {
		CHAIN = longest_chain
		return true, nil
	}

	return false, nil
}

func GetUuidAddress() string {
	return aDDRESS
}
