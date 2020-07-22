package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	//"time"

	//	"github.com/davecgh/go-spew/spew"
	"github.com/irateswami/Current-Stats-Fetcher/api"
	//	"github.com/irateswami/Current-Stats-Fetcher/database"
)

var (
	configs Configs
	year    string
	date    string
	wg      sync.WaitGroup
)

func init() {

	//year = time.Now().Format("2006")
	//date = time.Now().Format("20060102")

	// TODO remove this, for now it's just testing things
	year = "2019"
	date = "20190330"

	// Open the file
	secrets, err := os.Open("secrets.json")
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

	// Get the games for today ---------------------------------------------------
	todaysGames := api.GetDailyGames(configs.ApiKey, year, date)

	boxScores := new(api.BoxScores)

	wg.Add(len(todaysGames))

	// Create a go routine for each game and grab the box score ------------------
	for i := range todaysGames {
		go func(goI int) {
			defer wg.Done()
			api.GetBoxScore(todaysGames[goI], configs.ApiKey, year, date, boxScores)
		}(i)
	}

	wg.Wait()

	// Parse the box scores and grab all the player stats ------------------------
	var playerStats []api.PlayerStruct

	for i := range boxScores.BoxScores {
		stats := api.ParseBoxScore(boxScores.BoxScores[i])
		for j := range stats {
			playerStats = append(playerStats, stats[j])
		}
	}
}
