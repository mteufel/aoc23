package day1

import (
	"aoc23/util"
	"fmt"
	"regexp"
	"strconv"
)

var nonAlphanumericRegex = regexp.MustCompile(`\d*`)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Task11() {
	fmt.Println("\nDay 1, Part 1: Trebuchet?!")
	fmt.Println("===============================================")
	//lines := util.ReadFile("/day1/sample.txt.txt")
	lines := util.ReadFile("/day1/input.txt")
	total := 0
	value := 0
	for _, l := range lines {
		filtered := nonAlphanumericRegex.FindAllString(l, -1)
		value = convert(filtered)
		fmt.Println("", l, value)
		total += value
	}
	fmt.Println("\nResult=", total)
}

func Task12() {
	fmt.Println("\nDay 1, Part 2: Trebuchet?!")
	fmt.Println("===============================================")
	lines := util.ReadFile("/day1/input.txt")
	total := 0
	for _, line := range lines {
		digits := make([]string, 0)
		for i := 0; i < len(line); i++ {
			if line[i] >= '1' && line[i] <= '9' {
				digits = append(digits, string(line[i]))
				continue
			}

			for k, v := range digitMap {
				if i+len(k) > len(line) {
					continue
				}
				if line[i:i+len(k)] != k {
					continue
				}
				digits = append(digits, v)
			}
		}
		comb := digits[0] + digits[len(digits)-1]
		value, _ := strconv.Atoi(comb)
		total += value
		fmt.Println(line, value, total)
	}
	fmt.Println("Result=", total)

}

func convert(src []string) int {
	onlyDigits := make([]string, 0)
	for _, s := range src {
		if s != "" {
			onlyDigits = append(onlyDigits, s)
		}
	}
	complete := ""
	for _, s := range onlyDigits {
		complete += s
	}
	complete = complete[:1] + complete[len(complete)-1:]
	num, _ := strconv.Atoi(complete)
	return num
}
