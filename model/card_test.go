package model

import (
	"reflect"
	"testing"
)

func TestCard_FindAndHit(t *testing.T) {
	tests := []struct {
		name string
		card Card
		nums []int
		hit  [5][5]bool
		want []bool
	}{
		{
			name: "0",
			card: Card{
				Column{
					NewNum(1),
					NewNum(0),
					NewNum(6),
					NewNum(10),
					NewNum(4),
				},
				Column{
					NewNum(22),
					NewNum(26),
					NewNum(19),
					NewNum(15),
					NewNum(23),
				},
				Column{
					NewNum(36),
					NewNum(40),
					NewNum(43),
					NewNum(34),
					NewNum(44),
				},
				Column{
					NewNum(49),
					NewNum(47),
					NewNum(52),
					NewNum(53),
					NewNum(56),
				},
				Column{
					NewNum(62),
					NewNum(61),
					NewNum(69),
					NewNum(64),
					NewNum(60),
				},
			},
			nums: []int{2, 10, 15, 25, 26, 30, 44, 45, 49, 53, 60, 69},
			hit: [5][5]bool{
				{false, false, false, true, false},
				{false, true, false, true, false},
				{false, false, false, false, true},
				{true, false, false, true, false},
				{false, false, true, false, true},
			},
			want: []bool{
				false, true, true, false, true, false, true,
				false, true, true, true, true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]bool, len(tt.nums))
			for i, targetNum := range tt.nums {
				result[i] = tt.card.FindAndHit(targetNum)
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"FindAndHit() got: %v, want: %v", result, tt.want,
				)
			}
			for colNum, colHit := range tt.hit {
				col := tt.card[colNum]
				for rowNum, rowHit := range colHit {
					num := col[rowNum]
					if rowHit != num.IsHit() {
						t.Errorf(
							"FindAndHit() col: %d, row: %d, got: %v want: %v",
							colNum, rowNum, num.IsHit(), rowHit,
						)
					}
				}
			}
		})
	}
}

func TestCard_ColumnBingo(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want bool
	}{
		{
			name: "not bingo 1",
			card: Card{
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: false}, {hit: true}},
				{{hit: true}, {hit: true}, {hit: false}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: false}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: false}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
			},
			want: false,
		},
		{
			name: "not bingo 2",
			card: Card{
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false}},
			},
			want: false,
		},
		{
			name: "bingo 1",
			card: Card{
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
			},
			want: true,
		},
		{
			name: "bingo 2",
			card: Card{
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
			},
			want: true,
		},
		{
			name: "bingo 3",
			card: Card{
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.ColumnBingo(); got != tt.want {
				t.Errorf("ColumnBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_RowBingo(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want bool
	}{
		{
			name: "not bingo 1",
			card: Card{
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: false}, {hit: true}},
				{{hit: true}, {hit: true}, {hit: false}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: false}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: false}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
			},
			want: false,
		},
		{
			name: "not bingo 2",
			card: Card{
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
			},
			want: false,
		},
		{
			name: "bingo 1",
			card: Card{
				{{hit: true}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: true}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: true}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: true}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: true}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
			},
			want: true,
		},
		{
			name: "bingo 2",
			card: Card{
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
			},
			want: true,
		},
		{
			name: "bingo 3",
			card: Card{
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: true}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.RowBingo(); got != tt.want {
				t.Errorf("RowBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_DiagonalBingo(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want bool
	}{
		{
			name: "not bingo 1",
			card: Card{
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: false}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: false}, {hit: true}, {hit: true}, {hit: true}},
				{{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true}},
			},
			want: false,
		},
		{
			name: "bingo 1",
			card: Card{
				{{hit: true}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: true}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: true}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: true}},
			},
			want: true,
		},
		{
			name: "bingo 2",
			card: Card{
				{{hit: false}, {hit: false}, {hit: false}, {hit: false}, {hit: true}},
				{{hit: false}, {hit: false}, {hit: false}, {hit: true}, {hit: false}},
				{{hit: false}, {hit: false}, {hit: true}, {hit: false}, {hit: false}},
				{{hit: false}, {hit: true}, {hit: false}, {hit: false}, {hit: false}},
				{{hit: true}, {hit: false}, {hit: false}, {hit: false}, {hit: false}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.DiagonalBingo(); got != tt.want {
				t.Errorf("DiagonalBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}
