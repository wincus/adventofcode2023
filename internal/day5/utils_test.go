package day5

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
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			p:    common.Part1,
			want: 35,
		},
		{
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			p:    common.Part2,
			want: 46,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
