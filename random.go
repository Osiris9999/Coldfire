package coldfire

import (
	"math/rand"
	"time"
)

// RandomSelectStrNested returns a string array that was randomly selected from a nested list of strings
func RandomSelectStrNested(list [][]string) []string {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}
