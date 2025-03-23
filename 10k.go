package main

func main() {
	printWelcome()

	var p1, p2 = "", ""
	p1Name, p2Name := setup()
	p1Score, p2Score := 0, 0

	printRollFirst()

	// determine who goes first
	for {
		printTurn(p1Name)
		p1Roll := roll(1)
		printRoll(p1Roll)

		printTurn(p2Name)
		p2Roll := roll(1)
		printRoll(p2Roll)

		if p1Roll[0] > p2Roll[0] {
			p1 = p1Name
			p2 = p2Name
			printPlayerRollFirst(p1Name)
			break
		} else if p2Roll[0] > p1Roll[0] {
			p1 = p2Name
			p2 = p1Name
			printPlayerRollFirst(p2Name)
			break
		}
	}

	// play until someone reaches WIN_SCORE
	for {
		p1Score += play(p1, 0, ALL_DICE)
		printRoundScore(p1, p1Score)

		if p1Score >= WIN_SCORE {
			break
		}

		p2Score += play(p2, 0, ALL_DICE)
		printRoundScore(p2, p2Score)

		if p2Score >= WIN_SCORE {
			break
		}
	}

	// final round
	if p1Score >= WIN_SCORE {
		printFinalRoll(p2)
		p2Score += play(p2, 0, ALL_DICE)
	} else if p2Score >= WIN_SCORE {
		printFinalRoll(p1)
		p1Score += play(p1, 0, ALL_DICE)
	}

	// determine winner
	if p1Score > p2Score {
		printWinner(p1)
	} else {
		printWinner(p2)
	}

	printFinalScore(p1, p1Score)
	printFinalScore(p2, p2Score)
}
