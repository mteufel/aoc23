package day3

import (
	"fmt"
	"strconv"
)

func Task31() {
	fmt.Println("\nDay 3, Part 1: Gear Ratios")
	fmt.Println("===============================================")

	// My solution:
	// (1) load engine data into a 2d array OK
	// (2) detect and store start pos x,y of every potentially part OK
	// (3) scan elements around each potentially part to check if its is really a part OK
	// (4) store the information if it is really a part OK
	// (5) provide a list of parts with following values: X,Y,Id (PartNumber), IsPart (True/False) OK

	//engine := load("sample.txt")
	engine := load("input.txt")
	fmt.Println("Engine loaded: max X=" + strconv.Itoa(len(engine.EngineData)) + " max Y=" + strconv.Itoa(len(engine.EngineData[0])))
	result := 0
	for _, p := range engine.Parts {
		if p.IsPart {
			result += p.Value
		}
		fmt.Println("Processing ", p, result)
	}
	fmt.Println("Result=", result)

}

func Task32() {
	fmt.Println("\nDay 3, Part 2: Gear Ratios")
	fmt.Println("===============================================")

	// My solution:
	// (1) Find all parts that include a gear, remember its position OK
	// (2) Store all gear with its position in a list OK
	// (3) Iterate that list, take the first value check if there is another gear with the same coords OK
	//        yes: multiply both gears OK
	// (4) Remove the actual gear the referenced gear, if there is any OK
	// (5) Repeat until the list is empty OK

	//engine := load("sample.txt")
	engine := load("input.txt")
	fmt.Println("Engine loaded: max X=" + strconv.Itoa(len(engine.EngineData)) + " max Y=" + strconv.Itoa(len(engine.EngineData[0])))
	engine.PartsWithGears = AddPartsWithGears(engine.Parts)

	rest := engine.PartsWithGears
	result := 0
	total := 0
	for len(rest) > 0 {
		rest, result = CalulateGears(rest)
		if result > 0 {
			total += result
		}
	}
	fmt.Println("\nResult=", total)

}
