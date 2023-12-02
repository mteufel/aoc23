package day2

import (
	"aoc23/util"
	"fmt"
)

func Task21() {
	fmt.Println("\nDay 2, Part 1: Cube Conundrum")
	fmt.Println("===============================================")

	//lines := util.ReadFile("/day2/sample.txt")
	//maxRed := 12
	//maxGreen := 13
	//maxBlue := 14

	lines := util.ReadFile("/day2/input.txt")
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	result := 0
	for _, l := range lines {
		game := ParseGame(l, maxRed, maxGreen, maxBlue)
		game.ToString()
		if game.Valid {
			result += game.Id
		}

	}

	fmt.Println("\nResult=", result)

}

func Task22() {
	fmt.Println("\nDay 2, Part  2: Cube Conundrum")
	fmt.Println("===============================================")

	//lines := util.ReadFile("/day2/sample.txt")
	//maxRed := 12
	//maxGreen := 13
	//maxBlue := 14

	lines := util.ReadFile("/day2/input.txt")
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	result := 0
	for _, l := range lines {
		game := ParseGame(l, maxRed, maxGreen, maxBlue)
		game.ToString()
		result += game.Power

	}

	fmt.Println("\nResult=", result)

}
