package stats

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RollStat generates a single stat by rolling 4d6 and dropping the lowest roll
func RollStat() int {
	dice := make([]int, 4)
	for i := range dice {
		dice[i] = rand.Intn(6) + 1 // Roll a d6
	}

	min := dice[0]
	sum := 0
	for _, value := range dice {
		sum += value
		if value < min {
			min = value
		}
	}
	return sum - min // Drop the lowest roll
}

// GenerateStats generates all character stats
func GenerateStats() map[string]int {
	stats := map[string]int{
		"Strength":     RollStat(),
		"Dexterity":    RollStat(),
		"Constitution": RollStat(),
		"Intelligence": RollStat(),
		"Wisdom":       RollStat(),
		"Charisma":     RollStat(),
	}
	return stats
}
