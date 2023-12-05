package day5

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/wincus/adventofcode2023/internal/common"
)

type seeds []int

type seedRangeList []seedRange

type seedRange struct {
	start int
	end   int
}

type mapper map[mapperKey]router

type mapperKey string

type router []mapperItem

type mapperItem struct {
	sourceStart      int
	destinationStart int
	length           int
}

var order = []mapperKey{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

// Solve returns the solutions for day 5
func Solve(s []string, p common.Part) int {

	var m mapper
	var err error
	var r seeds
	var t seedRangeList

	if p == common.Part1 {

		r, err = getSeeds(s[0])

		if err != nil {
			log.Println(err)
			return 0
		}

		for _, seed := range r {
			// convert seed to seedRanges to simplify the logic
			t = append(t, seedRange{
				start: seed,
				end:   seed + 1,
			})
		}
	}

	if p == common.Part2 {

		t, err = getSeedRanges(s[0])

		if err != nil {
			log.Println(err)
			return 0
		}
	}

	// get mapping rules
	m, err = getMappers(s[1:])

	if err != nil {
		log.Println(err)
		return 0
	}

	//var output int
	var min int = math.MaxInt

	startTime := time.Now()

	for _, seedRange := range t {
		for seed := seedRange.start; seed < seedRange.end; seed++ {

			input := seed

			for _, key := range order {
				input = m[key].remap(input)
			}

			if input < min {
				min = input
			}
		}
	}

	log.Printf("time to finish part %v: %v", p, time.Since(startTime))

	return min

}

func getMappers(s []string) (mapper, error) {

	var m mapper = make(map[mapperKey]router)

	i := getMapperBlocks(s)

	for _, block := range i {

		// get mapper
		if strings.Contains(strings.Join(block, " "), "map") {
			mk, mi, err := getMapper(block)

			if err != nil {
				return m, err
			}

			m[mk] = mi

			continue
		}

		log.Printf("unknown block found ( will be ignored ): %s", block)

	}

	return m, nil
}

func getMapperBlocks(s []string) [][]string {

	var block []string
	var blocks [][]string

	for _, line := range s {

		if line == "" {

			if len(block) > 0 {
				blocks = append(blocks, block)
				block = []string{}
			}

			continue
		}

		block = append(block, strings.TrimSpace(line))

	}

	if len(block) > 0 {
		blocks = append(blocks, block)
	}

	return blocks
}

func getSeedRanges(s string) (seedRangeList, error) {

	var e seedRangeList

	if !strings.Contains(s, "seeds") {
		return e, errors.New("line does not contain seeds")
	}

	s = strings.Replace(s, "seeds:", "", -1)

	s = strings.TrimSpace(s)

	a := strings.Split(s, " ")

	for i := 0; i < len(a); i += 2 {

		start, err := strconv.Atoi(a[i])

		if err != nil {
			return e, fmt.Errorf("invalid seed range string %v: %v", s, err)
		}

		end, err := strconv.Atoi(a[i+1])

		if err != nil {
			return e, fmt.Errorf("invalid seed range string %v: %v", s, err)
		}

		e = append(e, seedRange{
			start: start,
			end:   start + end,
		})
	}

	return e, nil

}

func getSeeds(s string) (seeds, error) {

	var e seeds

	s = strings.TrimSpace(s)

	if !strings.Contains(s, "seeds") {
		return e, errors.New("line does not contain seeds")
	}

	s = strings.Replace(s, "seeds:", "", -1)

	a := strings.Split(s, " ")

	for _, v := range a {

		if v == "" {
			continue
		}

		i, err := strconv.Atoi(v)

		if err != nil {
			return e, fmt.Errorf("invalid seeds string %v: %v", s, err)
		}

		e = append(e, i)

	}

	return seeds(e), nil

}

func getMapper(s []string) (mapperKey, router, error) {

	var r mapperKey
	var i []mapperItem

	for _, line := range s {

		if line == "" {
			continue
		}

		// get mapper key
		if strings.Contains(line, "map") {

			a := strings.Split(line, " ")

			if len(a) != 2 {
				return r, i, errors.New("invalid mapper line")
			}

			r = mapperKey(a[0])

			continue
		}

		// get mapper items
		a := strings.Split(line, " ")

		if len(a) != 3 {
			return r, i, errors.New("invalid mapper line")
		}

		destinationStart, err := strconv.Atoi(a[0])

		if err != nil {
			return r, i, err
		}

		sourceStart, err := strconv.Atoi(a[1])

		if err != nil {
			return r, i, err
		}

		Items, err := strconv.Atoi(a[2])

		if err != nil {
			return r, i, err
		}

		i = append(i, mapperItem{
			sourceStart:      sourceStart,
			destinationStart: destinationStart,
			length:           Items,
		})
	}

	return r, i, nil

}

func (r router) remap(input int) int {

	for _, item := range r {
		if input < item.sourceStart+item.length && input >= item.sourceStart {
			return input - item.sourceStart + item.destinationStart
		}
	}

	return input
}
