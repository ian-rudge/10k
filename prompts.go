package main

import (
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
)

func askForScoringDice(dice []string) []string {
	sortedDice := append([]string{}, dice...)
	slices.Sort(sortedDice)

	keep := []string{}
	prompt := &survey.MultiSelect{
		Message: "Choose scoring dice:",
		Options: sortedDice,
	}
	fmt.Println("")
	err := survey.AskOne(prompt, &keep, survey.WithValidator(survey.Required))
	handleCtrlC(err)

	return keep
}

func askDecision(diceCount int) string {
	ans := ""
	prompt := &survey.Select{
		Message: fmt.Sprintf("You have %d dice left. Choose option:", diceCount),
		Options: []string{"Roll", "Stay"},
	}
	fmt.Println("")
	err := survey.AskOne(prompt, &ans, survey.WithValidator(survey.Required))
	handleCtrlC(err)

	return ans
}
