package model

type Card [5]Column

func NewCard(withFree bool) Card {
	var card Card
	for i := range card {
		card[i] = NewColumn(i)
	}
	if withFree {
		card[2][2].val = -1
		card[2][2].Hit()
	}
	return card
}

func (card Card) String() string {
	var s string
	for row := 0; row < 5; row++ {
		s += card[0][row].String() + card[1][row].String() + card[2][row].String() +
			card[3][row].String() + card[4][row].String()
		s += "\n"
	}

	return s
}

func (card *Card) FindAndHit(n int) bool {
	if n < 45 {
		if n < 15 {
			return card[0].FindAndHit(n)
		} else if n < 30 {
			return card[1].FindAndHit(n)
		} else {
			return card[2].FindAndHit(n)
		}
	} else {
		if n < 60 {
			return card[3].FindAndHit(n)
		} else {
			return card[4].FindAndHit(n)
		}
	}
}

func (card Card) ColumnBingo() bool {
	for _, column := range card {
		if column.Bingo() {
			return true
		}
	}

	return false
}

func (card Card) RowBingo() bool {
	for rowNum := 0; rowNum < 5; rowNum++ {
		rowBingo := true
		for columnNum := 0; columnNum < 5; columnNum++ {
			if !card[columnNum][rowNum].IsHit() {
				rowBingo = false
				break
			}
		}
		if rowBingo {
			return true
		}
	}

	return false
}

func (card Card) DiagonalBingo() bool {
	if !card[2][2].IsHit() {
		return false
	}

	if card[0][0].IsHit() && card[1][1].IsHit() && card[3][3].IsHit() && card[4][4].IsHit() {
		return true
	}

	if card[0][4].IsHit() && card[1][3].IsHit() && card[3][1].IsHit() && card[4][0].IsHit() {
		return true
	}

	return false
}

func (card Card) Bingo() bool {
	return card.RowBingo() || card.ColumnBingo() || card.DiagonalBingo()
}
