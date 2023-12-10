package main

import (
	"log"

	"github.com/wincus/adventofcode2023/internal/common"
	"github.com/wincus/adventofcode2023/internal/day7"
)

func main() {

	d, err := common.GetData(7)

	if err != nil {
		log.Panicf("no data, no game ... sorry!")
	}

	for _, p := range []common.Part{common.Part1, common.Part2} {
		log.Printf("Solution for Part %v: %v", p, day7.Solve(d, p))
	}
}
