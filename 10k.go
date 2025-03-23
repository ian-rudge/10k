package main

func main() {
	printWelcome()

	p1Name, p2Name := setup()
	p1Score, p2Score := 0, 0

	// play until someone reaches WIN_SCORE
	for {
		p1Score += play(p1Name, p1Score, 6)
		printRoundScore(p1Name, p1Score)

		if p1Score >= WIN_SCORE {
			break
		}

		p2Score += play(p2Name, p2Score, 6)
		printRoundScore(p2Name, p2Score)

		if p2Score >= WIN_SCORE {
			break
		}
	}

	// final round
	if p1Score >= WIN_SCORE {
		printFinalRoll(p2Name)
		p2Score += play(p2Name, p2Score, 6)
	} else if p2Score >= WIN_SCORE {
		printFinalRoll(p1Name)
		p1Score += play(p1Name, p1Score, 6)
	}

	// determine winner
	if p1Score > p2Score {
		printWinner(p1Name)
	} else {
		printWinner(p2Name)
	}
}
