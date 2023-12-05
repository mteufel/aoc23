package day5

import (
	"aoc23/util"
	"strconv"
	"strings"
)

type Fertilizer struct {
	Maps  []Map
	Seeds []int
}

type Combination struct {
	Source      int
	Destination int
	Range       int
}

type Map struct {
	Name         string
	Combinations []Combination
}

func (f *Fertilizer) GetMapByName(name string) Map {
	for _, m := range f.Maps {
		if m.Name == name {
			return m
		}
	}
	return Map{}
}

func (f *Fertilizer) GetLocationBySeed(seed int) int {
	result := 0
	m := f.GetMapByName("soil")
	result = m.GetValue(seed)
	m = f.GetMapByName("fertilizer")
	result = m.GetValue(result)
	m = f.GetMapByName("water")
	result = m.GetValue(result)
	m = f.GetMapByName("light")
	result = m.GetValue(result)
	m = f.GetMapByName("temperature")
	result = m.GetValue(result)
	m = f.GetMapByName("humidity")
	result = m.GetValue(result)
	m = f.GetMapByName("location")
	result = m.GetValue(result)
	return result
}

func (m *Map) GetValue(seed int) int {

	result := seed

	for _, combination := range m.Combinations {
		dest := combination.Destination
		for i := 0; i <= combination.Range-1; i++ {
			if dest == seed {
				return combination.Source + i
			}
			dest = dest + 1
		}
	}
	return result

}

/*
func (m *Map) GetCorrespondings() map[int]int {
	correspondings := make(map[int]int)
	for _, combination := range m.Combinations {
		dest := combination.Destination
		for i := 0; i <= combination.Range-1; i++ {
			correspondings[dest+i] = combination.Source + i
		}
	}
	return correspondings
}

*/

func Load(file string) Fertilizer {
	lines := util.ReadFile("/day5/" + file)

	// --------------------------------------
	// parse the seeds into an int-array
	// --------------------------------------
	seedString := strings.ReplaceAll(lines[0], "seeds: ", "")
	seeds := makeStringToIntArray(strings.TrimSpace(seedString))

	// --------------------------------------------------------
	// we start to read the file at the first map to parse
	// --------------------------------------------------------
	text := ""
	combinations := make([]Combination, 0)
	maps := make([]Map, 0)
	for i := 2; i < len(lines); i++ {
		switch true {
		case text == "":
			// in this case we are at the start of a new map
			text = strings.TrimSpace(strings.ReplaceAll(strings.Split(lines[i], "-")[2], ":", ""))
			text = strings.TrimSpace(strings.ReplaceAll(text, "map", ""))
			//fmt.Println("new map starts-->", text)
		case lines[i] == "" && text != "":
			// in this case we are at the end of one map
			maps = append(maps, Map{Name: text, Combinations: combinations})
			text = ""
			combinations = make([]Combination, 0)

			//fmt.Println("map ends------------------------", lines[i])
			text = ""
		default:
			// we just have to read and parse combinations
			arr := makeStringToIntArray(strings.TrimSpace(lines[i]))
			combinations = append(combinations, Combination{Source: arr[0], Destination: arr[1], Range: arr[2]})
			//fmt.Println("reading data: ", combinations)
		}
		if i == len(lines)-1 {
			maps = append(maps, Map{Name: text, Combinations: combinations})
		}

	}
	fertilizer := Fertilizer{Maps: maps, Seeds: seeds}
	return fertilizer

}

func makeStringToIntArray(s string) []int {
	tempArray := strings.Split(s, " ")
	intArray := make([]int, 0)
	for _, t := range tempArray {
		i, _ := strconv.Atoi(t)
		intArray = append(intArray, i)
	}
	return intArray
}

func MergeMaps(m1 map[int]int, m2 map[int]int) map[int]int {
	merged := make(map[int]int)
	for k, v := range m1 {
		merged[k] = v
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}
