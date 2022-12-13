package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/config"
	"github.com/rohitCodeRed/go_crypto/controllers"
	"github.com/rohitCodeRed/go_crypto/render"
	"github.com/rohitCodeRed/go_crypto/routes"
)

const USERNAME = "Rohit" //default user name
const URL = "localhost"
const PORT = "4000" //default Port number
const INITIAL_MONEY = 5.0

func main() {
	pUname := USERNAME
	pUrl := URL
	pPort := PORT

	if len(os.Args) > 2 {
		pUname = os.Args[1]
		pPort = os.Args[2]
	}
	pUrlPort := pUrl + ":" + pPort

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
	b.UserName = pUname
	b.Url = pUrlPort
	b.TOTAL_AMOUNT = INITIAL_MONEY

	fmt.Println("Sever Name :", pUname)
	fmt.Println("Sever running at port ", pPort)
	fmt.Println("Server Unique Address: ", b.GetUuidAddress())
	//fmt.Println(node_address)

	server := &http.Server{
		Addr:         ":" + pPort,
		Handler:      routes.Router(&b, &app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
