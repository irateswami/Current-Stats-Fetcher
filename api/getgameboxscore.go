package api

import (
	"fmt"
	"sync"
)

func GetBoxScore(games []Game, password string) BoxScore {

	fmt.Println("length: ", len(games))

	var wg sync.WaitGroup
	url := "https://api.mysportsfeeds.com/v2.1/pull/mlb/<YEAR>-regular/games/<DATE>-<AWAYTEAM>-<HOMETEAM>/boxscore.json"

	for i := range games {
		wg.Add(1)

	}

	//	client := &http.Client{}
	//
	//	//"https://api.mysportsfeeds.com/v2.1/pull/mlb/2019-regular/games/20190330-DET-TOR/boxscore.json"
	//	url := "https://api.mysportsfeeds.com/v2.1/pull/mlb/<YEAR>-regular/games/<DATE>-<AWAYTEAM>-<HOMETEAM>/boxscore.json"
	//
	//	request, err := http.NewRequest("GET", url, nil)
	//	if err != nil {
	//		log.Printf("request failed: %v\n", err)
	//	}
	//
	//	request.SetBasicAuth(password, "MYSPORTSFEEDS")
	//
	//	response, err := client.Do(request)
	//	if err != nil {
	//		log.Printf("response failed: %v\n", err)
	//	}
	//
	boxScore := new(BoxScore)
	//	json.NewDecoder(response.Body).Decode(boxScore)

	return *boxScore
}
