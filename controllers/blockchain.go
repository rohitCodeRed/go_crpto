package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/model"
)

func MineBlock(c *gin.Context) {
	previous_block := blockchain.Get_previous_block()
	previous_proof := previous_block.Proof
	proof := blockchain.Proof_of_work(previous_proof)
	previous_hash := blockchain.Hash(previous_block)
	sender := blockchain.GetUuidAddress()
	reciever := "you"
	amount := 1
	blockchain.Add_transaction(sender, reciever, amount)

	block := blockchain.Create_block(proof, previous_hash)

	c.JSON(http.StatusOK, gin.H{"message": "Block Mined succesfully",
		"index":         block.Index,
		"timestamp":     block.DateTime,
		"proof":         block.Proof,
		"previous_hash": block.Previous_hash,
		"transactions":  block.Transactions})

}

func GetChain(c *gin.Context) {
	block := blockchain.CHAIN
	c.JSON(http.StatusOK, gin.H{
		"chain":  block,
		"length": len(block)})
}

func IsChainValid(c *gin.Context) {
	is_valid := blockchain.Is_chain_valid(blockchain.CHAIN)
	if is_valid {
		c.JSON(http.StatusOK, gin.H{
			"message": "All good. The Blockchain is valid."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"chain": "May Day, we have a problem. The Blockchain is not valid."})
}

type transaction struct {
	Sender   string `json:"sender"`
	Reciever string `json:"reciever"`
	Amount   int    `json:"amount"`
}

func AddTransaction(c *gin.Context) {
	var data model.Transaction
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := blockchain.Add_transaction(data.Sender, data.Reciever, data.Amount)
	c.JSON(http.StatusOK, gin.H{
		"message": "This transaction will be added to Block with index: " + strconv.Itoa(index)})

}

func ConnectNode(c *gin.Context) {
	var nodes []model.Node

	if err := c.ShouldBindJSON(&nodes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, node := range nodes {
		blockchain.Add_node(&node)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All nodes are connected", "nodes_length": len(nodes)})

}

func Replace_chain(c *gin.Context) {
	is_chain_replaced := blockchain.Replace_node_chain()
	if !is_chain_replaced {
		c.JSON(http.StatusOK, gin.H{"message": "No need to replace chain."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Chain Replaced succesfully."})

}
