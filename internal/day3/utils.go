package day3

import (
	"github.com/wincus/adventofcode2023/internal/common"
)

// partChar is a struct that represents a character found in a part
// includes a list of gears surrounding the character
// is considered valid if it is surrounded by special characters
type partChar struct {
	c     rune
	valid bool
	gears map[gear]bool
}

// part is a struct that represents a part of a number
// includes a list of gears surrounding the part
// is considered valid if at least one of the characters is valid
type part struct {
	value int
	valid bool
	gears map[gear]bool
}

// gear is a struct that represents a gear
// is considered valid if it is surrounded by _exactly_
// two parts
type gear struct {
	row int
	col int
}

// Solve returns the solutions for day 3
func Solve(s []string, p common.Part) int {

	var sum int

	grid := getGrid(s)
	parts := getParts(grid)

	if p == common.Part1 {
		for _, x := range parts {
			if x.valid {
				sum += x.value
			}
		}
	}

	if p == common.Part2 {
		sum = findGearRatioSum(parts)
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

			// if the rune is a number, create a new partChar and append it
			// to the current list of partChars
			if r >= '0' && r <= '9' {

				currentpartChars = append(currentpartChars, partChar{
					c:     r,
					valid: isValid(rowNumber, colNumber, grid),
					gears: getGears(rowNumber, colNumber, grid),
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

	var g = make(map[gear]bool)

	var number int

	for _, pc := range partChars {

		if pc.c >= '0' && pc.c <= '9' {
			number = number*10 + int(pc.c-'0')
		}

		// merge all the gears found in the part
		for k, v := range pc.gears {
			g[k] = v
		}
	}

	p.value = number
	p.valid = isPartValid(partChars)
	p.gears = g

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

// getGears returns a slice of gears found in the grid
// next to a partChar located at row, col
func getGears(row, col int, grid [][]rune) map[gear]bool {

	var gears = make(map[gear]bool)

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

			// append gears
			if grid[row+r][col+c] == '*' {
				gears[gear{row: row + r, col: col + c}] = true
			}
		}
	}

	return gears

}

func findGearRatioSum(parts []part) int {

	var total int

	var allGears = make(map[gear]bool)

	// get uniq gears
	for _, p := range parts {
		for k, g := range p.gears {
			allGears[k] = g
		}
	}

	// for each gear find the parts that include it
	for k := range allGears {
		total += getGearRatio(parts, k)
	}

	return total

}

func getGearRatio(parts []part, g gear) int {

	var found []part

	for _, p := range parts {

		if !p.valid {
			continue
		}

		for k, x := range p.gears {
			if k == g && x {
				found = append(found, p)
			}
		}
	}

	if len(found) == 2 {
		return found[0].value * found[1].value
	}

	return 0

}
