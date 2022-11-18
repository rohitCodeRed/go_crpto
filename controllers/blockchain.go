package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/model"
)

type hashes struct {
	Index int    `json:"index"`
	Hash  string `json:"hash"`
}

type transaction struct {
	Sender   string  `json:"sender"`
	Reciever string  `json:"reciever"`
	Amount   float64 `json:"amount"`
}

func MineBlock(c *gin.Context, b *blockchain.BlockChain) {
	if len(b.TRANSACTIONS) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Transaction are not available.."})
		return
	}
	previous_block := b.Get_previous_block()
	previous_proof := previous_block.Proof
	proof := blockchain.Proof_of_work(previous_proof)
	previous_hash := blockchain.Hash(previous_block)
	sender := b.GetUuidAddress()
	reciever := b.UserName
	amount := 1.0
	b.Add_transaction(sender, reciever, amount)

	block := b.Create_block(proof, previous_hash)

	b.UpdateAmount()
	UpdateDataForUser(b.UserName, b)

	//ping to other node to update block..
	b.Ping_other_nodes_to_replaceChain()

	c.JSON(http.StatusOK, gin.H{"message": "Block Mined succesfully",
		"index":         block.Index,
		"timestamp":     block.DateTime,
		"proof":         block.Proof,
		"previous_hash": block.Previous_hash,
		"transactions":  block.Transactions})

}

func GetChain(c *gin.Context, b *blockchain.BlockChain) {
	block := b.CHAIN
	hashVar := make([]hashes, 0)

	for _, bl := range block {
		hash := hashes{Index: bl.Index, Hash: blockchain.Hash(bl)}
		hashVar = append(hashVar, hash)
	}

	c.JSON(http.StatusOK, gin.H{
		"chain":  block,
		"hashes": hashVar,
		"length": len(block)})
}

func IsChainValid(c *gin.Context, b *blockchain.BlockChain) {
	is_valid := b.Is_chain_valid(b.CHAIN)
	if is_valid {
		c.JSON(http.StatusOK, gin.H{
			"message": "All good. The Blockchain is valid."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"chain": "May Day, we have a problem. The Blockchain is not valid."})
}

func AddTransaction(c *gin.Context, b *blockchain.BlockChain) {
	var data model.Transaction
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := b.Add_transaction(data.Sender, data.Reciever, data.Amount)
	UpdateDataForUser(b.UserName, b)

	b.Ping_nodes_to_add_transaction(data.Sender, data.Reciever, data.Amount)
	c.JSON(http.StatusOK, gin.H{
		"message": "This transaction will be added with index: " + strconv.Itoa(index)})

}

func ConnectNode(c *gin.Context, b *blockchain.BlockChain) {
	var nodes []model.Node

	if err := c.ShouldBindJSON(&nodes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, node := range nodes {
		b.Add_node(&node)
	}

	UpdateDataForUser(b.UserName, b)
	c.JSON(http.StatusOK, gin.H{
		"message": "All nodes are connected", "nodes_length": len(nodes)})

}

func Replace_chain(c *gin.Context, b *blockchain.BlockChain) {
	is_chain_replaced := b.Replace_node_chain()
	if !is_chain_replaced {
		c.JSON(http.StatusOK, gin.H{"message": "No need to replace chain."})
		return
	}
	b.UpdateAmount()
	UpdateDataForUser(b.UserName, b)
	c.JSON(http.StatusOK, gin.H{"message": "Chain Replaced succesfully."})

}
