package main

import (
	"flag"
	"log"

	"github.com/wincus/adventofcode2023/internal/common"
)

func main() {

	day := flag.String("day", "", "day number")
	flag.Parse()

	if *day == "" {
		log.Panicf("day number is required")
	}

	if err := common.Generate(*day); err != nil {
		log.Panic(err)
	}
}
