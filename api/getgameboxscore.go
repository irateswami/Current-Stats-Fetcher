package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type Url struct {
	unpopulated string
	year        string
	date        string
	awayTeam    string
	homeTeam    string
}

func (url Url) populateBoxScoreUrl() string {

	newUrl := url.unpopulated
	newUrl = strings.Replace(newUrl, "<YEAR>", url.year, -1)
	newUrl = strings.Replace(newUrl, "<DATE>", url.date, -1)
	newUrl = strings.Replace(newUrl, "<AWAYTEAM>", url.awayTeam, -1)
	newUrl = strings.Replace(newUrl, "<HOMETEAM>", url.homeTeam, -1)

	return newUrl
}

func GetBoxScore(game Game, password string) BoxScore {

	client := &http.Client{}

	// EXAMPLE - https://api.mysportsfeeds.com/v2.1/pull/mlb/2019-regular/games/20190330-DET-TOR/boxscore.json
	url := Url{
		unpopulated: "https://api.mysportsfeeds.com/v2.1/pull/mlb/<YEAR>-regular/games/<DATE>-<AWAYTEAM>-<HOMETEAM>/boxscore.json",
		year:        time.Now().Format("2006"),
		date:        time.Now().Format("20060102"),
		awayTeam:    game.AwayTeam,
		homeTeam:    game.HomeTeam,
	}

	request, err := http.NewRequest("GET", url.populateBoxScoreUrl(), nil)
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

	return *boxScore
}
