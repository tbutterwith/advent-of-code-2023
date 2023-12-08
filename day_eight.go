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

	for _, line := range lines[2:] {
		node_name := line[0:3]

		nodes := strings.Split(line[7:15], ", ")

		map_nodes[node_name] = node{node_name, nodes[0], nodes[1]}
	}

	queue := instructions
	current_node := map_nodes["AAA"]
	total_steps := 0
	for len(queue) > 0 {
		next_instr := queue[0]
		queue = queue[1:]

		if current_node.name == "ZZZ" {
			log.Infof("Total steps: %d", total_steps)
			return
		}

		current_node = traverse(current_node, string(next_instr))
		total_steps += 1

		if len(queue) == 0 {
			queue = instructions
		}
	}

}

func traverse(current_node node, instruction string) (next_node node) {
	if instruction == "L" {
		next_node = map_nodes[current_node.left]
	} else {
		next_node = map_nodes[current_node.right]
	}
	return
}
