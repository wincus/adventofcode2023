package day1

import (
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2023/internal/common"
)

// Solve returns the solutions for day 1
func Solve(s []string, p common.Part) int {

	var sum int

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		if p == common.Part2 {
			line = strings.ReplaceAll(line, "zero", "z0o")
			line = strings.ReplaceAll(line, "one", "o1e")
			line = strings.ReplaceAll(line, "two", "t2o")
			line = strings.ReplaceAll(line, "three", "t3e")
			line = strings.ReplaceAll(line, "four", "f4r")
			line = strings.ReplaceAll(line, "five", "f5e")
			line = strings.ReplaceAll(line, "six", "s6x")
			line = strings.ReplaceAll(line, "seven", "s7n")
			line = strings.ReplaceAll(line, "eight", "e8t")
			line = strings.ReplaceAll(line, "nine", "n9e")
		}

		var n []int

		for _, c := range line {
			if i, err := strconv.Atoi(string(c)); err == nil {
				n = append(n, i)
			}
		}

		if len(n) == 1 {
			sum += n[0] * 11
		}

		if len(n) > 1 {
			sum += n[0]*10 + n[len(n)-1]
		}
	}

	return sum
}
