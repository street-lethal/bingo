package model

import (
	"testing"
)

func TestNewColumn(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "B",
			n:    0,
		},
		{
			name: "I",
			n:    1,
		},
		{
			name: "N",
			n:    2,
		},
		{
			name: "G",
			n:    3,
		},
		{
			name: "O",
			n:    4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewColumn(tt.n)
			duplicateChecker := map[int]bool{}
			for _, num := range got {
				val := num.Value()
				if val < tt.n*15 {
					t.Errorf("NewColumn() = %v, want >= %v", got, tt.n*15)
				}
				if val >= (tt.n+1)*15 {
					t.Errorf("NewColumn() = %v, want < %v", got, (tt.n+1)*15)
				}
				if duplicateChecker[val] {
					t.Errorf("NewColumn() = %v, duplicate: %v", got, val)
				}

				duplicateChecker[val] = true
			}
		})
	}
}

func TestColumn_FindAndHit(t *testing.T) {
	tests := []struct {
		name         string
		col          Column
		n            int
		want         string
		wantedResult bool
	}{
		{
			name: "0",
			col: Column{
				NewNum(10),
				NewNum(7),
				NewNum(13),
				NewNum(9),
				NewNum(2),
			},
			n:            9,
			want:         "  10  07  13 *09  02",
			wantedResult: true,
		},
		{
			name: "0",
			col: Column{
				NewNum(10),
				NewNum(7),
				NewNum(13),
				NewNum(9),
				NewNum(2),
			},
			n:            6,
			want:         "  10  07  13  09  02",
			wantedResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.col.FindAndHit(tt.n)
			if tt.col.String() != tt.want {
				t.Errorf("FindAndHit() -> %s, want: %s", tt.col.String(), tt.want)
			}
			if result != tt.wantedResult {
				t.Errorf("FindAndHit() = %v, want: %v", result, tt.wantedResult)
			}
		})
	}
}

func TestColumn_Bingo(t *testing.T) {
	tests := []struct {
		name string
		col  Column
		want bool
	}{
		{
			name: "not hit 1",
			col: Column{
				{hit: true}, {hit: true}, {hit: true}, {hit: false}, {hit: true},
			},
			want: false,
		},
		{
			name: "not hit 2",
			col: Column{
				{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: false},
			},
			want: false,
		},
		{
			name: "not hit 3",
			col: Column{
				{hit: false}, {hit: true}, {hit: true}, {hit: true}, {hit: true},
			},
			want: false,
		},
		{
			name: "hit",
			col: Column{
				{hit: true}, {hit: true}, {hit: true}, {hit: true}, {hit: true},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.col.Bingo(); got != tt.want {
				t.Errorf("Bingo() = %v, want %v", got, tt.want)
			}
		})
	}
}
