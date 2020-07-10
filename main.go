package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	//"time"

	"github.com/irateswami/Current-Stats-Fetcher/api"
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
	todaysGames := api.GetDailyGames(configs.ApiKey, year, date)

	boxScores := new(api.BoxScores)

	wg.Add(len(todaysGames))

	for i := range todaysGames {
		go func(goI int) {
			defer wg.Done()
			api.GetBoxScore(todaysGames[goI], configs.ApiKey, year, date, boxScores)
		}(i)
	}
	wg.Wait()

	fmt.Printf("%+v\n", len(boxScores.BoxScores))
}
