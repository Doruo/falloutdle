package random

import (
	"math/rand"
	"time"
)

// newRandom returns a new random using seed based on today date
func NewRandom() *rand.Rand {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	return rand.New(source)
}
