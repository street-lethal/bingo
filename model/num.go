package model

import "fmt"

type Num struct {
	val int
	hit bool
}

func NewNum(num int) Num {
	return Num{
		val: num,
		hit: false,
	}
}

func (num Num) Value() int {
	return num.val
}

func (num *Num) Hit() {
	num.hit = true
}

func (num Num) IsHit() bool {
	return num.hit
}

func (num Num) String() string {
	var s string

	if num.IsHit() {
		s += " *"
	} else {
		s += "  "
	}
	s += fmt.Sprintf("%02d", num.val)

	return s
}
