package main

import (
	"fmt"
	"strings"

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

func printTurn(name string) {
	color.Cyan("------------ %s's roll! ------------\n", name)
}

func printScore(score int) {
	color.Cyan("Current score: %d\n", score)
}

func printRoll(dice []string) {
	color.Magenta("Rolled: %s\n\n", strings.Join(dice, ", "))
}

func printBust() {
	color.Red("OOF! YOU BUSTED!\n")
}

func printRun(name string) {
	color.Green("%s scored a run!\n", name)
}

func print3pairs(name string) {
	color.Green("%s scored 3 pairs!\n", name)
}

func printFinalRoll(name string) {
	color.Cyan("%s's final roll!\n", name)
}

func printWinner(name string) {
	color.Green("%s wins!\n", name)
}
