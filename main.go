package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/irateswami/Current-Stats-Fetcher/api"
	//	"github.com/davecgh/go-spew/spew"
	//	"github.com/irateswami/Current-Stats-Fetcher/db"
)

var (
	configs Configs
	date    string
	year    string
)

func init() {

	// Get the date --------------------------------------------------------------
	//year = time.Now().Format("2006")
	//date = time.Now().Format("20060102")

	// TODO remove this, for now it's just testing things
	date = "20190330"
	year = "2019"

	// Open the file -------------------------------------------------------------
	secrets, err := os.Open("secrets.json")
	defer secrets.Close()
	if err != nil {
		log.Fatalf("no secrets found: %v\n", err)
	}

	// Read in the secrets -------------------------------------------------------
	bytesSecrets, err := ioutil.ReadAll(secrets)
	if err != nil {
		log.Fatalf("secrets failed byte conversion: %v\n", err)
	}

	json.Unmarshal(bytesSecrets, &configs)

}

func main() {

	playerStats := api.GetPlayerStats(configs.APIKey, year, date)

	_ = playerStats

}
