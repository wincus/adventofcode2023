package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2023/internal/common"
)

type Game struct {
	id    int
	plays []Play
}

type Play struct {
	blueCubes  int
	redCubes   int
	greenCubes int
}

type Color int

const (
	Blue Color = iota
	Red
	Green
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

// Solve returns the solutions for day 2
func Solve(s []string, p common.Part) int {

	var sum int

	for _, line := range s {

		if line == "" {
			continue
		}

		game, err := getGame(line)

		if err != nil {
			log.Printf("could not parse line %v: %v", line, err)
			continue
		}

		if p == common.Part1 {

			if game.isValid() {
				sum += game.id
			}
		}

		if p == common.Part2 {
			sum += game.getPower()
		}
	}

	return sum
}

// getGame gets a string like:
//
//	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
//
// and returns a Game struct

func getGame(s string) (Game, error) {

	var g Game
	var err error

	h := strings.Split(s, ":")
	if len(h) != 2 {
		return g, fmt.Errorf("could not parse game: %v", s)
	}

	i := strings.Split(h[0], " ")
	if len(i) != 2 {
		return g, fmt.Errorf("could not parse game id: %v", h[0])
	}

	// game id
	g.id, err = strconv.Atoi(i[1])

	if err != nil {
		return g, fmt.Errorf("could not parse game id: %v", i[1])
	}

	g.addPlays(h[1])

	return g, nil

}

// getPower returns the power of the game as described in Part 2
func (g *Game) getPower() int {

	var p Play

	for _, play := range g.plays {

		if play.redCubes > p.redCubes {
			p.redCubes = play.redCubes
		}

		if play.greenCubes > p.greenCubes {
			p.greenCubes = play.greenCubes
		}

		if play.blueCubes > p.blueCubes {
			p.blueCubes = play.blueCubes
		}
	}

	return p.redCubes * p.greenCubes * p.blueCubes
}

// isValid returns true if the game is valid
func (g *Game) isValid() bool {

	for _, play := range g.plays {
		if play.redCubes > MAX_RED || play.greenCubes > MAX_GREEN || play.blueCubes > MAX_BLUE {
			return false
		}
	}

	return true
}

// addPlays gets a string like:
//
//	"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
//
// and adds the play to the game
func (g *Game) addPlays(s string) {

	h := strings.Split(s, ";")

	for _, play := range h {

		var p Play

		k := strings.Split(play, ",")

		for _, l := range k {
			p.addCube(l)
		}

		g.plays = append(g.plays, p)

	}

}

// addCube gets a string like "3 blue" and
// adds the cube to the play
func (p *Play) addCube(s string) {

	// ensure s is stripped of whitespace
	s = strings.TrimSpace(s)

	m := strings.Split(s, " ")

	if len(m) != 2 {
		log.Printf("could not parse cube: %v", s)
		return
	}

	// cube color
	c, err := getColor(m[1])

	if err != nil {
		log.Printf("could not parse cube color: %v", m[1])
		return
	}

	// cube count
	n, err := strconv.Atoi(m[0])

	if err != nil {
		log.Printf("could not parse cube count: %v", m[0])
		return
	}

	switch c {
	case Blue:
		p.blueCubes += n
	case Red:
		p.redCubes += n
	case Green:
		p.greenCubes += n
	default:
		log.Printf("could not add cube: %v", c)
	}
}

// getColor returns the color of the cube
func getColor(s string) (Color, error) {

	switch s {
	case "blue":
		return Blue, nil
	case "red":
		return Red, nil
	case "green":
		return Green, nil
	default:
		return 0, fmt.Errorf("could not parse color: %v", s)
	}
}
