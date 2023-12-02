package utils

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadToSlice(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func StringToNumberString(s string) string {
	stringToNumberString := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return stringToNumberString[s]
	}

	return strconv.Itoa(i)
}

func MapValuesToOrderedSlice(m map[int]string) []string {
	var keys []int
	var values []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		values = append(values, m[k])
	}
	return values
}

func FindAllNumbersInString(s string, numbers []string) []string {

	var foundStringsMap map[int]string = make(map[int]string)

	for _, number := range numbers {
		if strings.Contains(s, number) {
			foundStringsMap[strings.LastIndex(s, number)] = number
			foundStringsMap[strings.Index(s, number)] = number
		}
	}

	return MapValuesToOrderedSlice(foundStringsMap)
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
