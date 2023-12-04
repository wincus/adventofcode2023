package day4

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2023/internal/common"
)

type card struct {
	id      int
	count   int
	winning map[int]bool
	list    map[int]bool
}

// Solve returns the solutions for day 4
func Solve(s []string, p common.Part) int {

	var sum int

	var cards []card

	for _, line := range s {

		if line == "" {
			continue
		}

		c, err := getCard(line)

		if err != nil {
			log.Println(err)
			continue
		}

		cards = append(cards, c)
	}

	if p == common.Part1 {
		sum = getCardScore(cards)
	}

	if p == common.Part2 {
		sum = countCards(cards)
	}

	return sum
}

// getCard gets a line like:
//
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// and returns a card struct
func getCard(s string) (card, error) {

	var c card

	// remove all whitespacing used for tabulation
	s = curate(s)

	p := strings.Split(s, ":")

	q := strings.Split(p[0], " ")

	if len(q) != 2 {
		return c, fmt.Errorf("invalid card string: %v", s)
	}

	id, err := strconv.Atoi(q[1])

	if err != nil {
		return c, fmt.Errorf("invalid card string: %v", s)
	}

	c.id = id

	r := strings.Split(p[1], "|")

	if len(r) != 2 {
		return c, fmt.Errorf("invalid card string: %v", s)
	}

	var w = make(map[int]bool)

	for _, x := range strings.Split(r[0], " ") {

		if x == "" {
			continue
		}

		y, err := strconv.Atoi(x)

		if err != nil {
			return c, fmt.Errorf("invalid card string: %v", s)
		}

		w[y] = true

	}

	c.winning = w

	var l = make(map[int]bool)

	for _, x := range strings.Split(r[1], " ") {

		if x == "" {
			continue
		}

		y, err := strconv.Atoi(x)

		if err != nil {
			return c, fmt.Errorf("invalid card string: %v", s)
		}

		l[y] = true

	}

	c.list = l
	c.count = 1

	return c, nil
}

func getCardScore(cards []card) int {

	var sum int

	for _, c := range cards {
		sum += calculateScore(c.countWinning())
	}

	return sum
}

func calculateScore(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	return calculateScore(n-1) * 2
}

func (c card) countWinning() int {

	var score int

	for k, v := range c.winning {
		for x, y := range c.list {
			if v && y && k == x {
				score++
			}
		}
	}

	return score

}

// curate removes all whitespacing used for tabulation
// to facilitate the parsing of the input
func curate(s string) string {

	var out strings.Builder
	var count int

	for _, x := range s {
		if x == ' ' && count == 0 {
			count++
			out.WriteRune(x)
		} else if x == ' ' && count > 0 {
			count++
			continue // ignore extra whitespacing
		} else {
			count = 0
			out.WriteRune(x)
		}
	}

	return out.String()
}

// countCards counts the number of cards using the rules
// for part 2
func countCards(cards []card) int {

	for i := 0; i < len(cards); i++ {

		n := cards[i].countWinning()

		if n > 0 {
			for j := 1; j <= n; j++ {

				if i+j >= len(cards) {
					break
				}

				cards[i+j].count += cards[i].count
			}
		}
	}

	var count int

	for _, c := range cards {
		count += c.count
	}

	return count
}
