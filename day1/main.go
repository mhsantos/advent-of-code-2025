package main

import (
	"fmt"
	"log/slog"
	"strconv"

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

	position := 50
	timesAt0 := 0

	for idx, line := range lines {
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			slog.Error("error converting steps", slog.String("line", line), slog.Any("error", err))
			return
		}
		switch line[0] {
		case 'L':
			position = moveDial(position, true, steps)
		case 'R':
			position = moveDial(position, false, steps)
		default:
			slog.Error("invalid input", slog.Int("line number", idx), slog.String("line", line), slog.Any("error", err))
			return
		}

		if position == 0 {
			timesAt0++
		}
	}
	fmt.Println("Password is:", timesAt0)
}

// move dial moves the dial to left or right and returns the new position
func moveDial(position int, left bool, steps int) int {
	if left {
		position -= steps % 100
		if position < 0 {
			position = 100 + position
		}
		return position
	}
	return (position + steps) % 100
}
