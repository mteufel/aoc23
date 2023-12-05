package day5

import (
	"fmt"
	"sort"
)

func Task51() {
	fmt.Println("\nDay 5, Part 1: If You Give A Seed A Fertilizer")
	fmt.Println("=====================================================")

	//fertilizer := Load("sample.txt")
	fertilizer := Load("input.txt")
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
