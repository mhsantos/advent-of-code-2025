package main

import (
	"fmt"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day3 <part1|part2> <input filename>"

func main() {
	part, filename := argparser.ParseArgs(2)
	if part == argparser.Invalid {
		fmt.Println(usage)
		return
	}

	lines := filereader.ReadInput(filename)
	fmt.Println("Maximum joltage is:", findMaxJoltage(lines))
}

func findMaxJoltage(lines []string) int {
	maxJoltage := 0
	for _, line := range lines {
		maxFirstDigit := rune(line[0])
		maxSecondDigit := '0'
		index := 1
		indexMaxDigit := 0
		for index < len(line)-1 {
			if rune(line[index]) > maxFirstDigit {
				maxFirstDigit = rune(line[index])
				indexMaxDigit = index
			}
			index++
			if maxFirstDigit == '9' {
				break
			}
		}
		index = indexMaxDigit + 1
		for index < len(line) {
			if rune(line[index]) > maxSecondDigit {
				maxSecondDigit = rune(line[index])
			}
			index++
			if maxSecondDigit == '9' {
				break
			}
		}
		joltage := int((maxFirstDigit-'0')*10 + (maxSecondDigit - '0'))
		maxJoltage += joltage
	}
	return maxJoltage
}
