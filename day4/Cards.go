package day4

import (
	"aoc23/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
}

func Load(file string) []Card {

	lines := util.ReadFile("/day4/" + file)
	cards := make([]Card, 0)

	for idx, line := range lines {

		left := strings.Split(line, "|")[0]
		right := strings.Split(line, "|")[1]

		cards = append(cards, Card{Id: idx + 1, WinningNumbers: toArray(strings.Split(left, ":")[1]), Numbers: toArray(right)})

	}

	return cards
}

func InitializeCardStack(cards []Card) map[int]int {
	stack := make(map[int]int)
	for _, card := range cards {
		stack[card.Id] = 1
	}
	return stack
}

func CheckCard(card Card) int {
	winningNumbers := card.WinningNumbers
	myNumbers := card.Numbers
	multipliers := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	count := -1
	result := 0
	for _, number := range myNumbers {
		if slices.Contains(winningNumbers, number) {
			count++
		}
	}

	if count > -1 {
		result = multipliers[count]
	}
	fmt.Println("Card, result -->", card, result)
	return result
}

func NumberOfMatches(card Card) int {
	count := 0
	for _, number := range card.Numbers {
		if slices.Contains(card.WinningNumbers, number) {
			count++
		}
	}
	return count
}

func toArray(data string) []int {
	pieces := strings.Split(data, " ")
	result := make([]int, 0)
	for _, n := range pieces {
		n = strings.TrimSpace(n)
		if n != "" {
			num, err := strconv.Atoi(n)
			util.CheckError(err)
			result = append(result, num)
		}

	}
	return result
}
