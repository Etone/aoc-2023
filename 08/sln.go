package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	id    string
	left  string
	right string
}

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

func part1() {
	currentNode := dessertMap["AAA"]
	var steps int
	for steps = 0; currentNode.id != "ZZZ"; steps++ {
		direction := string(instructions[steps%len(instructions)])
		if direction == "R" {
			currentNode = dessertMap[currentNode.right]
		}
		if direction == "L" {
			currentNode = dessertMap[currentNode.left]
		}
	}
	fmt.Printf("Part 1: %d\n", steps)
}

func part2() {
	currentPositions := make([]Node, 0, len(dessertMap))
	for k := range dessertMap {
		if strings.HasSuffix(k, "A") {
			currentPositions = append(currentPositions, dessertMap[k])
		}
	}

	var end bool
	var steps int
	for steps = 0; !end; steps++ {
		direction := string(instructions[steps%len(instructions)])
		end = true
		for i, currentPosition := range currentPositions {
			if direction == "R" {
				currentPositions[i] = dessertMap[currentPosition.right]
			}
			if direction == "L" {
				currentPositions[i] = dessertMap[currentPosition.left]
			}
		}
		for _, pos := range currentPositions {
			end = end && strings.HasSuffix(pos.id, "Z")
		}
	}

	fmt.Printf("Part 2: %d\n", steps)
}

func main() {
	initGlobals()
	part1()
	part2()
}
