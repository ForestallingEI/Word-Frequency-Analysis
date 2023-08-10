package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the text from a file
	file, err := os.Open("Tarzan-triumphant.txt") // Replace with your input file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a map to store word frequency
	wordFrequency := make(map[string]int)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // Split line into words

		// Count word frequency
		for _, word := range words {
			// Remove punctuation and convert to lowercase for better counting
			word = strings.TrimRight(strings.TrimLeft(word, ".,!?;:'\"()[]{}"), ".,!?;:'\"()[]{}")
			word = strings.ToLower(word)
			wordFrequency[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print word frequency
	for word, freq := range wordFrequency {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
