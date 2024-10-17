package usecase

import (
	"bingo/model"
	"fmt"
)

type Simulator struct {
	cards    int
	count    int64
	withFree bool
}

func NewSimulator(cards int, count int64, withFree bool) Simulator {
	return Simulator{cards, count, withFree}
}

type Result struct {
	rowBingoCount      int64
	columnBingoCount   int64
	diagonalBingoCount int64
	rollsCount         []int
}

func (s Simulator) Exec() error {
	var i int64

	fullResult := Result{
		rollsCount: make([]int, s.count),
	}

	var rollsMin, rollsMax int
	var rollsSum int64
	rollsMin = 75

	for i = 0; i < s.count; i++ {
		r, err := s.exec()
		if err != nil {
			return err
		}

		if r.rowBingo {
			fullResult.rowBingoCount++
		}
		if r.columnBingo {
			fullResult.columnBingoCount++
		}
		if r.diagonalBingo {
			fullResult.diagonalBingoCount++
		}
		fullResult.rollsCount[i] = r.rollsCount

		if rollsMin > r.rollsCount {
			rollsMin = r.rollsCount
		}
		if rollsMax < r.rollsCount {
			rollsMax = r.rollsCount
		}
		rollsSum += int64(r.rollsCount)
	}
	avg := float64(rollsSum) / float64(s.count)

	fmt.Printf(
		"Row    Bingo: %d\n"+
			"Column Bingo: %d\n"+
			"Diag   Bingo: %d\n\n"+
			"=== Rolls ===\n"+
			"Total: %d\n"+
			"Avg:   %v\n"+
			"Max:   %d\n"+
			"Min:   %d\n",
		fullResult.rowBingoCount,
		fullResult.columnBingoCount,
		fullResult.diagonalBingoCount,
		s.count, avg, rollsMax, rollsMin,
	)

	return nil
}

type result struct {
	rowBingo      bool
	columnBingo   bool
	diagonalBingo bool
	rollsCount    int
}

func (s Simulator) exec() (*result, error) {
	cards := make([]model.Card, s.cards)
	for i := range cards {
		cards[i] = model.NewCard(s.withFree)
	}

	machine := model.NewMachine()
	rs := result{}

	for {
		rolled := machine.Roll()
		if rolled == nil {
			return nil, fmt.Errorf("everything is rolled out")
		}

		bingo := false
		for i := range cards {
			if cards[i].FindAndHit(*rolled) {
				if cards[i].RowBingo() {
					rs.rowBingo = true
					bingo = true
				}
				if cards[i].ColumnBingo() {
					rs.columnBingo = true
					bingo = true
				}
				if cards[i].DiagonalBingo() {
					rs.diagonalBingo = true
					bingo = true
				}
			}
		}

		rs.rollsCount++
		if bingo {
			break
		}
	}

	return &rs, nil
}
