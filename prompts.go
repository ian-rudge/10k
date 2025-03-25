package main

import (
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
)

func promptForName(message string) (name string) {
	prompt := &survey.Input{
		Message: message,
	}
	err := survey.AskOne(prompt, &name)
	handleCtrlC(err)
	return name
}

func askForScoringDice(dice []string) []string {
	sortedDice := append([]string{}, dice...)
	slices.Sort(sortedDice)

	keep := []string{}
	prompt := &survey.MultiSelect{
		Message: "Choose scoring dice:",
		Options: sortedDice,
	}
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
	err := survey.AskOne(prompt, &ans, survey.WithValidator(survey.Required))
	handleCtrlC(err)

	return ans
}
