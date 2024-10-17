package model

import (
	"math/rand"
)

type Column [5]Num

func NewColumn(n int) Column {
	var col [5]Num

	list := rand.Perm(15)
	for i := range col {
		col[i] = NewNum(list[i] + n*15)
	}

	return col
}

func (col Column) String() string {
	var s string
	for _, num := range col {
		s += num.String()
	}
	return s
}

func (col *Column) FindAndHit(n int) bool {
	for i, num := range col {
		if n == num.Value() {
			col[i].Hit()
			return true
		}
	}
	return false
}

func (col Column) Bingo() bool {
	for _, num := range col {
		if !num.IsHit() {
			return false
		}
	}

	return true
}
