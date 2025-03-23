package main

import (
	"fmt"
	"strconv"
)

var WIN_SCORE = 10000
var ONE_THOUSAND = 1000

var scoreMap = map[string]int{
	// 1's
	"1:1": 100,
	"1:2": 200,

	// 3 of a kinds
	"1:3": 1000,
	"2:3": 200,
	"3:3": 300,
	"4:3": 400,
	"5:3": 500,
	"6:3": 600,

	// 4 of a kinds
	"1:4": 2000,
	"2:4": 400,
	"3:4": 600,
	"4:4": 800,
	"5:4": 1000,
	"6:4": 1200,

	// 5 of a kinds
	"1:5": 3000,
	"2:5": 800,
	"3:5": 1200,
	"4:5": 1600,
	"5:5": 2000,
	"6:5": 2400,

	// 6 of a kinds
	"1:6": 6000,
	"2:6": 1600,
	"3:6": 2400,
	"4:6": 3200,
	"5:6": 4000,
	"6:6": 4800,

	// 5's
	"5:1": 50,
	"5:2": 100,
}

func calculateScore(dice map[string]int) (score int, scoredDice int) {
	for num, cnt := range dice {
		points, exists := scoreMap[fmt.Sprintf("%s:%d", num, cnt)]
		if exists {
			score += points
			scoredDice += cnt
		}
	}
	return score, scoredDice
}

func checkRun(dice map[string]int) bool {
	if len(dice) != 6 {
		return false
	}

	for i := 1; i <= 6; i++ {
		die := dice[strconv.Itoa(i)]
		if die != 1 {
			return false
		}
	}
	return true
}

func check3pair(dice map[string]int) bool {
	if len(dice) < 3 {
		return false
	}

	pairCount := 0
	for _, count := range dice {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount >= 3
}
