package main

import (
	"fmt"
	"strings"
	"github.com/polysalama/aoc_2023/utils"
)

func main() {
	var sumPart1 int = 0
	var sumPart2 int = 0
	bagContents := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var lines []string = utils.ReadToSlice("./input.txt")
	var games = saveGames(lines)
	for i, game := range games {
		minCubes := make(map[string]int)
		var validGame bool = true
		for _, g := range game {
			for k, v := range g {
				validGame = validGame && (v <= bagContents[k])
				if v > minCubes[k] {
					minCubes[k] = v
				}
			}
		}
		if validGame {
			sumPart1 += (i + 1)
		}
		var gameSum = 1

		for _, v := range minCubes {
			gameSum *= v
		}
		sumPart2 += gameSum
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}

func saveGames(lines []string) [][]map[string]int {
	var games [][]map[string]int
	for _, game := range lines {
		var oneGame []map[string]int
		gameSets := strings.Split(game, ": ")[1]
		for _, set := range strings.Split(gameSets, "; ") {
			setResult := strings.Split(set, ", ")
			var setMap map[string]int = make(map[string]int)
			for _, result := range setResult {
				r := strings.Split(result, " ")
				setMap[r[1]] = utils.StringToInt(r[0])
			}
			oneGame = append(oneGame, setMap)
		}
		games = append(games, oneGame)
	}

	return games
}
