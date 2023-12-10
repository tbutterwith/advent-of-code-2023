package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isSplitChar(char string) bool {
	for _, v := range [2]string{",", ";"} {
		if v == char {
			return true
		}
	}
	return false
}

func main() {
	// Open the file
	file, err := os.Open("inputs/2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	power_sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		substrings := strings.Split(line, ":")

		cube_counts := make(map[string]int)

		cube_counts["red"] = 0
		cube_counts["green"] = 0
		cube_counts["blue"] = 0

		count := ""
		colour := ""

		for _, char := range substrings[1] {
			str_char := string(char)
			if unicode.IsDigit(char) { // parse the numbers
				count += string(char)
			} else if str_char == " " { // skip spaces
				continue
			} else if isSplitChar(str_char) { // reset if we hit a delimiter
				count = ""
				colour = ""
			} else { // build the colour string
				colour += str_char
			}

			current_max, exists := cube_counts[colour]
			if exists {
				int_count, _ := strconv.Atoi(count)
				if int_count > current_max {
					cube_counts[colour] = int_count
				}
			}
		}

		// calculate powers
		game_power := cube_counts["red"] * cube_counts["green"] * cube_counts["blue"]

		power_sum += game_power
	}

	println(power_sum)
}
