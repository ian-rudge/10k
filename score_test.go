package main

import (
	"fmt"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		dice       map[string]int
		wantScore  int
		wantScored int
	}{
		{map[string]int{"1": 1}, 100, 1},
		{map[string]int{"1": 3}, 1000, 3},
		{map[string]int{"2": 3}, 200, 3},
		{map[string]int{"5": 1}, 50, 1},
		{map[string]int{"5": 3}, 500, 3},
		{map[string]int{"1": 1, "5": 1}, 150, 2},
		{map[string]int{"1": 2, "5": 2}, 300, 4},
		{map[string]int{"1": 3, "2": 3}, 1200, 6},
		{map[string]int{"1": 4}, 2000, 4},
		{map[string]int{"2": 4}, 400, 4},
		{map[string]int{"3": 4}, 600, 4},
		{map[string]int{"4": 4}, 800, 4},
		{map[string]int{"5": 4}, 1000, 4},
		{map[string]int{"6": 4}, 1200, 4},
		{map[string]int{"2": 5}, 800, 5},
		{map[string]int{"3": 5}, 1200, 5},
		{map[string]int{"4": 5}, 1600, 5},
		{map[string]int{"5": 5}, 2000, 5},
		{map[string]int{"6": 5}, 2400, 5},
		{map[string]int{"2": 6}, 1600, 6},
		{map[string]int{"3": 6}, 2400, 6},
		{map[string]int{"4": 6}, 3200, 6},
		{map[string]int{"5": 6}, 4000, 6},
		{map[string]int{"6": 6}, 4800, 6},
		{map[string]int{"1": 2}, 200, 2},
		{map[string]int{"1": 5}, 3000, 5},
		{map[string]int{"1": 6}, 6000, 6},
		{map[string]int{"5": 2}, 100, 2},
		{map[string]int{"1": 1, "2": 1, "3": 1, "4": 1, "5": 1, "6": 1}, 150, 2},
		{map[string]int{"1": 3, "5": 3}, 1500, 6},
		{map[string]int{"1": 4, "5": 2}, 2100, 6},
		{map[string]int{"2": 2, "3": 2, "4": 2}, 0, 0},
		{map[string]int{"1": 2, "2": 2, "3": 2}, 200, 2},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Dice: %v", tt.dice), func(t *testing.T) {
			gotScore, gotScored := calculateScore(tt.dice)
			if gotScore != tt.wantScore {
				t.Errorf("calculateScore() gotScore = %v, want %v", gotScore, tt.wantScore)
			}
			if gotScored != tt.wantScored {
				t.Errorf("calculateScore() gotScored = %v, want %v", gotScored, tt.wantScored)
			}
		})
	}
}
func TestCheckRun(t *testing.T) {
	tests := []struct {
		dice map[string]int
		want bool
	}{
		{map[string]int{"1": 1, "2": 1, "3": 1, "4": 1, "5": 1, "6": 1}, true},
		{map[string]int{"1": 1, "2": 1, "3": 1, "4": 1, "5": 1}, false},
		{map[string]int{"1": 1, "2": 1, "3": 1, "4": 1, "5": 1, "7": 1}, false},
		{map[string]int{"1": 1, "2": 1, "3": 1, "4": 1, "5": 1, "6": 1, "7": 1}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Dice: %v", tt.dice), func(t *testing.T) {
			got := checkRun(tt.dice)
			if got != tt.want {
				t.Errorf("checkRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCheck3Pair(t *testing.T) {
	tests := []struct {
		dice map[string]int
		want bool
	}{
		{map[string]int{"1": 2, "2": 2, "3": 2}, true},
		{map[string]int{"1": 2, "2": 2, "3": 2, "4": 1}, true},
		{map[string]int{"1": 2, "2": 2, "3": 1, "4": 1}, false},
		{map[string]int{"1": 2, "2": 2, "3": 2, "4": 2}, true},
		{map[string]int{"1": 2, "2": 2, "3": 1, "4": 1, "5": 1}, false},
		{map[string]int{"1": 2, "2": 2, "3": 2, "4": 1, "5": 1}, true},
		{map[string]int{"1": 1, "2": 1, "3": 1, "4": 1, "5": 1, "6": 1}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Dice: %v", tt.dice), func(t *testing.T) {
			got := check3pair(tt.dice)
			if got != tt.want {
				t.Errorf("check3pair() = %v, want %v", got, tt.want)
			}
		})
	}
}
