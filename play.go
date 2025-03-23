package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func printRoundScore(name string, score int) {
	fmt.Printf("%s's score: %d\n", name, score)
}

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

func getRollCount(dice []string) map[string]int {
	count := map[string]int{}

	for _, d := range dice {
		count[d]++
	}

	return count
}

func play(name string, currentScore int, diceCount int) int {
	fmt.Printf("------ %s's roll! ------\n", name)

	dice := roll(diceCount)
	fmt.Printf("Rolled: %s\n", strings.Join(dice, ", "))

	if !hasOneOrFive(dice) {
		fmt.Printf("%s rolled no 1s or 5s.\n", name)
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
		fmt.Printf("%s rolled a run!\n", name)
		return play(name, currentScore+1000, diceCount-len(keep))
	}

	if check3pair(keptDice) {
		fmt.Printf("%s rolled 3 pair!\n", name)
		return play(name, currentScore+1000, diceCount-len(keep))
	}

	score := 0
	for num, cnt := range keptDice {
		score += scoreMap[fmt.Sprintf("%s:%d", num, cnt)]
	}

	return play(name, currentScore+score, diceCount-len(keep))
}
