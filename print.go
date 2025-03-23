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

func printRollFirst() {
	fmt.Println("")
	fmt.Println("Rolling to see who goes first.")
	fmt.Println("")
}

func printPlayerRollFirst(name string) {
	fmt.Println("")
	fmt.Printf("%s rolls first!\n", name)
	fmt.Println("")
}

func printWelcome() {
	fmt.Println("Welcome to 10,000!")
}

func printTurn(name string) {
	color.Cyan("------------ %s's roll! ------------\n", name)
}

func printOnBoard(name string) {
	color.Yellow("%s is on the board!\n", name)
}

func printScore(score int) {
	color.Yellow("Current score: %d\n", score)
}

func printFinalScore(name string, score int) {
	color.Yellow("%s's final score: %d\n", name, score)
}

func printRoll(dice []string) {
	color.Magenta("Rolled: %s\n", strings.Join(dice, ", "))
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
	color.Red("%s's final roll!\n", name)
	fmt.Println("")
}

func printWinner(name string) {
	color.Green("%s wins!\n", name)
}
