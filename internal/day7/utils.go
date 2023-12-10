package day7

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2023/internal/common"
)

type Card rune

type Hand [5]Card

type Play struct {
	H   Hand
	Bid int
}

type Game struct {
	play []Play
	p    common.Part
}

type HandType int

const (
	HIGH_CARD HandType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

// card values
var VALUES = map[common.Part]string{
	common.Part1: "23456789TJQKA",
	common.Part2: "J23456789TQKA",
}

// Solve returns the solutions for day 7
func Solve(s []string, p common.Part) int {

	g, err := Parse(s, p)

	if err != nil {
		log.Fatal(err)
	}

	var sum int

	sort.Sort(g)

	for i := 0; i < len(g.play); i++ {
		sum += g.play[i].Bid * (len(g.play) - i)
	}

	return sum
}

func Parse(s []string, p common.Part) (Game, error) {

	var g Game
	var play []Play

	for _, x := range s {

		if x == "" {
			continue
		}

		a := strings.Fields(x)

		if len(a) != 2 {
			return g, fmt.Errorf("invalid input: %v", x)
		}

		bid, err := strconv.Atoi(a[1])

		if err != nil {
			return g, fmt.Errorf("invalid bid: %v", a[1])
		}

		h, err := getHand(a[0])

		if err != nil {
			return g, fmt.Errorf("invalid hand: %v", a[0])
		}

		play = append(play, Play{H: h, Bid: bid})
	}

	g.play = play
	g.p = p
	return g, nil

}

func getHand(s string) (Hand, error) {

	if len(s) != 5 {
		return Hand{}, fmt.Errorf("invalid hand count: %v", len(s))
	}

	var h Hand

	for i, c := range s {
		h[i] = Card(c)
	}

	return h, nil

}

func (g Game) Len() int {
	return len(g.play)
}

func (g Game) Swap(i, j int) {
	g.play[i], g.play[j] = g.play[j], g.play[i]
}

// Less returns true if i is less than j
func (g Game) Less(i, j int) bool {

	vi, vj := g.play[i].H.GetHandType(g.p), g.play[j].H.GetHandType(g.p)

	if vi != vj {
		return vi > vj
	}

	for k := 0; k < 5; k++ {
		if g.play[i].H[k] != g.play[j].H[k] {
			return g.play[i].H[k].getValue(g.p) > g.play[j].H[k].getValue(g.p)
		}
	}

	return false
}

func (h Hand) GetHandType(p common.Part) HandType {

	var m = make(map[Card]int)
	var l []int
	var jokers int

	for _, c := range h {
		m[c]++
	}

	// joker is considered a wild card in part 2
	if p == common.Part2 {
		jokers = m[Card('J')]
		delete(m, Card('J'))

		// if there are 5 jokers, it's a five of a kind
		if jokers == 5 {
			return FIVE_OF_A_KIND
		}
	}

	for _, v := range m {
		l = append(l, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(l)))

	// joker is considered a wild card in part 2
	if p == common.Part2 {
		l[0] += jokers
	}

	if l[0] == 5 {
		return FIVE_OF_A_KIND
	}

	if l[0] == 4 {
		return FOUR_OF_A_KIND
	}

	if l[0] == 3 && l[1] == 2 {
		return FULL_HOUSE
	}

	if l[0] == 3 {
		return THREE_OF_A_KIND
	}

	if l[0] == 2 && l[1] == 2 {
		return TWO_PAIR
	}

	if l[0] == 2 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func (c Card) getValue(p common.Part) int {
	v := strings.Index(VALUES[p], string(c))
	return v
}

func (c Card) String() string {
	return string(c)
}
