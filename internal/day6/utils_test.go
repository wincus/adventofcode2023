package day6

import (
	"testing"

	"github.com/wincus/adventofcode2023/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  int
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			p:    common.Part1,
			want: 288,
		},
		{
			input: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			p:    common.Part2,
			want: 71503,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
