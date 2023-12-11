package day7

import (
	"fmt"
)

func Task71() {
	fmt.Println("\nDay 7, Part 1: Camel Cards")
	fmt.Println("=====================================================")

	//hands := Load("sample.txt", WithoutJoker)
	hands := Load("input.txt", WithoutJoker)

	result := 0
	for idx, h := range hands {
		fmt.Println("Rank, Hand --> ", idx+1, h.Cards, h.GetType(), h.Bid)
		result = result + (h.Bid * (idx + 1))

	}

	print("Result: ", result)

}

func Task72() {
	fmt.Println("\nDay 7, Part 2: Camel Cards")
	fmt.Println("=====================================================")

	//hands := Load("sample2.txt", WithJoker)
	//hands := Load("sample.txt", WithJoker)
	hands := Load("input.txt", WithJoker)

	result := 0
	for idx, h := range hands {
		fmt.Println("Rank, Hand --> ", idx+1, h.Cards, h.GetType(), h.Bid)
		result = result + (h.Bid * (idx + 1))
	}

	print("Result: ", result)
}
