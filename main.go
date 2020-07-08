package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/irateswami/Current-Stats-Fetcher/api"
)

var (
	configs Configs
)

type BoxScores struct {
	BoxScores []api.BoxScore
	Mut       sync.Mutex
}

func init() {

	// Open the file
	secrets, err := os.Open("configs.json")
	defer secrets.Close()
	if err != nil {
		log.Fatalf("no secrets found: %v\n", err)
	}

	// Read in the secrets
	bytesSecrets, err := ioutil.ReadAll(secrets)
	if err != nil {
		log.Fatalf("secrets failed byte conversion: %v\n", err)
	}

	json.Unmarshal(bytesSecrets, &configs)

}

func main() {
	//todaysGames := api.GetDailyGames(configs.ApiKey)

	//var boxScores BoxScores
	//	playerStats := api.GetBoxScore(todaysGames, configs.ApiKey)
	//fmt.Printf("%+v\n", playerStats)
}
