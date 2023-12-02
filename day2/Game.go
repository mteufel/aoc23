package day2

import (
	"aoc23/util"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Original string
	Id       int
	Blue     int
	Red      int
	Green    int
	Power    int
	Valid    bool
}

func (g *Game) ToString() {
	fmt.Println(strconv.Itoa(g.Id) + ": red=" + strconv.Itoa(g.Red) + " green=" + strconv.Itoa(g.Green) + " blue=" + strconv.Itoa(g.Blue) + " power=" + strconv.Itoa(g.Power) + " value=" + strconv.FormatBool(g.Valid) + " | " + g.Original)
}

func ParseGame(line string, maxRed int, maxGreen int, maxBlue int) Game {
	var game Game
	line = strings.TrimSpace(line)
	game.Original = line
	game.Valid = true
	game.Red = 0
	game.Green = 0
	game.Blue = 0

	// parse the string
	var id = strings.Split(line, ":")[0]
	id = strings.ReplaceAll(id, "Game ", "")
	var cubes = strings.Split(line, ":")[1]
	var sets = strings.Split(cubes, ";")

	// sort and calculate the cubes
	for _, cubeSet := range sets {
		cubeSet = strings.TrimSpace(cubeSet)
		cubes := strings.Split(cubeSet, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			value, err := strconv.Atoi(strings.Split(cube, " ")[0])
			util.CheckError(err)
			if strings.Contains(cube, "red") {
				if value > game.Red {
					game.Red = value
				}
				if value > maxRed {
					game.Valid = false
				}
			}
			if strings.Contains(cube, "green") {
				if value > game.Green {
					game.Green = value
				}
				if value > maxGreen {
					game.Valid = false
				}

			}
			if strings.Contains(cube, "blue") {
				if value > game.Blue {
					game.Blue = value
				}
				if value > maxBlue {
					game.Valid = false
				}
			}

		}
	}

	game.Id, _ = strconv.Atoi(id)
	game.Power = game.Red * game.Green * game.Blue

	return game
}

func isGameValid(game Game, red int, green int, blue int) bool {
	if game.Red <= red && game.Green <= green && game.Blue <= blue {
		return true
	}
	return false
}
