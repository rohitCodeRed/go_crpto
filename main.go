package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/routes"
)

func main() {
	node_address := blockchain.New()
	fmt.Println("Server Unique Address: ", node_address)

	server := &http.Server{
		Addr:         ":4000",
		Handler:      routes.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
