package day5

import (
	"fmt"
	"sort"
	"sync"
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

	var wg sync.WaitGroup

	//fertilizer := Load("sample.txt")
	fertilizer := Load("input.txt")

	locs := make([]int, 0)

	for i := 0; i <= len(fertilizer.Seeds)-1; i += 2 {
		fmt.Println("--> New range: start,range ", fertilizer.Seeds[i], fertilizer.Seeds[i+1])
		seedRange := fertilizer.CalculateSeedsForRange(fertilizer.Seeds[i], fertilizer.Seeds[i+1])

		ch := make(chan int)

		for _, seed := range seedRange {
			wg.Add(1)
			go func(fertilizer Fertilizer, seed int, ch chan<- int) {
				defer wg.Done()
				ch <- fertilizer.GetLocationBySeed(seed)
			}(fertilizer, seed, ch)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for i := range ch {
			locs = append(locs, i)
		}
	}

	sort.Ints(locs)
	lowestLocation := locs[0]
	fmt.Println("--->", lowestLocation)
}
