package day8

import (
	"fmt"
	"log"

	"strings"

	"github.com/wincus/adventofcode2023/internal/common"
)

const (
	START          = "AAA"
	MAX_ITERATIONS = 1 << 15
)

type node struct {
	name  string
	left  *node
	right *node
}

type instruction struct {
	pos        int
	directions []DIRECTION
}

type DIRECTION int

const (
	LEFT DIRECTION = iota
	RIGHT
)

// Solve returns the solutions for day 8
func Solve(s []string, p common.Part) int {

	// get instructions
	i, err := getInstructions(s[0])

	if err != nil {
		log.Printf("could not parse instructions: %v", err)
		return 0
	}

	// get nodes
	singleStart, multiStart, err := buildGraph(s[1:])

	if err != nil {
		log.Printf("could not parse nodes: %v", err)
		return 0
	}

	if p == common.Part1 {
		return travelSinglePath(singleStart, i)

	}

	if p == common.Part2 {

		var counts []int

		for _, start := range multiStart {
			counts = append(counts, travelSinglePath(start, i))
		}

		// return the least common multiple of all counts
		// the alternative ( traversing the graph for all nodes in
		// parallel would have taken 3 days to compute according to
		// my estimations )
		return LCM(counts)

	}

	return 0
}

func (i *instruction) getNextDirection() (DIRECTION, error) {

	defer func() {
		i.pos++
	}()

	if len(i.directions) == 0 {
		return 0, fmt.Errorf("no directions")
	}

	if i.pos >= len(i.directions) {
		i.pos = 0
	}

	return i.directions[i.pos], nil
}

func parseDirection(r rune) (DIRECTION, error) {
	switch r {
	case 'L':
		return LEFT, nil
	case 'R':
		return RIGHT, nil
	default:
		return 0, fmt.Errorf("invalid direction: %v", r)
	}
}

func getInstructions(s string) (instruction, error) {

	var directions []DIRECTION

	for _, r := range s {
		d, err := parseDirection(r)

		if err != nil {
			return instruction{}, err
		}

		directions = append(directions, d)
	}

	return instruction{directions: directions}, nil
}

func buildGraph(s []string) (*node, []*node, error) {

	var t = make(map[string]*node)

	var singleStart *node
	var multiStart = make([]*node, 0)

	for _, line := range s {

		if len(line) != 16 {
			continue
		}

		var top, left, right string

		// 0123456789012345
		// AAA = (BBB, CCC)
		top = line[0:3]
		left = line[7:10]
		right = line[12:15]

		// register nodes
		if _, ok := t[top]; !ok {
			t[top] = &node{name: top}
		}

		if _, ok := t[left]; !ok {
			t[left] = &node{name: left}
		}

		if _, ok := t[right]; !ok {
			t[right] = &node{name: right}
		}

		// register start nodes for multi path
		if t[top].isStart() {
			multiStart = append(multiStart, t[top])
		}

		// register start node for single path
		if t[top].name == START {
			singleStart = t[top]
		}

		// set left and right
		t[top].left = t[left]
		t[top].right = t[right]
	}

	return singleStart, multiStart, nil
}

func (b *node) GetNode(d DIRECTION) *node {

	switch d {
	case LEFT:
		return b.left
	case RIGHT:
		return b.right
	default:
		return nil
	}

}

func (b *node) isStart() bool {
	return strings.HasSuffix(b.name, "A")
}

func (b *node) isEnd() bool {
	return strings.HasSuffix(b.name, "Z")
}

func travelSinglePath(start *node, i instruction) int {

	var next *node = start
	var count int

	for {

		if next.isEnd() {
			return count
		}

		d, err := i.getNextDirection()

		if err != nil {
			log.Printf("could not get next direction: %v", err)
			return 0
		}

		next = next.GetNode(d)

		count++

		if count > MAX_ITERATIONS {
			log.Printf("max iterations reached")
			return 0
		}
	}
}

// LCM returns the least common multiple of a slice of integers
func LCM(nums []int) int {

	var max int

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	var lcm int = max

	for {
		var found = true

		for _, num := range nums {
			if lcm%num != 0 {
				found = false
				break
			}
		}

		if found {
			return lcm
		}

		lcm += max
	}
}
