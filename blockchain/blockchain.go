package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rohitCodeRed/go_crypto/model"
)

type SocketInfo struct {
	Conn   net.Conn
	Opcode byte
	User   string
}
type BlockChain struct {
	CHAIN             []model.Block       `json:"chain"`
	TRANSACTIONS      []model.Transaction `json:"transactions"`
	NODES             []model.Node        `json:"nodes"`
	Uuid              string              `json:"uuid"`
	TOTAL_AMOUNT      float64             `json:"amount"`
	UserName          string              `json:"name"`
	Url               string              `json:"url"`
	IsBlockChanged    bool
	SocketConnections []SocketInfo
}

// var Uuuid string

// var CHAIN []model.Block
// var TRANSACTIONS []model.Transaction
// var NODES []model.Node

func (b *BlockChain) New() string {
	b.CHAIN = []model.Block{}
	b.TRANSACTIONS = []model.Transaction{}
	b.Create_block(1, "0")
	b.Uuid = uuid.New().String()
	b.IsBlockChanged = true
	b.NODES = []model.Node{}
	b.TOTAL_AMOUNT = 0.0
	return b.Uuid
}

func (b *BlockChain) Create_block(proof int, previous_hash string) model.Block {

	newBlock := model.Block{Index: len(b.CHAIN),
		Proof:         proof,
		DateTime:      time.Now().String(),
		Previous_hash: previous_hash,
		Transactions:  b.TRANSACTIONS}

	b.TRANSACTIONS = []model.Transaction{}

	b.CHAIN = append(b.CHAIN, newBlock)

	b.IsBlockChanged = true //data chnaged...
	//controllers.UpdateDataForUser(b.UserName,&b)
	return newBlock

}

func (b *BlockChain) Get_previous_block() model.Block {
	return b.CHAIN[len(b.CHAIN)-1]
}

func (b *BlockChain) Is_chain_valid(pChain []model.Block) bool {
	previous_block := b.CHAIN[0]
	block_index := 1
	for block_index > len(b.CHAIN) {
		block := b.CHAIN[block_index]
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

func (b *BlockChain) Add_transaction(sender string, reciever string, amount float64) int {
	pTransaction := model.Transaction{Sender: sender, Reciever: reciever, Amount: amount}
	b.TRANSACTIONS = append(b.TRANSACTIONS, pTransaction)
	prevBlock := b.Get_previous_block()

	b.IsBlockChanged = true
	return prevBlock.Index
}

func (b *BlockChain) Add_node(node *model.Node) {
	pNode := model.Node{Address: node.Address, Name: node.Name}
	b.NODES = append(b.NODES, pNode)
	b.IsBlockChanged = true
}

func (b *BlockChain) Replace_node_chain() bool {
	//networks := NODES
	longest_chain := []model.Block{}
	max_length := len(b.CHAIN)

	for _, node := range b.NODES {
		url := node.Address
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
		if length > max_length && b.Is_chain_valid(pChain) {
			max_length = length
			longest_chain = pChain
		}

	}

	if len(longest_chain) > 0 {
		b.CHAIN = longest_chain
		b.IsBlockChanged = true
		b.TRANSACTIONS = []model.Transaction{}
		return true
	}

	return false
}

func (b *BlockChain) Ping_other_nodes_to_replaceChain() {
	for _, node := range b.NODES {
		url := node.Address
		resp, err := http.Get("http://" + url + "/replace_chain")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
	}
}

func (b *BlockChain) Ping_nodes_to_add_transaction(sender string, reciever string, amount float64) {
	pTransaction := model.Transaction{Sender: sender, Reciever: reciever, Amount: amount}
	postBody, _ := json.Marshal(pTransaction)
	responseBody := bytes.NewBuffer(postBody)

	for _, node := range b.NODES {
		url := node.Address
		resp, err := http.Post("http://"+url+"/update_transaction", "application/json", responseBody)
		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	}
}

func (b *BlockChain) GetUuidAddress() string {
	return b.Uuid
}

func (b *BlockChain) UpdateAmount() {
	chain := b.CHAIN
	for _, chainVal := range chain {
		transactions := chainVal.Transactions
		for _, transVal := range transactions {
			if transVal.Reciever == b.UserName {
				b.TOTAL_AMOUNT += transVal.Amount
			}

			if transVal.Sender == b.UserName {
				b.TOTAL_AMOUNT -= transVal.Amount
			}
		}
	}
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
	j, err := json.Marshal(pBlock)

	if err != nil {
		fmt.Println(err)
	}
	hash256 := sha256.New()
	hash256.Write(j)
	// blockHash := sha256.Sum256([]byte(b))
	newHash := hex.EncodeToString(hash256.Sum(nil))

	return newHash

}
