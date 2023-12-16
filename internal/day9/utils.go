package day9

import (
	"log"
	"strings"

	"github.com/wincus/adventofcode2023/internal/common"
)

// Solve returns the solutions for day 9
func Solve(s []string, p common.Part) int {

	var count int

	for _, line := range s {

		var input [][]int

		if line == "" {
			continue
		}

		l := strings.Fields(line)

		i, err := common.ToInt(l)

		if err != nil {
			log.Printf("could not convert input to int: %v", err)
			return 0
		}

		input = append(input, i)

		if p == common.Part1 {
			input = fill(input)

			// add a zero to the end of the last line
			input[len(input)-1] = append(input[len(input)-1], 0)

			output := reduceBack(input)

			count += getLastItem(getLastSlice(output))
		}

		if p == common.Part2 {
			input = fill(input)

			// add a zero to the start of the last line
			input[len(input)-1] = append([]int{0}, input[len(input)-1]...)

			output := reduceFront(input)

			count += getLastSlice(output)[0]
		}
	}

	return count
}

func reduceBack(a [][]int) [][]int {

	if len(a) == 1 {
		return a
	}

	var output = make([][]int, len(a)-1)

	output = a[:len(a)-1]
	last := getLastItem(getLastSlice(a))
	output[len(output)-1] = append(output[len(output)-1], last+getLastItem(output[len(output)-1]))

	return reduceBack(output)
}

func reduceFront(a [][]int) [][]int {

	if len(a) == 1 {
		return a
	}

	var output = make([][]int, len(a)-1)

	output = a[:len(a)-1]
	last := getFirstItem(getLastSlice(a))
	output[len(output)-1] = append([]int{output[len(output)-1][0] - last}, output[len(output)-1]...)

	return reduceFront(output)
}

func fill(a [][]int) [][]int {

	previous := a[len(a)-1]

	if isZero(previous) {
		return a
	}

	z := make([]int, len(previous)-1)

	for i := 0; i < len(previous)-1; i++ {
		z[i] = previous[i+1] - previous[i]
	}

	a = append(a, z)

	return fill(a)
}

func getLastSlice(a [][]int) []int {
	return a[len(a)-1]
}

func getLastItem(a []int) int {
	return a[len(a)-1]
}

func getFirstItem(a []int) int {
	return a[0]
}

func isZero(a []int) bool {

	for _, v := range a {
		if v != 0 {
			return false
		}
	}

	return true
}
