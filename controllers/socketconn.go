package controllers

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/model"
)

func GetRealTimeData(c *gin.Context, b *blockchain.BlockChain) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		log.Printf("upgrade error: %s", err)
		return
	}
	// store connections
	if len(b.SocketConnections) > 0 {
		for index, pConn := range b.SocketConnections {
			if pConn.User == b.UserName {
				b.SocketConnections[index].Conn = conn
			}
		}
	} else {
		b.SocketConnections = append(b.SocketConnections, blockchain.SocketInfo{Conn: conn, Opcode: 0x1, User: b.UserName})
	}

	log.Println("Client connected info: ", c.Request.Header)

	defer conn.Close()

	for {
		_, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Printf("read message error: %v", err)
			b.IsBlockChanged = true
			return
		}

		log.Println("op-----------:", op)

		if b.IsBlockChanged {
			jsonBlock, err := json.Marshal(GetBlockInfoForSocket(b))
			if err != nil {
				log.Println("error:", err)
			}

			err = wsutil.WriteServerMessage(conn, 0x1, jsonBlock)
			if err != nil {
				log.Printf("write message error: %v", err)
				return
			}
			b.IsBlockChanged = false
		}

	}
}

type BlockInfo struct {
	CHAIN        []model.Block       `json:"chain"`
	TRANSACTIONS []model.Transaction `json:"transactions"`
	NODES        []model.Node        `json:"nodes"`
	Uuid         string              `json:"uuid"`
	TOTAL_AMOUNT float64             `json:"amount"`
	UserName     string              `json:"name"`
	Url          string              `json:"url"`
	Hash         []hashes            `json:"hash"`
}

func GetBlockInfoForSocket(b *blockchain.BlockChain) *BlockInfo {
	block := b.CHAIN
	hashVar := make([]hashes, 0)

	for _, bl := range block {
		hash := hashes{Index: bl.Index, Hash: blockchain.Hash(bl)}
		hashVar = append(hashVar, hash)
	}

	blockInfo := BlockInfo{CHAIN: b.CHAIN, TRANSACTIONS: b.TRANSACTIONS, NODES: b.NODES, Uuid: b.Uuid, TOTAL_AMOUNT: b.TOTAL_AMOUNT, Hash: hashVar, UserName: b.UserName, Url: b.Url}
	return &blockInfo
}

func UpdateDataForUser(username string, b *blockchain.BlockChain) {
	for _, val := range b.SocketConnections {
		if val.User == username {
			cConn := val.Conn
			jsonBlock, err := json.Marshal(GetBlockInfoForSocket(b))
			if err != nil {
				log.Println("error:", err)
			}

			err = wsutil.WriteServerMessage(cConn, ws.OpCode(val.Opcode), jsonBlock)
			if err != nil {
				log.Printf("write message error: %v", err)
				return
			}
			b.IsBlockChanged = false
		}
	}
}
