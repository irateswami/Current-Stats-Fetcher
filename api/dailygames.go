package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Game struct {
	HomeTeam   string
	HomeTeamID int
	AwayTeam   string
	AwayTeamID int
}

func GetDailyGames(password string) []Game {
	client := &http.Client{}

	url := struct {
		unpopulated string
		year        string
		date        string
	}{
		"https://api.mysportsfeeds.com/v2.1/pull/mlb/<YEAR>-regular/date/<DATE>/games.json",
		time.Now().Format("2006"),
		time.Now().Format("20060102"),
	}

	newUrl := url.unpopulated
	newUrl = strings.Replace(newUrl, "<YEAR>", url.year, -1)
	newUrl = strings.Replace(newUrl, "<DATE>", url.date, -1)

	if request, err := http.NewRequest("GET", newUrl, nil); err != nil {
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
