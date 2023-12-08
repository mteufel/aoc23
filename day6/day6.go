package day6

import (
	"fmt"
)

func Task61() {
	fmt.Println("\nDay 6, Part 1: Wait For It")
	fmt.Println("=====================================================")

	race, distance := Load("sample.txt")
	fmt.Println("Race----->", race)
	fmt.Println("Distance->", distance)

	result := AnalyzeRace(7, 9)
	fmt.Println(result)
}
