package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day5 <part1|part2> <input filename>"

type interval struct {
	id   int
	kind rune
}

func main() {
	part, filename := argparser.ParseArgs(2)
	if part == argparser.Invalid {
		fmt.Println(usage)
		return
	}

	lines := filereader.ReadInput(filename)
	if part == argparser.Part1 {
		intervals, ids, err := processLines(lines)
		if err != nil {
			fmt.Println("Error processing lines:", err)
			return
		}
		count := 0
		for _, id := range ids {
			if isInInterval(id, intervals) {
				count++
			}
		}
		fmt.Println("Number of IDs in intervals:", count)
		return
	} else {
		intervals, _, err := processLines(lines)
		if err != nil {
			fmt.Println("Error processing lines:", err)
			return
		}
		count := 0
		for i := 1; i < len(intervals); i++ {
			if intervals[i].kind == 'E' && intervals[i-1].kind != 'S' {
				fmt.Println("Unexpected E after E")
				return
			}
			if intervals[i].kind == 'S' && intervals[i-1].kind != 'E' {
				fmt.Println("Unexpected S after S")
				return
			}
			if intervals[i].kind == 'E' {
				count += intervals[i].id - intervals[i-1].id + 1
			}
		}
		fmt.Println("Total number of fresh items:", count)
	}
}

// processLines gets the id ranges and inserts them in a map o interval in the folloing format:
// 1:S,5:E,13:S,147:E,165:S,2030:E
func processLines(lines []string) ([]interval, []int, error) {
	intervals := make([]interval, 0)
	ids := []int{}
	processingIntervals := true
	for _, line := range lines {
		if processingIntervals {
			if line == "" {
				processingIntervals = false
				continue
			}
			startend := strings.Split(line, "-")
			start, err := strconv.Atoi(startend[0])
			if err != nil {
				return nil, nil, err
			}
			end, err := strconv.Atoi(startend[1])
			if err != nil {
				return nil, nil, err
			}
			insertInterval(&intervals, interval{start, 'S'}, interval{end, 'E'})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, err
			}
			ids = append(ids, id)
		}
	}
	return intervals, ids, nil
}

// insertInteval uses binary search to figure where to put the start and end values.
// it also takes care of overlapping intervals by seeing a position where the interval is suppose
// to be inserted is inside or outside of an existing interval.
func insertInterval(intervals *[]interval, start interval, end interval) {
	if len(*intervals) == 0 {
		*intervals = append(*intervals, start, end)
	} else {
		newIntervals := make([]interval, 0)
		startPos, _ := slices.BinarySearchFunc(*intervals, start.id, func(i interval, id int) int {
			return i.id - id
		})
		if startPos > 0 {
			newIntervals = append(newIntervals, (*intervals)[:startPos]...)
			if (*intervals)[startPos-1].kind == 'E' {
				newIntervals = append(newIntervals, start)
			}
		} else {
			newIntervals = append(newIntervals, start)
		}
		endPos, ok := slices.BinarySearchFunc(*intervals, end.id, func(i interval, id int) int {
			return i.id - id
		})
		if endPos < len(*intervals) {
			if (*intervals)[endPos].kind == 'E' {
				newIntervals = append(newIntervals, (*intervals)[endPos:]...)
			} else { // it's S
				if ok {
					newIntervals = append(newIntervals, (*intervals)[endPos+1:]...)
				} else {
					newIntervals = append(newIntervals, end)
					newIntervals = append(newIntervals, (*intervals)[endPos:]...)
				}
			}
		} else {
			newIntervals = append(newIntervals, end)
		}
		*intervals = newIntervals
	}
}

// isInInterval uses binary search to figure if a value is inside the ranges of valid ids
func isInInterval(id int, intervals []interval) bool {
	intervalPos, ok := slices.BinarySearchFunc(intervals, id, func(i interval, id int) int {
		return i.id - id
	})
	if ok {
		return true
	}
	if intervalPos == 0 || intervalPos == len(intervals) {
		return false
	}
	if intervals[intervalPos-1].kind == 'S' {
		return true
	}
	return false
}
