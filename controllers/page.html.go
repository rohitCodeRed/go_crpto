package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/config"
	"github.com/rohitCodeRed/go_crypto/model"
	"github.com/rohitCodeRed/go_crypto/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(c *gin.Context, b *blockchain.BlockChain) {
	//fmt.Fprintf(w, "this is the home page")
	render.RenderTemplate(c, "home.page.tmpl", &model.CoinData{})
}

func (m *Repository) About(c *gin.Context, b *blockchain.BlockChain) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	//m.App.Session.

	render.RenderTemplate(c, "about.page.tmpl", &model.CoinData{
		StringMap: stringMap,
	})

}
