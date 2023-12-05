package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (file_array []string) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		file_array = append(file_array, scanner.Text())
	}

	return
}
