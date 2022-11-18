package model

type CoinData struct {
	Uuid     string `json:"uuid"`
	UserName string `json:"name"`
	Url      string `json:"url"`
	Warning  string
	Error    string
}
