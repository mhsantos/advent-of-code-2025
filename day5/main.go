package main

import (
	"fmt"
	"slices"
	"sort"
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
		intervals, ids, err := intervalsAndIDs(lines)
		if err != nil {
			fmt.Println("Error processing lines:", err)
			return
		}
		count := 0
		for _, id := range ids {
			if isIDInIntervals(id, intervals) {
				count++
			}
		}
		fmt.Println("Number of IDs in intervals:", count)
		return
	}
}

func intervalsAndIDs(lines []string) ([][]int, []int, error) {
	processingIntervals := true
	ids := make([]int, 0)
	intervals := make([][]int, 0)
	for _, line := range lines {
		if processingIntervals {
			if line == "" {
				processingIntervals = false
				continue
			}
			interval := strings.Split(line, "-")
			start, err := strconv.Atoi(interval[0])
			if err != nil {
				return nil, nil, err
			}
			end, err := strconv.Atoi(interval[1])
			if err != nil {
				return nil, nil, err
			}
			intervals = append(intervals, []int{start, end})

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

func isIDInIntervals(id int, intervals [][]int) bool {
	for _, interval := range intervals {
		if id >= interval[0] && id <= interval[1] {
			return true
		}
	}
	return false
}

// Alternative approach using maps and binary search
func processLines(lines []string) (map[int]rune, []int, error) {
	intervals := make(map[int]rune)
	ids := []int{}
	processingIntervals := true
	for _, line := range lines {
		if processingIntervals {
			if line == "" {
				processingIntervals = false
				continue
			}
			interval := strings.Split(line, "-")
			start, err := strconv.Atoi(interval[0])
			if err != nil {
				return nil, nil, err
			}
			intervals[start] = 'S'
			end, err := strconv.Atoi(interval[1])
			if _, ok := intervals[end]; !ok {
				intervals[end] = 'E'
			}
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

// Alternative approach using maps and binary search
func intervalAsSlice(intervalMap map[int]rune) []interval {
	sortedIntervals := make([]interval, 0)
	intervals := make([]interval, 0)
	for id, kind := range intervalMap {
		sortedIntervals = append(sortedIntervals, interval{id: id, kind: kind})
	}
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i].id < sortedIntervals[j].id
	})
	prevKind := 'X'
	for _, interv := range sortedIntervals {
		if interv.kind == 'S' && prevKind != 'S' {
			intervals = append(intervals, interv)
			prevKind = 'S'
		}
		if interv.kind == 'E' {
			if prevKind == 'E' {
				intervals = append(intervals[:len(intervals)-1], interv)
			} else {
				intervals = append(intervals, interv)
			}
			prevKind = 'E'
		}
	}
	return intervals
}

// Alternative approach using maps and binary search
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
