//Follow the Money : main.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func walletTxsDrop(ws []string) {
	var wallet_stock []wallet
	var string addr_list

	for i := 0; i < len(ws); i++ {
		addr_list := addr_list + ws[i] + ','
	}
	url := "https://blockexplorer.com/api/addrs/" + addr_list + "/txs"
	ftmClient := http.Client{
		Timeout: time.Second * 180, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "ftm")
	res, getErr := ftmClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	for i := 0; i < len(ws); i++ {
		jsonErr := json.Unmarshal(body, &wallet_stock[i])
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}
	return wallet_stock
}

func authenticate(string login, string password) {

}
func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	router.Run()
}
