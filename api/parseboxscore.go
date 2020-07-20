package api

func ParseBoxScore(boxScore BoxScore) []PlayerStruct {

	var allPlayers []PlayerStruct
	awayPlayers := boxScore.Stats.Away.Players
	homePlayers := boxScore.Stats.Away.Players

	for i := range awayPlayers {
		newPlayer := awayPlayers[i]
		allPlayers = append(allPlayers, newPlayer)
	}

	for i := range homePlayers {
		newPlayer := homePlayers[i]
		allPlayers = append(allPlayers, newPlayer)
	}

	return allPlayers
}
