package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func roll(num int) []string {
	dice := make([]string, num)
	for i := range dice {
		dice[i] = strconv.Itoa(rand.IntN(6) + 1)
	}
	return dice
}

func hasOneOrFive(dice []string) bool {
	for _, num := range dice {
		if num == "1" || num == "5" {
			return true
		}
	}
	return false
}

// Returns a map of the count of each dice value
func getRollCount(dice []string) map[string]int {
	count := map[string]int{}

	for _, d := range dice {
		count[d]++
	}

	return count
}

func play(name string, currentScore int, diceCount int) int {
	printTurn(name)

	if diceCount < 6 {
		ans := ""
		prompt := &survey.Select{
			Message: fmt.Sprintf("You have %d dice left. Choose option:", diceCount),
			Options: []string{"Roll", "Stay"},
		}
		err := survey.AskOne(prompt, &ans, survey.WithValidator(survey.Required))
		handleCtrlC(err)

		if ans == "Stay" {
			return currentScore
		}
	}

	dice := roll(diceCount)
	printRoll(dice)

	if !hasOneOrFive(dice) {
		printBust()
		return 0
	}

	sortedDice := make([]string, len(dice))
	copy(sortedDice, dice)
	slices.Sort(sortedDice)

	keep := []string{}
	prompt := &survey.MultiSelect{
		Message: "Choose scoring dice:",
		Options: sortedDice,
	}

	err := survey.AskOne(prompt, &keep, survey.WithValidator(survey.Required))
	handleCtrlC(err)

	keptDice := getRollCount(keep)

	if checkRun(keptDice) {
		printRun(name)
		updatedScore := currentScore + ONE_THOUSAND
		printScore(updatedScore)
		return play(name, updatedScore, 6)
	}

	if check3pair(keptDice) {
		print3pairs(name)
		updatedScore := currentScore + ONE_THOUSAND
		printScore(updatedScore)
		return play(name, updatedScore, 6)
	}

	score, scoredDice := calculateScore(keptDice)
	printScore(currentScore + score)
	availDice := diceCount - scoredDice

	return play(name, currentScore+score, availDice)
}
