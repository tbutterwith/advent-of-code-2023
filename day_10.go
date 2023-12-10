package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type maze_node struct {
	name      string
	tile_type byte
	prev      string
	next      string
}

var maze map[string]maze_node
var path_nodes []string

func make_id(x, y int) (id string) {
	id = strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, ",")
	return
}

func main() {
	lines := ReadFile("inputs/10.txt")
	log.SetLevel(log.DebugLevel)

	maze = make(map[string]maze_node)

	var start_node string

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			id := make_id(x, y)

			tile := lines[y][x]

			if tile == '.' {
				continue
			}

			if tile == 'S' {
				start_node = make_id(x, y)
				tile = infer_start_node_type(x, y, lines)
			}

			cur_node, ok := maze[id]
			if !ok {
				cur_node = maze_node{id, tile, "", ""}
				maze[id] = cur_node
			} else {
				cur_node.tile_type = tile
			}

			var prev_tile_id string
			var next_tile_id string

			if tile == '|' {
				prev_tile_id = make_id(x, y-1)
				next_tile_id = make_id(x, y+1)
			} else if tile == '-' {
				prev_tile_id = make_id(x-1, y)
				next_tile_id = make_id(x+1, y)
			} else if tile == 'L' {
				prev_tile_id = make_id(x, y-1)
				next_tile_id = make_id(x+1, y)
			} else if tile == 'J' {
				prev_tile_id = make_id(x, y-1)
				next_tile_id = make_id(x-1, y)
			} else if tile == '7' {
				prev_tile_id = make_id(x-1, y)
				next_tile_id = make_id(x, y+1)
			} else if tile == 'F' {
				prev_tile_id = make_id(x, y+1)
				next_tile_id = make_id(x+1, y)
			}

			cur_node.prev = prev_tile_id
			cur_node.next = next_tile_id

			maze[id] = cur_node
		}
	}
	traverse_maze(start_node, "", start_node, maze, 0)

	calc_inner_area(lines, path_nodes)
}

func calc_inner_area(maze, path_nodes []string) {
	area := 0
	for y, line := range maze {
		in_path := false
		for x, tile := range line {
			id := make_id(x, y)
			if slices.Contains(path_nodes, id) {
				if slices.Contains([]string{"|", "L", "J"}, string(tile)) {
					in_path = !in_path
				}
			} else if in_path {
				log.Debug("", "y", y, "x", x)
				area += 1
			}
		}
	}

	log.Infof("Area: %d", area)
}

func traverse_maze(start_id, prev_id, node_id string, maze map[string]maze_node, step_count int) {
	current_node := maze[node_id]

	if !slices.Contains(path_nodes, node_id) {
		path_nodes = append(path_nodes, node_id)
	}

	if node_id == start_id && step_count > 0 {
		log.Info(step_count / 2)
		return
	} else {

		if prev_id == current_node.next {
			traverse_maze(start_id, node_id, current_node.prev, maze, step_count+1)
		} else {
			traverse_maze(start_id, node_id, current_node.next, maze, step_count+1)
		}
	}
}

func infer_start_node_type(x, y int, lines []string) (start_type byte) {
	var north byte
	var south byte
	var east byte
	var west byte

	if x < len(lines[0])-1 {
		east = lines[y][x+1]
	}
	if x >= 1 {
		west = lines[y][x-1]
	}

	if y < len(lines)-1 {
		south = lines[y+1][x]
	}
	if y >= 1 {
		north = lines[y-1][x]
	}

	if north == '|' {
		if east == '-' || east == 'J' || east == '7' {
			start_type = 'L'
		} else if west == '-' || west == 'L' || west == 'F' {
			start_type = 'J'
		} else if south == '|' {
			start_type = '|'
		}
	} else if east == '-' || east == 'J' || east == 'F' || east == '7' {
		if west == '-' {
			start_type = '-'
		} else if south == '|' {
			start_type = 'F'
		} else if south == 'J' {
			start_type = 'F'
		}
	} else if west == '-' || west == 'L' || west == 'F' {
		if south == '|' {
			start_type = '7'
		}
	}

	log.Debugf("Start is %s", string(start_type))

	return
}
