package day6

import (
	"aoc23/util"
	"regexp"
	"strconv"
)

var nonAlphanumericRegex = regexp.MustCompile(`\d*`)

func Load(file string) ([]int, []int) {
	lines := util.ReadFile("/day6/" + file)

	race := make([]int, 0)
	distance := make([]int, 0)

	for idx, line := range lines {
		if idx == 0 {
			race = parseLine(line)
		}
		if idx == 1 {
			distance = parseLine(line)
		}
	}
	return race, distance
}

func parseLine(line string) []int {
	filtered := nonAlphanumericRegex.FindAllString(line, -1)
	onlyDigits := make([]int, 0)
	for _, s := range filtered {
		if s != "" {
			n, _ := strconv.Atoi(s)
			onlyDigits = append(onlyDigits, n)
		}
	}
	return onlyDigits
}

func AnalyzeRace(race int, distance int) []int {
	options := make([]int, 0)
	for i := 1; i <= race; i++ {
		millimeters := (race - i) * i
		if millimeters > 0 {
			options = append(options, millimeters)
		}
	}
	return options
}
