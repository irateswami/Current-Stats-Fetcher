package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetBoxScore(game Game, password string, year string, date string, boxScores *BoxScores) {

	client := &http.Client{}

	// EXAMPLE - https://api.mysportsfeeds.com/v2.1/pull/mlb/2019-regular/games/20190330-DET-TOR/boxscore.json
	url := struct {
		unpopulated string
		year        string
		date        string
		awayTeam    string
		homeTeam    string
	}{
		"https://api.mysportsfeeds.com/v2.1/pull/mlb/<YEAR>-regular/games/<DATE>-<AWAYTEAM>-<HOMETEAM>/boxscore.json",
		year,
		date,
		game.AwayTeam,
		game.HomeTeam,
	}

	newUrl := url.unpopulated
	newUrl = strings.Replace(newUrl, "<YEAR>", url.year, -1)
	newUrl = strings.Replace(newUrl, "<DATE>", url.date, -1)
	newUrl = strings.Replace(newUrl, "<AWAYTEAM>", url.awayTeam, -1)
	newUrl = strings.Replace(newUrl, "<HOMETEAM>", url.homeTeam, -1)

	fmt.Printf("%+v\n", newUrl)

	//request, err := http.NewRequest("GET", url.populateBoxScoreUrl(), nil)
	request, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		log.Printf("request failed: %v\n", err)
	}

	request.SetBasicAuth(password, "MYSPORTSFEEDS")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("response failed: %v\n", err)
	}

	boxScore := new(BoxScore)
	json.NewDecoder(response.Body).Decode(boxScore)

	boxScores.Mut.Lock()
	boxScores.BoxScores = append(boxScores.BoxScores, *boxScore)
	boxScores.Mut.Unlock()

}
