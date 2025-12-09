package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day6 <part1|part2> <input filename>"

type operation struct {
	operator rune
	elements []int
}

func main() {
	part, filename := argparser.ParseArgs(2)
	if part == argparser.Invalid {
		fmt.Println(usage)
		return
	}

	lines := filereader.ReadInput(filename)
	if part == argparser.Part1 {
		operations, err := processLines(lines)
		if err == nil {
			fmt.Println("Grand total: ", runAllOperations(operations))
		}
	}
}

func processLines(lines []string) ([]operation, error) {
	operations := make([]operation, 0)
	for lineNumber, line := range lines {
		values := strings.Fields(line)
		for idx, value := range values {
			if len(operations) == idx {
				operations = append(operations, operation{})
			}
			if lineNumber == len(lines)-1 {
				operations[idx].operator = rune(value[0])
			} else {
				operand, err := strconv.Atoi(value)
				if err != nil {
					fmt.Println("error converting value", err)
					return nil, err
				}
				operations[idx].elements = append(operations[idx].elements, operand)
			}
		}
	}
	return operations, nil
}

func runAllOperations(ops []operation) int {
	total := 0
	for _, op := range ops {
		total += runOperation(op)
	}
	return total
}

func runOperation(op operation) int {
	total := 0
	if op.operator == '*' {
		total = 1
	}
	for _, value := range op.elements {
		if op.operator == '+' {
			total += value
		} else {
			total *= value
		}
	}
	return total
}
