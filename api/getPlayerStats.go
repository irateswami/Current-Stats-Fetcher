package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// GetPlayerStats Grabs all the player stat totals from a certain date and year
func GetPlayerStats(password, year, date string) PlayerStats {

	client := &http.Client{}

	url := struct {
		unpopulated string
		year        string
		date        string
	}{
		"https://api.mysportsfeeds.com/v2.1/pull/mlb/<YEAR>-regular/player_stats_totals.json?Date=<DATE>",
		date,
		year,
	}

	newURL := url.unpopulated
	newURL = strings.Replace(newURL, "<DATE>", url.date, -1)
	newURL = strings.Replace(newURL, "<YEAR>", url.date, -1)

	//request, err := http.NewRequest("GET", url.populateBoxScoreUrl(), nil)
	request, err := http.NewRequest("GET", newURL, nil)
	if err != nil {
		log.Printf("request failed: %v\n", err)
	}

	request.SetBasicAuth(password, "MYSPORTSFEEDS")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("response failed: %v\n", err)
	}

	playerStats := new(PlayerStats)

	json.NewDecoder(response.Body).Decode(&playerStats)

	return *playerStats

}
