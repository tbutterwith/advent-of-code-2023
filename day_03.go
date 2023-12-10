package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func main() {
	// Open the file
	file, err := os.Open("inputs/3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var schematics [][]rune

	for scanner.Scan() {
		var line_slice []rune

		for _, char := range scanner.Text() {
			line_slice = append(line_slice, char)
		}

		schematics = append(schematics, line_slice)
	}

	// search_schematic_for_all_nums(schematics)
	get_gears(schematics)
}

func get_gears(schematic [][]rune) {
	gear_total := 0
	for y, line := range schematic {
		for x, char := range line {
			if string(char) == "*" {
				gear_total += calculate_gears(schematic, x, y)
			}
		}
	}
	fmt.Println(gear_total)
}

func calculate_gears(schematic [][]rune, x int, y int) int {
	min_x := x - 1
	max_x := x + 1

	min_y := y - 1
	max_y := y + 1

	if min_x < 0 {
		min_x = 0
	}
	if min_y < 0 {
		min_y = 0
	}

	if max_x >= len(schematic[0]) {
		max_x = len(schematic[0]) - 1
	}
	if max_y >= len(schematic) {
		max_y = len(schematic) - 1
	}

	var found_nums []int
	total := 0

	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {
			char := schematic[y][x]
			if unicode.IsDigit(char) {
				// get the full number
				full_num := get_full_num(schematic, x, y)
				// put in found_nums if doesn't exist
				if !slices.Contains(found_nums, full_num) {
					found_nums = append(found_nums, full_num)
				}
			}
		}
	}

	if len(found_nums) == 2 {
		total = found_nums[0] * found_nums[1]
	}

	return total
}

func get_full_num(schematic [][]rune, x_coord, y_coord int) (full_num int) {
	line := schematic[y_coord]

	num_str := string(line[x_coord])
	pointer := x_coord - 1

	for pointer >= 0 && unicode.IsDigit(line[pointer]) {
		num_str = string(line[pointer]) + num_str
		pointer--
	}

	pointer = x_coord + 1
	for pointer < len(line) && unicode.IsDigit(line[pointer]) {
		num_str += string(line[pointer])
		pointer++
	}

	full_num, _ = strconv.Atoi(num_str)

	return
}

func search_schematic_for_all_nums(schematic [][]rune) {

	total := 0

	for y, line := range schematic {
		num := ""
		search := false
		min_x := 0
		max_x := 0
		min_y := 0
		max_y := 0

		for x, char := range line {
			// if we find a number
			if unicode.IsDigit(char) {

				if !search {
					min_x = x - 1
					min_y = y - 1
				}
				num += string(char)
				max_x = x + 1
				max_y = y + 1

				search = true

			}
			if !unicode.IsDigit(char) || x == len(line)-1 {
				if num != "" {
					if is_valid(schematic, min_x, min_y, max_x, max_y) {
						num_int, _ := strconv.Atoi(num)
						total += num_int
					}

					num = ""
					search = false
				}
			}
		}
	}

	println(total)
}

func is_valid(schematic [][]rune, min_x, min_y, max_x, max_y int) (is_valid bool) {
	if min_x < 0 {
		min_x = 0
	}
	if min_y < 0 {
		min_y = 0
	}

	if max_x >= len(schematic[0]) {
		max_x = len(schematic[0]) - 1
	}
	if max_y >= len(schematic) {
		max_y = len(schematic) - 1
	}
	is_valid = false
	print("\n")
	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {
			char := schematic[y][x]
			if !unicode.IsDigit(char) && string(char) != "." && string(char) != "\n" {
				is_valid = true
			}
		}
		print("\n")
	}

	return
}
