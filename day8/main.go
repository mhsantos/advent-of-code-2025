package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-of-code-2025/internal/argparser"
	"github.com/mhsantos/advent-of-code-2025/internal/filereader"
)

const usage = "Usage: go run ./day8 <part1|part2> <input filename>"

type box struct {
	id      int
	x, y, z int
}

type boxDistance struct {
	boxA, boxB box
	distance   float64
}

func main() {
	part, filename := argparser.ParseArgs(2)
	if part == argparser.Invalid {
		fmt.Println(usage)
		return
	}

	lines := filereader.ReadInput(filename)
	if part == argparser.Part1 {
		boxes, err := getBoxes(lines)
		if err != nil {
			return
		}
		distances := shortestDistances(boxes)
		circuits, _, _ := connectBoxes(distances, 1000, 1000)
		prod3Largest := prodNLargest(circuits, 3)
		fmt.Println("3 largest multiplied:", prod3Largest)
	} else {
		boxes, err := getBoxes(lines)
		if err != nil {
			return
		}
		distances := shortestDistances(boxes)
		_, _, lastTwo := connectBoxes(distances, len(boxes), -1)
		lastTwoX := boxes[lastTwo[0]].x * boxes[lastTwo[1]].x
		fmt.Println("Last two connected x multiplied", lastTwoX)
	}
}

func getBoxes(lines []string) ([]box, error) {
	boxes := make([]box, 0, len(lines))
	for idx, line := range lines {
		values := strings.Split(line, ",")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("error converting x value on line ", idx)
			return nil, err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("error converting y value on line ", idx)
			return nil, err
		}
		z, err := strconv.Atoi(values[2])
		if err != nil {
			fmt.Println("error converting z value on line ", idx)
			return nil, err
		}
		boxes = append(boxes, box{idx, x, y, z})
	}
	return boxes, nil
}

// shortestDistances returns an n x n size slice of boxDistance with the distance between
// each box to all other boxes
func shortestDistances(boxes []box) []boxDistance {
	distances := make([]boxDistance, 0, len(boxes))
	for a := 0; a < len(boxes); a++ {
		for b := a + 1; b < len(boxes); b++ {
			boxA := boxes[a]
			boxB := boxes[b]
			distanceX := boxA.x - boxB.x
			if distanceX < 0 {
				distanceX = distanceX * -1
			}
			distanceY := boxA.y - boxB.y
			if distanceY < 0 {
				distanceY = distanceY * -1
			}
			distanceZ := boxA.z - boxB.z
			if distanceZ < 0 {
				distanceZ = distanceZ * -1
			}
			distance := euclideanDistance(distanceX, distanceY, distanceZ)
			distances = append(distances, boxDistance{boxA, boxB, distance})
		}
	}
	slices.SortFunc(distances, func(a, b boxDistance) int {
		if a.distance < b.distance {
			return -1
		} else if b.distance < a.distance {
			return 1
		} else {
			return 0
		}
	})
	return distances
}

// connectBoxes connects boxes into circuits.
// When the n parameter is informed, it connects n boxes together a map from circuit to boxes and a second map from box to circuit
// When n <= 0, it joins boxes together until there's only one circuit. It still returns the maps described above, plus a 2 positions array
// with the ids of the last 2 boxes connected
func connectBoxes(distances []boxDistance, nBoxes, n int) (map[string][]box, map[int]string, [2]int) {
	circuits := make(map[string][]box, 0)
	boxCircuit := make(map[int]string, 0)
	var lastTwo [2]int

	circuitIndex := 0
	i := 0
outer:
	for {
		if n > 0 && i == n {
			break
		}
		if len(circuits) == 1 {
			for _, v := range circuits {
				if len(v) == nBoxes {
					break outer
				}
				break
			}
		}
		distance := distances[i]
		i++
		circuitA, okA := boxCircuit[distance.boxA.id]
		circuitB, okB := boxCircuit[distance.boxB.id]
		if okA {
			if okB { // both boxes are on their circuits
				if circuitA == circuitB { // both boxes are on the same circuit. Nothing to be done
					continue
				}
				// boxes are on different circuits.
				// Merging both circuits by moving members of circuitB to circuitA
				boxesB := circuits[circuitB]
				for _, boxB := range boxesB {
					boxCircuit[boxB.id] = circuitA
					circuits[circuitA] = append(circuits[circuitA], boxB)
				}
				lastTwo[0] = distance.boxA.id
				lastTwo[1] = distance.boxB.id
				delete(circuits, circuitB)
			} else { // boxB is not in any circuit. Will add it to boxA's circuit
				boxCircuit[distance.boxB.id] = circuitA
				circuits[circuitA] = append(circuits[circuitA], distance.boxB)
				lastTwo[0] = distance.boxA.id
				lastTwo[1] = distance.boxB.id
			}
		} else {
			if okB { // boxA is not in any circuit. Will add it to boxB's circuit
				boxCircuit[distance.boxA.id] = circuitB
				circuits[circuitB] = append(circuits[circuitB], distance.boxA)
			} else { // None of the boxes are in any circuit. Will create a circuit and add them both to it
				circuitName := fmt.Sprintf("circuit%d", circuitIndex)
				members := make([]box, 0, 2)
				members = append(members, distance.boxA)
				members = append(members, distance.boxB)
				circuits[circuitName] = members
				boxCircuit[distance.boxA.id] = circuitName
				boxCircuit[distance.boxB.id] = circuitName
				circuitIndex++
			}
			lastTwo[0] = distance.boxA.id
			lastTwo[1] = distance.boxB.id
		}
	}
	return circuits, boxCircuit, lastTwo
}

func prodNLargest(circuits map[string][]box, n int) int {
	prod := 1
	sizes := orderdedFromLargest(circuits)
	for i := range n {
		prod *= sizes[i]
	}
	return prod
}

func orderdedFromLargest(circuits map[string][]box) []int {
	sizes := make([]int, 0, len(circuits))
	for _, v := range circuits {
		sizes = append(sizes, len(v))
	}
	slices.SortFunc(sizes, func(a, b int) int {
		return b - a
	})
	return sizes
}

func euclideanDistance(x, y, z int) float64 {
	return math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2) + math.Pow(float64(z), 2))
}
