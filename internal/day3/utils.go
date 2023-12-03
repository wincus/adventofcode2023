package day3

import (
	"github.com/wincus/adventofcode2023/internal/common"
)

// partChar is a struct that represents a character found in a part
// is considered valid if it is surrounded by special characters
type partChar struct {
	c     rune
	valid bool
	// g     gear
}

// part is a struct that represents a part of a number
// is considered valid if at least one of the characters is valid
type part struct {
	value int
	valid bool
	// gears []gear
}

// // gear is a struct that represents a gear
// // is considered valid if it is surrounded by _exactly_
// // two parts
// type gear struct {
// 	row int
// 	col int
// }

// Solve returns the solutions for day 3
func Solve(s []string, p common.Part) int {

	var sum int

	grid := getGrid(s)

	if p == common.Part1 {
		for _, x := range getParts(grid) {
			if x.valid {
				sum += x.value
			}
		}
	}

	if p == common.Part2 {
		for _, x := range getGears(grid) {
			sum += x.value
		}
	}

	return sum

}

// getGrid returns a 2D array of runes from a slice of strings
func getGrid(s []string) [][]rune {

	var grid [][]rune

	// parse input string into a 2D array of runes
	for _, line := range s {

		if line == "" {
			continue
		}

		var row []rune

		for _, r := range line {
			row = append(row, r)
		}

		grid = append(grid, row)

	}

	return grid

}

// getParts returns a slice of parts found in the grid
func getParts(grid [][]rune) []part {

	var numbers []part

	var currentpartChars []partChar

	// for each row in the grid
	for rowNumber, row := range grid {

		// for each rune in the row
		for colNumber, r := range row {

			// if the rune is a number, add it to the list of numbers
			if r >= '0' && r <= '9' {

				currentpartChars = append(currentpartChars, partChar{
					c:     r,
					valid: isValid(rowNumber, colNumber, grid),
				})

			} else {

				// if the rune is not a number, check if we have a valid number
				if len(currentpartChars) > 0 {
					numbers = append(numbers, getPart(currentpartChars))
				}

				// reset the list of partChars
				currentpartChars = []partChar{}
			}
		}
	}

	return numbers

}

// getPart returns a part from a slice of partChars
func getPart(partChars []partChar) part {

	var p part

	var number int

	for _, pc := range partChars {

		if pc.c >= '0' && pc.c <= '9' {
			number = number*10 + int(pc.c-'0')
		}
	}

	p.value = number
	p.valid = isPartValid(partChars)

	return p

}

// isPartValid iterates over all the chars found in a part and returns true
// if at least one of them is valid
func isPartValid(partChars []partChar) bool {

	for _, pc := range partChars {
		if pc.valid {
			return true
		}
	}

	return false

}

// isValid returns true if the surrounding characters found are
// special ( that is, not a number or space or a dot )
func isValid(row, col int, grid [][]rune) bool {

	for _, r := range []int{-1, 0, 1} {
		for _, c := range []int{-1, 0, 1} {

			// skip the current position
			if r == 0 && c == 0 {
				continue
			}

			// skip out of bounds
			if row+r < 0 || row+r >= len(grid) || col+c < 0 || col+c >= len(grid[row+r]) {
				continue
			}

			// skip spaces
			if grid[row+r][col+c] == ' ' {
				continue
			}

			// skip dots
			if grid[row+r][col+c] == '.' {
				continue
			}

			// skip numbers
			if grid[row+r][col+c] >= '0' && grid[row+r][col+c] <= '9' {
				continue
			}

			// if we get here, we found a special character
			return true

		}
	}

	return false

}
