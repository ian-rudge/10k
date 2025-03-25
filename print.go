package main

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

var system = color.New(color.FgWhite).PrintfFunc()
var info = color.New(color.FgCyan).PrintfFunc()
var success = color.New(color.FgGreen).PrintfFunc()
var warning = color.New(color.FgYellow).PrintfFunc()
var danger = color.New(color.FgRed).PrintfFunc()

func printRoundOverScore(name string, score int) {
	fmt.Println("")
	success("------------ ROUND OVER  ------------\n")
	printTotalScore(name, score)
	fmt.Println("")
}

func printRoundOverOnBoard(name string) {
	fmt.Println("")
	success("------------ ROUND OVER! ------------\n")
	warning("%s needs to score %d to get on the board!\n", name, ON_BOARD_SCORE)
	fmt.Println("")
}

func printRollFirst() {
	fmt.Println("")
	system("Rolling to see who goes first.\n")
	fmt.Println("")
}

func printPlayerRollFirst(name string) {
	fmt.Println("")
	system("%s rolls first!\n", name)
	fmt.Println("")
}

func printWelcome() {
	figure.NewColorFigure("10k!", "", "green", true).Print()
	fmt.Println("")
}

func printTurn(name string) {
	system("------------ %s's roll! ------------\n", name)
}

func printTotalScore(name string, score int) {
	success("%s's total score: %d\n", name, score)
}

func printRoundScore(score int) {
	info("Round score: %d\n", score)
}

func printFinalScore(name string, score int) {
	info("%s's final score: %d\n", name, score)
}

func printRoll(dice []string) {
	info("Rolled: %s\n", strings.Join(dice, ", "))
}

func printBust(lostScore int) {
	if lostScore > 0 {
		danger("OOF! You busted and lost %d points!\n", lostScore)
	} else {
		danger("OOF! You busted!\n")
	}
}

func printRun(name string) {
	success("%s scored a run!\n", name)
}

func printScoringRolls(name string, rolls []string) {
	success("%s scored %s\n", name, strings.Join(rolls, ", "))
}

func print3pairs(name string) {
	success("%s scored 3 pairs!\n", name)
}

func printFinalRoll(name string) {
	info("%s's final roll!\n", name)
	fmt.Println("")
}

func printTie() {
	success("It's a tie!\n")
}

func printWinner(name string) {
	success("%s wins!\n", name)
}
