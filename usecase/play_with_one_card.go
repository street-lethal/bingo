package usecase

import (
	"bingo/model"
	"fmt"
)

type PlayWithOneCardUseCase struct{}

func NewPlayWithOneCardUseCase() PlayWithOneCardUseCase {
	return PlayWithOneCardUseCase{}
}

func (u PlayWithOneCardUseCase) Play(withFree bool) error {
	card := model.NewCard(withFree)
	fmt.Println(card)

	machine := model.NewMachine()

	for {
		rolled := machine.Roll()
		if rolled == nil {
			return fmt.Errorf("everything is rolled out")
		}

		fmt.Printf("%d ", *rolled)
		if card.FindAndHit(*rolled) && card.Bingo() {
			break
		}
	}
	fmt.Print("\n\n")

	fmt.Println(card)

	return nil
}
