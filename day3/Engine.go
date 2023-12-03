package day3

import (
	"aoc23/util"
	"strconv"
	"strings"
)

type Engine struct {
	EngineData     [][]string
	Parts          []Part
	PartsWithGears []Part
}

type Gear struct {
	X int
	Y int
}

type Part struct {
	X      int
	Y      int
	Id     string
	Value  int
	IsPart bool
	Gear   Gear
}

func newPart(x int, y int, id string, isPart bool) Part {
	return Part{X: x, Y: y, Id: id, Value: 0, IsPart: isPart}
}

func newEngine(engineData [][]string, parts []Part) Engine {
	return Engine{EngineData: engineData, Parts: parts}
}

func load(file string) Engine {

	var x, y int
	engineData := make([][]string, 0)
	parts := make([]Part, 0)
	lines := util.ReadFile("/day3/" + file)

	//                Y
	//         0      1       2
	//      ---------------------
	//   0 |[0][0]  [0][1]  [0][2]
	//X  1 |[1][0]  [1][1]  [1][2]
	//   2 |[2][0]  [2][1]  [2][2]

	x = 0
	for _, line := range lines {

		chars := strings.Split(line, "")
		temp := make([]string, 0)
		fragment := ""
		y = 0
		for _, char := range chars {
			temp = append(temp, char)
			if util.IsInt(char) {
				fragment += char
			} else {
				if strings.Compare(fragment, "") != 0 {

					part := newPart(x, y-len(fragment), fragment, false)
					parts = append(parts, part)
					fragment = ""
				}

			}
			y++
		}
		if strings.Compare(fragment, "") != 0 {
			part := newPart(x, (y)-len(fragment), fragment, false)
			parts = append(parts, part)
			fragment = ""
		}
		engineData = append(engineData, temp)
		x++
	}

	engine := newEngine(engineData, parts)

	// check for parts
	checkedParts := make([]Part, 0)
	for idx, p := range parts {

		p.IsPart, p.Gear = isPart(engine, p)
		p.Value, _ = strconv.Atoi(p.Id)
		p.Id = strconv.Itoa(idx)
		checkedParts = append(checkedParts, p)
	}
	engine.Parts = checkedParts

	return engine
}

func isPart(engine Engine, partToCheck Part) (bool, Gear) {

	startX := partToCheck.X
	startY := partToCheck.Y
	value := ""
	valid := true
	gear := Gear{X: -1, Y: -1}

	// check left
	if (startY - 1) >= 0 {
		value = engine.EngineData[startX][startY-1]
		if util.IsInt(value) || strings.Compare(value, ".") != 0 {
			valid = false
		}
		if isGear(value) {
			gear = Gear{X: startX, Y: startY - 1}
		}
	}
	// check right
	if ((startY + 1) + len(partToCheck.Id)) < len(engine.EngineData[0]) {
		value = engine.EngineData[startX][startY+len(partToCheck.Id)]
		if util.IsInt(value) || strings.Compare(value, ".") != 0 {
			valid = false
		}
		if isGear(value) {
			gear = Gear{X: startX, Y: startY + len(partToCheck.Id)}
		}
	}
	// check top
	if (startX - 1) >= 0 {
		s := startY - 1
		if s < 0 {
			s = 0
		}
		for i := s; i < startY+len(partToCheck.Id)+1; i++ {
			if i < len(engine.EngineData[0]) {
				value = engine.EngineData[startX-1][i]
				if util.IsInt(value) || strings.Compare(value, ".") != 0 {
					valid = false
				}
				if isGear(value) {
					gear = Gear{X: startX - 1, Y: i}
				}
			}

		}
	}

	// check bottom
	if (startX + 1) < len(engine.EngineData) {
		s := startY - 1
		if s < 0 {
			s = 0
		}
		for i := s; i < startY+len(partToCheck.Id)+1; i++ {
			if i < len(engine.EngineData[0]) {
				value = engine.EngineData[startX+1][i]
				if util.IsInt(value) || strings.Compare(value, ".") != 0 {
					valid = false
				}
				if isGear(value) {
					gear = Gear{X: startX + 1, Y: i}
				}
			}

		}
	}

	return !valid, gear
}

func AddPartsWithGears(parts []Part) []Part {
	withGearsOnly := make([]Part, 0)
	for _, p := range parts {
		if p.Gear.X != -1 && p.Gear.Y != -1 {
			withGearsOnly = append(withGearsOnly, p)
		}
	}
	return withGearsOnly
}

func isGear(s string) bool {
	if strings.Compare(s, "*") == 0 {
		return true
	}
	return false
}

func removePart(parts []Part, part Part) []Part {
	var new []Part
	for _, v := range parts {
		if v.Id != part.Id {
			new = append(new, v)
		}
	}
	return new
}

func CalulateGears(parts []Part) ([]Part, int) {
	part := parts[0]
	parts = removePart(parts, part)
	var otherPart Part
	var resultParts []Part
	result := -1
	for _, p := range parts {
		if p.Id != part.Id && p.Gear.X == part.Gear.X && p.Gear.Y == part.Gear.Y {
			otherPart = p
			break
		}
	}
	if otherPart != (Part{}) {
		result = part.Value * otherPart.Value
		resultParts = removePart(parts, otherPart)
		resultParts = removePart(resultParts, part)
	} else {
		resultParts = removePart(parts, part)
	}

	return resultParts, result
}
