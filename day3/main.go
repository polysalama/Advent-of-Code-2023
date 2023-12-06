package main

import (
	"fmt"
	"regexp"
	"github.com/polysalama/aoc_2023/utils"
)

type Symbol struct {
	x, y, height, width int
	value               string
	adjacent            []Part
}

type Part struct {
	x, y, height, width, value int
}

func main() {
	var reSymbol = regexp.MustCompile(`[^\w.]`)
	var reNumber = regexp.MustCompile(`\d+`)
	symbols := []Symbol{}
	parts := []Part{}

	var lines []string = utils.ReadToSlice("./input.txt")
	for i, line := range lines {
		for _, p := range reSymbol.FindAllStringIndex(line, -1) {
			symbols = append(
				symbols,
				Symbol{i - 1, p[0] - 1, 3, 3, line[p[0]:p[1]], []Part{}},
			)
		}
		for _, p := range reNumber.FindAllStringIndex(line, -1) {
			parts = append(
				parts,
				Part{i, p[0], 1, p[1] - p[0], utils.StringToInt(line[p[0]:p[1]])},
			)
		}
	}

	var sumPart1 int = 0
	for i, s := range symbols {
		for _, n := range parts {
			if isIntersecting(s, n) {
				symbols[i].adjacent = append(symbols[i].adjacent, n)
				sumPart1 += n.value
			}
		}
	}

	fmt.Println("Part 1:", sumPart1)

	var sumPart2 int = 0
	for _, s := range symbols {
		if len(s.adjacent) == 2 {
			sumPart2 += s.adjacent[0].value * s.adjacent[1].value
		}
	}

	fmt.Println("Part 2:", sumPart2)
}

func isIntersecting(obj1 Symbol, obj2 Part) bool {
	return obj1.x < obj2.x+obj2.height &&
		obj1.x+obj1.height > obj2.x &&
		obj1.y < obj2.y+obj2.width &&
		obj1.y+obj1.width > obj2.y
}
