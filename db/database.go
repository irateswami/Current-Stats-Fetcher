package db

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/irateswami/Current-Stats-Fetcher/api"
)

func init() {

	// TODO intake secrets for db ------------------------------------------------

	// TODO connect to database --------------------------------------------------

}

func Insert(players []api.PlayerStruct) {

	for i := range players {

		// Marshal each playerStruct into a json ------------------------------------
		playerBytes, err := json.Marshal(players[i])
		if err != nil {
			log.Printf("jason marshal failed: %v\n", err)
		}

		// Convert that byte array into a io.Reader ---------------------------------
		playerReader := bytes.NewReader(playerBytes)

		// Unmarshall onto a new dbPlayerStruct -------------------------------------
		newDbPlayerStruct := new(api.PlayerStruct)
		json.NewDecoder(playerReader).Decode(&newDbPlayerStruct)

		spew.Dump(newDbPlayerStruct)

	}

}
