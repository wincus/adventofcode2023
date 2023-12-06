package day6

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2023/internal/common"
)

type race struct {
	time     int
	distance int
}

type races []race

// Solve returns the solutions for day 6
func Solve(s []string, p common.Part) int {

	var r races

	times, distances, err := parse(s)

	if err != nil {
		log.Printf("error parsing input: %v", err)
		return 0
	}

	if p == common.Part1 {
		for x := 0; x < len(times); x++ {

			t, err := strconv.Atoi(times[x])

			if err != nil {
				log.Printf("error converting time to int: %v", err)
				return 0
			}

			d, err := strconv.Atoi(distances[x])

			if err != nil {
				log.Printf("error converting distance to int: %v", err)
				return 0
			}

			r = append(r, race{
				time:     t,
				distance: d,
			})
		}
	}

	if p == common.Part2 {

		var ts strings.Builder
		var ds strings.Builder

		for x := 0; x < len(times); x++ {
			ts.WriteString(times[x])
			ds.WriteString(distances[x])

		}

		t, err := strconv.Atoi(ts.String())

		if err != nil {
			log.Printf("error converting time to int: %v", err)
			return 0
		}

		d, err := strconv.Atoi(ds.String())

		if err != nil {
			log.Printf("error converting distance to int: %v", err)
			return 0
		}

		r = append(r, race{
			time:     t,
			distance: d,
		})
	}

	var mul = 1

	for _, race := range r {
		x, y := getMinMax(race.time, race.distance)

		factor := y - x + 1
		mul *= factor
	}

	return mul

}

// parse returns a slice of times and distances __strings__
func parse(s []string) ([]string, []string, error) {

	var times []string
	var distances []string

	for _, line := range s {

		x := strings.Fields(line)

		if strings.Contains(line, "Time") {
			times = x[1:]
		}

		if strings.Contains(line, "Distance") {
			distances = x[1:]
		}
	}

	return times, distances, nil

}

func getMinMax(tmax, dmax int) (int, int) {

	// t = -b +- sqrt(b^2 - 4ac) / 2a
	t := float64(tmax)
	d := float64(dmax)

	disc := math.Pow(t, 2) - 4*d

	x0 := (-1*t + math.Sqrt(disc)) / -2
	x1 := (-1*t - math.Sqrt(disc)) / -2

	i0 := int(math.Ceil(x0))
	i1 := int(math.Floor(x1))

	// correct in case the root matches exactly ( !! )
	if i0*(tmax-i0) == dmax {
		i0++
	}

	if i1*(tmax-i1) == dmax {
		i1--
	}

	return i0, i1

}
