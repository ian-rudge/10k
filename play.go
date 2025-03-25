package main

import (
	"math/rand/v2"
	"strconv"
)

var ALL_DICE = 6

func rollForFirstTurn(playerName string) []string {
	printTurn(playerName)
	roll := roll(1)
	printRoll(roll)
	return roll
}

func playRound(playerName string, playerScore *int, onBoard *bool) bool {
	printTurn(playerName)
	printTotalScore(playerName, *playerScore)
	roundScore := playGame(playerName, 0, ALL_DICE)
	if !*onBoard && roundScore >= ON_BOARD_SCORE {
		*onBoard = true
	}
	if *onBoard {
		*playerScore += roundScore
		printRoundOverScore(playerName, *playerScore)
	} else {
		printRoundOverOnBoard(playerName)
	}

	if *playerScore >= WIN_SCORE {
		return true
	}

	return false
}

func playGame(playerName string, currentScore int, diceCount int) int {
	printRoundScore(currentScore)

	if diceCount < ALL_DICE {
		ans := askDecision(diceCount)
		if ans == "Stay" {
			return currentScore
		}
	}

	dice := roll(diceCount)
	printRoll(dice)

	_, potentialScoredDice := calculateScore(getRollCount(dice))
	if potentialScoredDice == 0 {
		printBust(currentScore)
		return 0
	}

	keep := askForScoringDice(dice)
	keptDice := getRollCount(keep)

	if checkRun(keptDice) {
		printRun(playerName)
		updatedScore := currentScore + ONE_THOUSAND
		return playGame(playerName, updatedScore, ALL_DICE)
	}

	if check3pair(keptDice) {
		print3pairs(playerName)
		updatedScore := currentScore + ONE_THOUSAND
		return playGame(playerName, updatedScore, ALL_DICE)
	}

	score, scoredDice := calculateScore(keptDice)
	if scoredDice == 0 {
		printBust(currentScore)
		return 0
	}

	printScoringRolls(playerName, getScoringRollsText(keptDice))

	updatedScore := currentScore + score

	if diceCount == scoredDice {
		return playGame(playerName, updatedScore, ALL_DICE)
	}

	availDice := diceCount - scoredDice
	return playGame(playerName, updatedScore, availDice)
}

func roll(num int) []string {
	dice := make([]string, num)
	for i := range dice {
		dice[i] = strconv.Itoa(rand.IntN(ALL_DICE) + 1)
	}
	return dice
}

// Returns a map of the count of each dice value
func getRollCount(dice []string) map[string]int {
	count := map[string]int{}
	for _, d := range dice {
		count[d]++
	}
	return count
}
