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
	runeMatrix := convertToRuneMatrix(lines)
	var accessible int
	if part == argparser.Part1 {
		accessible = findAccessibleRolls(runeMatrix, false)
	} else {
		stillAccessible := findAccessibleRolls(runeMatrix, true)
		for stillAccessible != 0 {
			accessible += stillAccessible
			stillAccessible = findAccessibleRolls(runeMatrix, true)
		}
	}
	println("Accessible rolls:", accessible)
}

// convertToRuneMatrix converts a slice of strings to a matrix of runes to allow for easier
// replacement of individual characters.
func convertToRuneMatrix(lines []string) [][]rune {
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

// findAccessibleRolls parses the input map and return the count of rolls (cells with the @ symbol)
// that have less than 4 adjacent rolls. This includes diagonal adjacency.
// The replace parameters indicates if it whould replace the roll with a '.' when it's accessible.
func findAccessibleRolls(lines [][]rune, replace bool) int {
	count := 0
	rows := len(lines)
	cols := len(lines[0])

	for row := range rows {
		for col := range cols {
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
				if replace {
					lines[row][col] = '.'
				}
			}
		}
	}
	return count
}
