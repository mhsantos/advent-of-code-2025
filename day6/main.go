package main

import (
	"fmt"
	"math"
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
	var operations []operation
	var err error
	if part == argparser.Part1 {
		operations, err = processLines(lines)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			return
		}
	} else {
		operations = processLinesPart2(lines)
	}
	fmt.Println("Grand total: ", runAllOperations(operations))
}

func processLinesPart2(lines []string) []operation {
	operations := make([]operation, 0)
	var op operation
	for idx, val := range lines[len(lines)-1] {
		if val != ' ' {
			if idx > 0 {
				operations = append(operations, op)
			}
			op = operation{}
			op.operator = val
		}
		decimalPlace := 0
		operand := 0
		for i := len(lines) - 2; i > -1; i-- {
			if len(lines[i]) > idx {
				if lines[i][idx] != ' ' {
					digit := int(lines[i][idx] - '0')
					operand += digit * int(math.Pow(float64(10), float64(decimalPlace)))
					decimalPlace++
				}
			}
		}
		if operand > 0 {
			op.elements = append(op.elements, operand)
		}
	}
	operations = append(operations, op)
	return operations
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
