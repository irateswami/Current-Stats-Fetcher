package api

import "time"

type PlayerStats struct {
	LastUpdatedOn     time.Time `json:"lastUpdatedOn"`
	PlayerStatsTotals []struct {
		Player struct {
			ID              int    `json:"id"`
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			PrimaryPosition string `json:"primaryPosition"`
			JerseyNumber    int    `json:"jerseyNumber"`
			CurrentTeam     struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"currentTeam"`
			CurrentRosterStatus string      `json:"currentRosterStatus"`
			CurrentInjury       interface{} `json:"currentInjury"`
			Height              string      `json:"height"`
			Weight              int         `json:"weight"`
			BirthDate           string      `json:"birthDate"`
			Age                 int         `json:"age"`
			BirthCity           string      `json:"birthCity"`
			BirthCountry        string      `json:"birthCountry"`
			Rookie              bool        `json:"rookie"`
			HighSchool          interface{} `json:"highSchool"`
			College             interface{} `json:"college"`
			Handedness          struct {
				Bats   string `json:"bats"`
				Throws string `json:"throws"`
			} `json:"handedness"`
			OfficialImageSrc    string        `json:"officialImageSrc"`
			SocialMediaAccounts []interface{} `json:"socialMediaAccounts"`
		} `json:"player"`
		Team struct {
			ID           int    `json:"id"`
			Abbreviation string `json:"abbreviation"`
		} `json:"team"`
		Stats struct {
			GamesPlayed int `json:"gamesPlayed"`
			Batting     struct {
				AtBats                       int     `json:"atBats"`
				Runs                         int     `json:"runs"`
				Hits                         int     `json:"hits"`
				SecondBaseHits               int     `json:"secondBaseHits"`
				ThirdBaseHits                int     `json:"thirdBaseHits"`
				Homeruns                     int     `json:"homeruns"`
				EarnedRuns                   int     `json:"earnedRuns"`
				UnearnedRuns                 int     `json:"unearnedRuns"`
				RunsBattedIn                 int     `json:"runsBattedIn"`
				BatterWalks                  int     `json:"batterWalks"`
				BatterSwings                 int     `json:"batterSwings"`
				BatterStrikes                int     `json:"batterStrikes"`
				BatterStrikesFoul            int     `json:"batterStrikesFoul"`
				BatterStrikesMiss            int     `json:"batterStrikesMiss"`
				BatterStrikesLooking         int     `json:"batterStrikesLooking"`
				BatterTagOuts                int     `json:"batterTagOuts"`
				BatterForceOuts              int     `json:"batterForceOuts"`
				BatterPutOuts                int     `json:"batterPutOuts"`
				BatterGroundBalls            int     `json:"batterGroundBalls"`
				BatterFlyBalls               int     `json:"batterFlyBalls"`
				BatterLineDrives             int     `json:"batterLineDrives"`
				Batter2SeamFastballs         int     `json:"batter2SeamFastballs"`
				Batter4SeamFastballs         int     `json:"batter4SeamFastballs"`
				BatterCurveballs             int     `json:"batterCurveballs"`
				BatterChangeups              int     `json:"batterChangeups"`
				BatterCutters                int     `json:"batterCutters"`
				BatterSliders                int     `json:"batterSliders"`
				BatterSinkers                int     `json:"batterSinkers"`
				BatterSplitters              int     `json:"batterSplitters"`
				BatterStrikeouts             int     `json:"batterStrikeouts"`
				StolenBases                  int     `json:"stolenBases"`
				CaughtBaseSteals             int     `json:"caughtBaseSteals"`
				BatterStolenBasePct          float64 `json:"batterStolenBasePct"`
				BattingAvg                   float64 `json:"battingAvg"`
				BatterOnBasePct              float64 `json:"batterOnBasePct"`
				BatterSluggingPct            float64 `json:"batterSluggingPct"`
				BatterOnBasePlusSluggingPct  float64 `json:"batterOnBasePlusSluggingPct"`
				BatterIntentionalWalks       int     `json:"batterIntentionalWalks"`
				HitByPitch                   int     `json:"hitByPitch"`
				BatterSacrificeBunts         int     `json:"batterSacrificeBunts"`
				BatterSacrificeFlies         int     `json:"batterSacrificeFlies"`
				TotalBases                   int     `json:"totalBases"`
				ExtraBaseHits                int     `json:"extraBaseHits"`
				BatterDoublePlays            int     `json:"batterDoublePlays"`
				BatterTriplePlays            int     `json:"batterTriplePlays"`
				BatterGroundOuts             int     `json:"batterGroundOuts"`
				BatterFlyOuts                int     `json:"batterFlyOuts"`
				BatterGroundOutToFlyOutRatio float64 `json:"batterGroundOutToFlyOutRatio"`
				PitchesFaced                 int     `json:"pitchesFaced"`
				PlateAppearances             int     `json:"plateAppearances"`
				LeftOnBase                   int     `json:"leftOnBase"`
			} `json:"batting"`
			Pitching struct {
				Wins                          int     `json:"wins"`
				Losses                        int     `json:"losses"`
				EarnedRunAvg                  float64 `json:"earnedRunAvg"`
				Saves                         int     `json:"saves"`
				SaveOpportunities             int     `json:"saveOpportunities"`
				InningsPitched                float64 `json:"inningsPitched"`
				HitsAllowed                   int     `json:"hitsAllowed"`
				SecondBaseHitsAllowed         int     `json:"secondBaseHitsAllowed"`
				ThirdBaseHitsAllowed          int     `json:"thirdBaseHitsAllowed"`
				RunsAllowed                   int     `json:"runsAllowed"`
				EarnedRunsAllowed             int     `json:"earnedRunsAllowed"`
				HomerunsAllowed               int     `json:"homerunsAllowed"`
				PitcherWalks                  int     `json:"pitcherWalks"`
				PitcherSwings                 int     `json:"pitcherSwings"`
				PitcherStrikes                int     `json:"pitcherStrikes"`
				PitcherStrikesFoul            int     `json:"pitcherStrikesFoul"`
				PitcherStrikesMiss            int     `json:"pitcherStrikesMiss"`
				PitcherStrikesLooking         int     `json:"pitcherStrikesLooking"`
				PitcherGroundBalls            int     `json:"pitcherGroundBalls"`
				PitcherFlyBalls               int     `json:"pitcherFlyBalls"`
				PitcherLineDrives             int     `json:"pitcherLineDrives"`
				Pitcher2SeamFastballs         int     `json:"pitcher2SeamFastballs"`
				Pitcher4SeamFastballs         int     `json:"pitcher4SeamFastballs"`
				PitcherCurveballs             int     `json:"pitcherCurveballs"`
				PitcherChangeups              int     `json:"pitcherChangeups"`
				PitcherCutters                int     `json:"pitcherCutters"`
				PitcherSliders                int     `json:"pitcherSliders"`
				PitcherSinkers                int     `json:"pitcherSinkers"`
				PitcherSplitters              int     `json:"pitcherSplitters"`
				PitcherSacrificeBunts         int     `json:"pitcherSacrificeBunts"`
				PitcherSacrificeFlies         int     `json:"pitcherSacrificeFlies"`
				PitcherStrikeouts             int     `json:"pitcherStrikeouts"`
				PitchingAvg                   float64 `json:"pitchingAvg"`
				WalksAndHitsPerInningPitched  float64 `json:"walksAndHitsPerInningPitched"`
				CompletedGames                int     `json:"completedGames"`
				Shutouts                      int     `json:"shutouts"`
				BattersHit                    int     `json:"battersHit"`
				PitcherIntentionalWalks       int     `json:"pitcherIntentionalWalks"`
				GamesFinished                 int     `json:"gamesFinished"`
				Holds                         int     `json:"holds"`
				PitcherDoublePlays            int     `json:"pitcherDoublePlays"`
				PitcherTriplePlays            int     `json:"pitcherTriplePlays"`
				PitcherGroundOuts             int     `json:"pitcherGroundOuts"`
				PitcherFlyOuts                int     `json:"pitcherFlyOuts"`
				PitcherWildPitches            int     `json:"pitcherWildPitches"`
				Balks                         int     `json:"balks"`
				PitcherStolenBasesAllowed     int     `json:"pitcherStolenBasesAllowed"`
				PitcherCaughtStealing         int     `json:"pitcherCaughtStealing"`
				PickoffAttempts               int     `json:"pickoffAttempts"`
				Pickoffs                      int     `json:"pickoffs"`
				TotalBattersFaced             int     `json:"totalBattersFaced"`
				PitchesThrown                 int     `json:"pitchesThrown"`
				WinPct                        float64 `json:"winPct"`
				PitcherGroundOutToFlyOutRatio float64 `json:"pitcherGroundOutToFlyOutRatio"`
				PitcherOnBasePct              float64 `json:"pitcherOnBasePct"`
				PitcherSluggingPct            float64 `json:"pitcherSluggingPct"`
				PitcherOnBasePlusSluggingPct  float64 `json:"pitcherOnBasePlusSluggingPct"`
				StrikeoutsPer9Innings         float64 `json:"strikeoutsPer9Innings"`
				WalksAllowedPer9Innings       float64 `json:"walksAllowedPer9Innings"`
				HitsAllowedPer9Innings        float64 `json:"hitsAllowedPer9Innings"`
				StrikeoutsToWalksRatio        float64 `json:"strikeoutsToWalksRatio"`
				PitchesPerInning              float64 `json:"pitchesPerInning"`
				PitcherAtBats                 int     `json:"pitcherAtBats"`
			} `json:"pitching"`
			Fielding struct {
				InningsPlayed             float64 `json:"inningsPlayed"`
				TotalChances              int     `json:"totalChances"`
				FielderTagOuts            int     `json:"fielderTagOuts"`
				FielderForceOuts          int     `json:"fielderForceOuts"`
				FielderPutOuts            int     `json:"fielderPutOuts"`
				OutsFaced                 int     `json:"outsFaced"`
				Assists                   int     `json:"assists"`
				Errors                    int     `json:"errors"`
				FielderDoublePlays        int     `json:"fielderDoublePlays"`
				FielderTriplePlays        int     `json:"fielderTriplePlays"`
				FielderStolenBasesAllowed int     `json:"fielderStolenBasesAllowed"`
				FielderCaughtStealing     int     `json:"fielderCaughtStealing"`
				FielderStolenBasePct      float64 `json:"fielderStolenBasePct"`
				PassedBalls               int     `json:"passedBalls"`
				FielderWildPitches        int     `json:"fielderWildPitches"`
				FieldingPct               float64 `json:"fieldingPct"`
				RangeFactor               float64 `json:"rangeFactor"`
			} `json:"fielding"`
			Miscellaneous struct {
				GamesStarted int `json:"gamesStarted"`
			} `json:"miscellaneous"`
		} `json:"stats"`
	} `json:"playerStatsTotals"`
	References struct {
		TeamReferences []struct {
			ID                  int           `json:"id"`
			City                string        `json:"city"`
			Name                string        `json:"name"`
			Abbreviation        string        `json:"abbreviation"`
			HomeVenue           interface{}   `json:"homeVenue"`
			TeamColoursHex      []interface{} `json:"teamColoursHex"`
			SocialMediaAccounts []struct {
				MediaType string `json:"mediaType"`
				Value     string `json:"value"`
			} `json:"socialMediaAccounts"`
			OfficialLogoImageSrc interface{} `json:"officialLogoImageSrc"`
		} `json:"teamReferences"`
		PlayerReferences []struct {
			ID              int    `json:"id"`
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			PrimaryPosition string `json:"primaryPosition"`
			JerseyNumber    int    `json:"jerseyNumber"`
			CurrentTeam     struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"currentTeam"`
			CurrentRosterStatus string      `json:"currentRosterStatus"`
			CurrentInjury       interface{} `json:"currentInjury"`
			Height              string      `json:"height"`
			Weight              int         `json:"weight"`
			BirthDate           string      `json:"birthDate"`
			Age                 int         `json:"age"`
			BirthCity           string      `json:"birthCity"`
			BirthCountry        string      `json:"birthCountry"`
			Rookie              bool        `json:"rookie"`
			HighSchool          interface{} `json:"highSchool"`
			College             string      `json:"college"`
			Handedness          struct {
				Bats   string `json:"bats"`
				Throws string `json:"throws"`
			} `json:"handedness"`
			OfficialImageSrc    string        `json:"officialImageSrc"`
			SocialMediaAccounts []interface{} `json:"socialMediaAccounts"`
		} `json:"playerReferences"`
		PlayerStatReferences []struct {
			Category     string `json:"category"`
			FullName     string `json:"fullName"`
			Description  string `json:"description"`
			Abbreviation string `json:"abbreviation"`
			Type         string `json:"type"`
		} `json:"playerStatReferences"`
	} `json:"references"`
}
