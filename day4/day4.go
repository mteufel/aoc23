package day4

import (
	"fmt"
)

func Task41() {
	fmt.Println("\nDay 4, Part 1: Scratchcards")
	fmt.Println("===============================================")

	result := 0
	//cards := Load("sample.txt")
	cards := Load("input.txt")

	for _, card := range cards {
		result += CheckCard(card)
	}

	fmt.Println("Result=", result)
}

func Task42() {
	fmt.Println("\nDay 4, Part 2: Scratchcards")
	fmt.Println("===============================================")

	//result := 0
	//cards := Load("sample.txt")
	cards := Load("input.txt")

	stack := InitializeCardStack(cards)

	for _, card := range cards {
		copies := stack[card.Id]
		matches := NumberOfMatches(card)
		max := card.Id + matches
		for i := card.Id + 1; i <= max; i++ {
			stack[i] = stack[i] + (1 * copies)
		}
	}
	result := 0
	for _, m := range stack {
		result = result + m
	}
	fmt.Println("Result=", result)
}
