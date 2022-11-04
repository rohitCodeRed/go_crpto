package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohitCodeRed/go_crypto/blockchain"
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

	blockchain.Create_block(proof, previous_hash)

	c.JSON(http.StatusOK, gin.H{"message": "Block Mined succesfully"})
}

func GetChain(c *gin.Context) {

}

func IsChainValid(c *gin.Context) {

}

func AddTransaction(c *gin.Context) {

}

func ConnectNode(c *gin.Context) {

}

func Replace_chain(c *gin.Context) {

}
