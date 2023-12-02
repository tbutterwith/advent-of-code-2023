package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	game_sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		substrings := strings.Split(line, ":")

		game_num_str := string(regexp.MustCompile(`\d+`).Find([]byte(substrings[0])))
		game_num, _ := strconv.Atoi(game_num_str)

		count := ""
		colour := ""

		is_valid_game := true

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

			num_count, _ := strconv.Atoi(count)
			if colour == "red" && num_count > 12 {
				is_valid_game = false
				break
			} else if colour == "green" && num_count > 13 {
				is_valid_game = false
				break
			} else if colour == "blue" && num_count > 14 {
				is_valid_game = false
				break
			}
		}

		if is_valid_game {
			game_sum += game_num
		}
	}

	println(game_sum)
}
