package main

import (
	"fmt"
	"math"

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
	if part == argparser.Part1 {
		fmt.Println("Maximum joltage is:", findMaxJoltage(lines, 2))
		return
	}
	fmt.Println("Maximum joltage is:", findMaxJoltage(lines, 12))
}

func findMaxJoltage(lines []string, decimalPlaces int) int {
	maxJoltage := 0
	for _, line := range lines {
		joltage := 0
		currentDecimal := 0
		currentDecimalMax := rune(0)
		indexToStart := 0
		for currentDecimal < decimalPlaces {
			index := indexToStart
			for index < len(line)-(decimalPlaces-1-currentDecimal) {
				currentDigit := rune(line[index]) - '0'
				if currentDigit > currentDecimalMax {
					currentDecimalMax = currentDigit
					indexToStart = index + 1
				}
				if currentDecimalMax == 9 {
					indexToStart = index + 1
					index++
					break
				}
				index++
			}
			joltage += int(currentDecimalMax) * int(math.Pow(10, float64(decimalPlaces-currentDecimal-1)))
			currentDecimal++
			currentDecimalMax = 0
			index = currentDecimal
		}
		maxJoltage += joltage
	}
	return maxJoltage
}
