package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed example.txt
var input string

type Node struct {
	id    string
	left  string
	right string
}

type test func(string) bool

var dessertMap map[string]Node

var instructions string

func initGlobals() {
	dessertMap = map[string]Node{}
	parseInput()
}

func parseInput() {
	moveInstructions, dessert, _ := strings.Cut(input, "\n\n")
	instructions = moveInstructions

	nodes := strings.Split(dessert, "\n")
	dessertRegex := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)

	for _, nodeDescription := range nodes {
		matches := dessertRegex.FindStringSubmatch(nodeDescription)

		nodeId := matches[1]
		leftId := matches[2]
		rightId := matches[3]

		dessertMap[nodeId] = Node{id: nodeId, left: leftId, right: rightId}
	}
}

func findCycleLength(pos Node) int {
	seenEndpoints := map[Node]int{}
	var steps int

	cycleLength, instructionPtr, endPoint := solve(0, pos, func(s string) bool { return strings.HasSuffix(s, "Z") })

	for {
		_, isPresent := seenEndpoints[endPoint]
		seenEndpoints[endPoint] = instructionPtr
		steps, instructionPtr, endPoint = solve(instructionPtr, endPoint, func(s string) bool { return strings.HasSuffix(s, "Z") })
		cycleLength += steps

		// Im to stupid as for some reason it just doe not work as condition in the for loop
		if isPresent && seenEndpoints[endPoint] == instructionPtr {
			break
		}
	}

	return cycleLength
}

func solve(startInstruction int, start Node, endCondition test) (int, int, Node) {
	var steps int
	currentNode := start
	for steps = startInstruction; !endCondition(currentNode.id); steps++ {
		direction := string(instructions[steps%len(instructions)])
		if direction == "R" {
			currentNode = dessertMap[currentNode.right]
		}
		if direction == "L" {
			currentNode = dessertMap[currentNode.left]
		}
	}
	return steps, (startInstruction + steps) % len(instructions), currentNode
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
					t := b
					b = a % b
					a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
					result = LCM(result, integers[i])
	}

	return result
}

func part1() {
	steps, _, _ := solve(0, dessertMap["AAA"], func(s string) bool { return s == "ZZZ" })
	fmt.Printf("Part 1: %d\n", steps)
}

func part2() {
	startingPoints := make([]string, 0, len(dessertMap))
	for k := range dessertMap {
		if strings.HasSuffix(k, "A") {
			startingPoints = append(startingPoints, k)
		}
	}

	cycles := []int{}
	for _, startingPoint := range startingPoints {
		cycles = append(cycles, findCycleLength(dessertMap[startingPoint]))
	}
	//lcm cycles
	lcm := LCM(cycles[0], cycles[1], cycles[2:]...)
	fmt.Printf("Part 2: %d\n", lcm)
}

func main() {
	initGlobals()
	// part1()
	part2()
}
