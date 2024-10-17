package main

import (
	"bingo/usecase"
)

func main() {
	//u := usecase.NewPlayWithOneCardUseCase()
	//if err := u.Play(false); err != nil {
	//	panic(err)
	//}

	s := usecase.NewSimulator(5, 1000, true)
	if err := s.Exec(); err != nil {
		panic(err)
	}
}
