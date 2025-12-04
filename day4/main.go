package main

import (
	"fmt"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day4 <part1|part2> <input filename>"

func main() {
	part, filename := argparser.ParseArgs(4)
	if part == argparser.Invalid {
		fmt.Println(usage)
		return
	}

	lines := filereader.ReadInput(filename)
	accessible := findAccessibleRolls(lines)
	println("Accessible rolls:", accessible)
}

// findAccessibleRolls parses the input map and return the count of rolls (cells with the @ symbol)
// that have less than 4 adjacent rolls. This includes diagonal adjacency.
func findAccessibleRolls(lines []string) int {
	count := 0
	rows := len(lines)
	cols := len(lines[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if lines[row][col] != '@' {
				continue
			}
			adjacentRolls := 0
			// Check all 8 directions
			rowscols := [][2]int{
				{-1, -1}, {-1, 0}, {-1, 1},
				{0, -1}, {0, 1},
				{1, -1}, {1, 0}, {1, 1},
			}
			for _, rowcol := range rowscols {
				newRow := row + rowcol[0]
				newCol := col + rowcol[1]
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
					if lines[newRow][newCol] == '@' {
						adjacentRolls++
					}
				}
			}
			if adjacentRolls < 4 {
				count++
			}
		}
	}
	return count
}
