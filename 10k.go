package main

func main() {
	printWelcome()

	p1, p2 := "", ""
	p1Name, p2Name := setup()
	p1Score, p2Score := 0, 0
	p1OnBoard, p2OnBoard := false, false

	printRollFirst()

	for {
		p1Roll, p2Roll := rollForFirstTurn(p1Name), rollForFirstTurn(p2Name)
		if p1Roll[0] > p2Roll[0] {
			p1, p2 = p1Name, p2Name
			break
		}
		if p2Roll[0] > p1Roll[0] {
			p1, p2 = p2Name, p1Name
			break
		}
	}

	printPlayerRollFirst(p1)

	for {
		if playRound(p1, &p1Score, &p1OnBoard) {
			break
		}
		if playRound(p2, &p2Score, &p2OnBoard) {
			break
		}
	}

	// final round
	if p1Score >= WIN_SCORE {
		printFinalRoll(p2)
		p2Score += playGame(p2, 0, ALL_DICE)
	} else if p2Score >= WIN_SCORE {
		printFinalRoll(p1)
		p1Score += playGame(p1, 0, ALL_DICE)
	}

	if p1Score > p2Score {
		printWinner(p1)
	} else if p2Score > p1Score {
		printWinner(p2)
	} else {
		printTie()
	}

	printFinalScore(p1, p1Score)
	printFinalScore(p2, p2Score)
}
