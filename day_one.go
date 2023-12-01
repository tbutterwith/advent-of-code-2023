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
	file, err := os.Open("inputs/1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	total := 0

	pattern := `(one|two|three|four|five|six|seven|eight|nine|\d)`

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		first_num := ""
		last_num := ""
		for _, char_rune := range line {
			if unicode.IsDigit(char_rune) {
				if first_num == "" {
					first_num = string(char_rune)
				}
				last_num = string(char_rune)
			}
		}
		num_str := first_num + last_num

		num, _ := strconv.Atoi(num_str)

		total += num
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	println(total)

}
