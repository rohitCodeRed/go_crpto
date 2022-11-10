package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/config"
	"github.com/rohitCodeRed/go_crypto/controllers"
	"github.com/rohitCodeRed/go_crypto/render"
	"github.com/rohitCodeRed/go_crypto/routes"
)

func main() {
	var b blockchain.BlockChain
	var app config.AppConfig
	app.InProduction = false

	tc, errT := render.CreateTemplate()
	if errT != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplate(&app)

	repo := controllers.NewRepo(&app)
	controllers.NewHandler(repo)

	b.New()
	fmt.Println("Server Unique Address: ", b.GetUuidAddress())
	//fmt.Println(node_address)

	server := &http.Server{
		Addr:         ":4000",
		Handler:      routes.Router(&b, &app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
