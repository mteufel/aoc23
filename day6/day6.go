package day6

import (
	"fmt"
)

func Task61() {
	fmt.Println("\nDay 6, Part 1: Wait For It")
	fmt.Println("=====================================================")

	//race, distance := Load("sample.txt")
	race, distance := Load("input.txt")
	fmt.Println("Race----->", race)
	fmt.Println("Distance->", distance)

	result := 1
	for idx, race := range race {
		options, ways := AnalyzeRace(race, distance[idx])
		fmt.Println(options, ways)
		result = result * ways
	}
	fmt.Println("Result--> ", result)

}

func Task62() {
	fmt.Println("\nDay 6, Part 2: Wait For It")
	fmt.Println("=====================================================")

	//race, distance := Load("sample2.txt")
	race, distance := Load("input2.txt")
	fmt.Println("Race----->", race)
	fmt.Println("Distance->", distance)

	result := 1
	for idx, race := range race {
		_, ways := AnalyzeRace(race, distance[idx])
		result = result * ways
	}
	fmt.Println("Result--> ", result)

}
