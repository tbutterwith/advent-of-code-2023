package main

import (
	"bufio"
	"fmt"
	"os"
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

	search_schematic(schematics)
}

func search_schematic(schematic [][]rune) {

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

func is_valid(schematic [][]rune, min_x int, min_y int, max_x int, max_y int) bool {
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
	is_valid := false
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

	return is_valid
}
