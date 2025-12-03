package argparser

import (
	"fmt"
	"os"
	"slices"
)

// ParseArgs checks the command line arguments and returns the corresponding filename
// if the arguments are valid, otherwise, it prints usage and returns false
func ParseArgs(day int, usage string) (string, bool) {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(usage)
		return "", false
	}
	validArgs := []string{"part1", "part2", "example"}
	if !slices.Contains(validArgs, args[0]) {
		fmt.Println(usage)
		return "", false
	}
	switch args[0] {
	case "part1":
		return fmt.Sprintf("day%da.txt", day), true
	case "part2":
		return fmt.Sprintf("day%db.txt", day), true
	case "example":
		return fmt.Sprintf("day%dexample.txt", day), true
	default:
		fmt.Println(usage)
		return "", false
	}
}
