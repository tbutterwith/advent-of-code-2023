package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var num_map map[string]string

func main() {
	// Open the file
	file, err := os.Open("inputs/1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	num_map = make(map[string]string)
	num_map["one"] = "1"
	num_map["two"] = "2"
	num_map["three"] = "3"
	num_map["four"] = "4"
	num_map["five"] = "5"
	num_map["six"] = "6"
	num_map["seven"] = "7"
	num_map["eight"] = "8"
	num_map["nine"] = "9"
	num_map["zero"] = "0"

	nums := [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	total := 0

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()

		first_digit := ""
		last_digit := ""

		// iterate on line char by char
		for i, char := range line {
			// if digit, stop
			if unicode.IsDigit(char) {
				if first_digit == "" {
					first_digit = string(char)
				}
				last_digit = string(char)
			} else {
				// if starts to match string, look ahead
				for _, num_string := range nums {
					char_string := string(char)
					char_lookahead := 0

					// if it's a valid starting path
					for strings.HasPrefix(num_string, char_string) {
						num_char, ok := num_map[char_string]

						// if we've matched, set the vars
						if ok {
							if first_digit == "" {
								first_digit = num_char
							}
							last_digit = num_char
							break
						} else { // otherwise keep going
							char_lookahead += 1
						}
						// don't go too far along the line
						if len(line) <= i+char_lookahead {
							break
						}
						char_string = char_string + string(line[i+char_lookahead])
					}
				}
			}
		}

		num_str := first_digit + last_digit
		num, _ := strconv.Atoi(num_str)
		total += num
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	println(total)

}
