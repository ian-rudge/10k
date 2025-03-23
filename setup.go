package main

import (
	"github.com/AlecAivazis/survey/v2"
)

func setup() (string, string) {
	p1Name := ""
	p1Prompt := &survey.Input{
		Message: "Player 1 name",
	}
	err := survey.AskOne(p1Prompt, &p1Name)
	handleCtrlC(err)

	p2Name := ""
	p2Prompt := &survey.Input{
		Message: "Player 2 name",
	}
	err = survey.AskOne(p2Prompt, &p2Name)
	handleCtrlC(err)

	return p1Name, p2Name
}
