package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func toTitleCase(s string) string {
	return cases.Title(language.English).String(s)
}

func setup() (p1Name string, p2Name string) {
	p1Name, p2Name = promptForName("Player 1 name"), promptForName("Player 2 name")
	return toTitleCase(p1Name), toTitleCase(p2Name)
}
