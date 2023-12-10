package main

import (
	"strings"

	"github.com/charmbracelet/log"
)

type node struct {
	name  string
	left  string
	right string
}

var map_nodes map[string]node

func main() {
	lines := ReadFile("inputs/8.txt")
	log.SetLevel(log.DebugLevel)

	instructions := lines[0]

	map_nodes = make(map[string]node)

	var starting_nodes []node
	for _, line := range lines[2:] {
		node_name := line[0:3]

		nodes := strings.Split(line[7:15], ", ")

		map_nodes[node_name] = node{node_name, nodes[0], nodes[1]}

		if string(node_name[2]) == "A" {
			starting_nodes = append(starting_nodes, map_nodes[node_name])
		}
	}

	current_nodes := starting_nodes
	var steps []int

	for _, node := range current_nodes {
		steps = append(steps, find_z_for_node(node, instructions))
	}

	log.Debug(LCM(steps[0], steps[1], steps[2:]...))

}

func find_z_for_node(starting_node node, instructions string) (steps int) {
	steps = 0
	queue := instructions
	current_node := starting_node
	for len(queue) > 0 {
		next_instr := queue[0]
		queue = queue[1:]

		if string(current_node.name[2]) == "Z" {
			return
		}

		current_node = traverse(current_node, string(next_instr))
		steps += 1

		if len(queue) == 0 {
			queue = instructions
		}
	}

	return
}

func traverse(current_node node, instruction string) (next_node node) {
	if instruction == "L" {
		next_node = map_nodes[current_node.left]
	} else {
		next_node = map_nodes[current_node.right]
	}
	return
}

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
