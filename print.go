package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printRoundScore(name string, score int) {
	fmt.Println("")
	color.Green("ROUND OVER\n")
	color.Green("%s's score: %d\n", name, score)
	fmt.Println("")
}

func printWelcome() {
	fmt.Println("Welcome to 10,000!")
}

func printScore(score int) {
	color.Cyan("Current score: %d\n", score)
}
