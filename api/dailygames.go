package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Game struct {
	HomeTeam   string
	HomeTeamID int
	AwayTeam   string
	AwayTeamID int
}

func GetDailyGames(password string) []Game {
	client := &http.Client{}

	// TODO: dynamically allocate dates in url
	request, err := http.NewRequest("GET", "https://api.mysportsfeeds.com/v2.1/pull/mlb/2019-regular/date/20190330/games.json", nil)
	if err != nil {
		log.Printf("request failed: %v\n", err)
	}

	request.SetBasicAuth(password, "MYSPORTSFEEDS")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("response failed: %v\n", err)
	}

	if response.StatusCode != 200 {
		log.Printf("bad status code: %v\n", response.Status)
	}

	games := new(Games)

	json.NewDecoder(response.Body).Decode(&games)

	var allDailyGames []Game

	for i := range games.Games {
		game := Game{
			HomeTeam:   games.Games[i].Schedule.HomeTeam.Abbreviation,
			HomeTeamID: games.Games[i].Schedule.HomeTeam.ID,
			AwayTeam:   games.Games[i].Schedule.AwayTeam.Abbreviation,
			AwayTeamID: games.Games[i].Schedule.AwayTeam.ID,
		}
		allDailyGames = append(allDailyGames, game)
	}

	return allDailyGames
}
