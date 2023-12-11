package main

import (
	"math"

	"github.com/charmbracelet/log"
)

type coord struct {
	x int
	y int
}

func main() {
	lines := ReadFile("inputs/11.txt")
	log.SetLevel(log.DebugLevel)

	var galaxy_coordinates []coord

	var clear_x map[int]bool
	var clear_y map[int]bool
	clear_x = make(map[int]bool)
	clear_y = make(map[int]bool)

	for y, line := range lines {
		clear_y[y] = true
		for x, tile := range line {
			if tile == '#' {
				galaxy_coordinates = append(galaxy_coordinates, coord{x, y})
				clear_y[y] = false
				clear_x[x] = false
			}
			_, ok := clear_x[x]
			if !ok {
				clear_x[x] = true
			}
		}
	}

	empty_x := get_trues_from_map(clear_x)
	empty_y := get_trues_from_map(clear_y)

	log.Debug("", "x", empty_x)
	log.Debug("", "y", empty_y)

	distances := calc_distances_with_empty(galaxy_coordinates, empty_x, empty_y)

	log.Info(distances)
}

func get_trues_from_map(input map[int]bool) (out []int) {
	for key, value := range input {
		if value {
			out = append(out, key)
		}
	}

	return
}

func calc_distances_with_empty(galaxy_coordinates []coord, empty_x, empty_y []int) (total_distances int) {
	space := 999999

	for i, co_ord := range galaxy_coordinates {
		log.Debug(co_ord.y, co_ord.x)
		x1 := add_space(co_ord.x, space, empty_x)
		y1 := add_space(co_ord.y, space, empty_y)

		log.Debug(y1, x1)

		for j := i; j < len(galaxy_coordinates); j++ {
			x2 := add_space(galaxy_coordinates[j].x, space, empty_x)
			y2 := add_space(galaxy_coordinates[j].y, space, empty_y)

			distance := math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1))
			total_distances += int(distance)
		}
	}

	return
}

func add_space(num, space_val int, empty_spaces []int) int {
	new_num := num
	for _, val := range empty_spaces {
		if val <= num {
			new_num += space_val
		}
	}

	return new_num
}
