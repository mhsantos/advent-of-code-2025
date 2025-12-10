package main

import (
	"fmt"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day7 <part1|part2> <input filename>"

func main() {
	part, filename := argparser.ParseArgs(2)
	if part == argparser.Invalid {
		fmt.Println(usage)
		return
	}

	lines := filereader.ReadInput(filename)
	if part == argparser.Part1 {
		fmt.Println("Beam splits", countBeamSplits(lines, false))
	} else {
		fmt.Println("Possible paths", traverseManifold(findStartIndex(lines[0]), lines))
	}

}

func findStartIndex(line string) int {
	for idx, char := range line {
		if char == 'S' {
			return idx
		}
	}
	return -1
}

func countBeamSplits(lines []string, print bool) int {
	count := 0
	rows := len(lines)
	cols := 0
	var levels [][]rune
	for _, line := range lines {
		if cols == 0 {
			cols = len(line)
		}
		levels = append(levels, []rune(line))
	}
	for row := 1; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if levels[row-1][col] == 'S' {
				levels[row][col] = '|'
				break
			}
			if levels[row-1][col] == '|' {
				if levels[row][col] == '.' {
					levels[row][col] = '|'
				}
				if levels[row][col] == '^' {
					count++
					if levels[row][col-1] == '.' {
						levels[row][col-1] = '|'
					}
					if levels[row][col+1] == '.' {
						levels[row][col+1] = '|'
					}
				}
			}
		}
		if print {
			fmt.Println(string(levels[row]), count)
		}
	}
	return count
}

func traverseManifold(col int, lines []string) int {
	row := 2
	cols := len(lines[0])
	counter := make([]int, len(lines[0]))
	counter[col] = 1
	for row < len(lines) {
		newCounter := make([]int, cols)
		for col := range cols {
			if lines[row][col] == '^' {
				newCounter[col-1] += counter[col]
				newCounter[col+1] += counter[col]
			} else {
				newCounter[col] += counter[col]
			}
		}
		counter = newCounter
		row += 2
	}
	total := 0
	for _, count := range counter {
		total += count
	}
	return total
}

func printAllLevels(levels [][]rune) {
	for _, level := range levels {
		fmt.Println(string(level))
	}
}
