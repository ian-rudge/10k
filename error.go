package main

import (
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
)

func handleCtrlC(err error) {
	if err != nil && err == terminal.InterruptErr {
		os.Exit(1)
	}
}
