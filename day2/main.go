package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day1 <part1|part2> <input filename>"

func main() {
	part, filename := argparser.ParseArgs(2)
	if part == argparser.Invalid {
		fmt.Println(usage)
		return
	}

	lines := filereader.ReadInput(filename)
	invalidSum, err := findInvalidIdsSum(lines[0], part == argparser.Part1)
	if err != nil {
		fmt.Println("Error finding invalid IDs:", err)
		return
	}
	fmt.Println("Sum of invalid IDs is:", invalidSum)

}

func findInvalidIdsSum(line string, part1 bool) (int, error) {
	sum := 0
	line = strings.ReplaceAll(line, " ", "")
	idRanges := strings.Split(line, ",")
	for _, idRange := range idRanges {
		bounds := strings.Split(idRange, "-")
		from, err := strconv.Atoi(bounds[0])
		if err != nil {
			return 0, errors.New("invalid id " + bounds[0])
		}
		to, err := strconv.Atoi(bounds[1])
		if err != nil {
			return 0, errors.New("invalid id " + bounds[1])
		}
		for i := from; i <= to; i++ {
			asStr := strconv.Itoa(i)
			if part1 {
				if isIdInvalidPart1(asStr) {
					sum += i
				}
			} else {
				if isIdInvalidPart2(asStr) {
					sum += i
				}
			}
		}
	}
	return sum, nil
}

func isIdInvalidPart1(id string) bool {
	if len(id)%2 != 0 {
		return false
	}
	mid := len(id) / 2
	return id[:mid] == id[mid:]
}

func isIdInvalidPart2(id string) bool {
outer:
	for chunkSize := 1; chunkSize <= len(id)/2; chunkSize++ {
		if len(id)%chunkSize != 0 {
			continue
		}
		firstChunk := id[0:chunkSize]
		for index := chunkSize; index < len(id); index += chunkSize {
			if id[index:index+chunkSize] != firstChunk {
				continue outer
			}
		}
		return true
	}
	return false
}
