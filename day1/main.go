package main

import (
	"fmt"
	"github.com/polysalama/aoc_2023/utils"
	"strconv"
)

func main() {
	var lines []string = utils.ReadToSlice("./day1/input.txt")

	var sumPart1 int = 0
	var sumPart2 int = 0

	numbers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	numbersAsWords := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	for _, line := range lines {

		partOneNumbers := utils.FindAllNumbersInString(line, numbers)
		partTwoNumbers := utils.FindAllNumbersInString(line, append(append([]string{}, numbers...), numbersAsWords...))

		n1, err1 := strconv.Atoi(
			utils.StringToNumberString(partOneNumbers[0]) + utils.StringToNumberString(partOneNumbers[len(partOneNumbers)-1]),
		)
		if err1 != nil {
			fmt.Println(err1)
		}
		sumPart1 += n1

		n2, err2 := strconv.Atoi(
			utils.StringToNumberString(partTwoNumbers[0]) + utils.StringToNumberString(partTwoNumbers[len(partTwoNumbers)-1]),
		)

		if err2 != nil {
			fmt.Println(err2)
		}
		sumPart2 += n2
	}

	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}
