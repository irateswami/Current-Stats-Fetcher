package api

import (
	"sync"
)

type BoxScores struct {
	BoxScores []BoxScore
	Mut       sync.Mutex
}
