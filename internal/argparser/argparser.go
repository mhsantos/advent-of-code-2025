package argparser

import (
	"os"
	"slices"
)

const (
	Invalid int = -1
	Part1   int = 1
	Part2   int = 2
)

// ParseArgs checks the command line arguments and returns the test type as an integer and the
// intput filename. If the arguments are invalid, it returns -1 and "".
func ParseArgs(day int) (int, string) {
	args := os.Args[1:]
	if len(args) != 2 {
		return Invalid, ""
	}
	validArgs := []string{"part1", "part2"}
	if !slices.Contains(validArgs, args[0]) {
		return Invalid, ""
	}
	switch args[0] {
	case "part1":
		return Part1, args[1]
	case "part2":
		return Part2, args[1]
	default:
		return Invalid, ""
	}
}
