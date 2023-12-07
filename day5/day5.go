package day5

import (
	"fmt"
	"sort"
)

func Task51() {
	fmt.Println("\nDay 5, Part 1: If You Give A Seed A Fertilizer")
	fmt.Println("=====================================================")

	fertilizer := Load("sample.txt")
	//fertilizer := Load("input.txt")
	//fmt.Println("--->", fertilizer)

	locs := make([]int, 0)
	for _, seed := range fertilizer.Seeds {
		fmt.Println("processing seed:", seed)
		locs = append(locs, fertilizer.GetLocationBySeed(seed))
	}
	sort.Ints(locs)
	lowestLocation := locs[0]
	fmt.Println("--->", lowestLocation)
}

func Task52() {
	fmt.Println("\nDay 5, Part 2: If You Give A Seed A Fertilizer")
	fmt.Println("=====================================================")

	// fertilizer := Load("sample.txt")
	fertilizer := Load("input.txt")

	locs := make([]int, 0)

	for i := 0; i <= len(fertilizer.Seeds)-1; i += 2 {
		fmt.Println("--> New range: start,range ", fertilizer.Seeds[i], fertilizer.Seeds[i+1])
		seedRange := fertilizer.CalculateSeedsForRange(fertilizer.Seeds[i], fertilizer.Seeds[i+1])

		for _, seed := range seedRange {
			loc := fertilizer.GetLocationBySeed(seed - 1)
			locs = append(locs, loc)
		}

	}

	sort.Ints(locs)
	lowestLocation := locs[0]
	fmt.Println("--->", lowestLocation)
}
