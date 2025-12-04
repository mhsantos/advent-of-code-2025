package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day1 <part1|part2|example>"

func main() {

	filename, ok := argparser.ParseArgs(1, usage)
	if !ok {
		return
	}

	lines := filereader.ReadInput(filename)

	if strings.HasSuffix(filename, "a.txt") || strings.HasSuffix(filename, "example.txt") {
		part1(lines, true)
	} else {
		part1(lines, false)
	}

}

func part1(lines []string, atZero bool) {
	position := 50
	timesAt0 := 0
	times := 0
	crossedTimes := 0

	for idx, line := range lines {
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			slog.Error("error converting steps", slog.String("line", line), slog.Any("error", err))
			return
		}
		switch line[0] {
		case 'L':
			position, times = moveDial(position, true, steps)
		case 'R':
			position, times = moveDial(position, false, steps)
		default:
			slog.Error("invalid input", slog.Int("line number", idx), slog.String("line", line), slog.Any("error", err))
			return
		}

		if position == 0 {
			timesAt0++
		}
		crossedTimes += times
	}
	if atZero {
		fmt.Println("Password is:", timesAt0)
	} else {
		fmt.Println("Password is:", crossedTimes)
	}
}

// move dial moves the dial to left or right and returns the new position and how many times it crossed 0
func moveDial(position int, left bool, steps int) (int, int) {
	times := steps / 100
	steps = steps % 100
	for i := 0; i < steps; i++ {
		if left {
			position--
		} else {
			position++
		}
		switch position {
		case -1:
			position = 99
		case 100:
			position = 0
			times++
		case 0:
			times++
		}
	}
	return position, times
}
