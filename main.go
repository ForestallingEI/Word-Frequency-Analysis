package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Read the text from a file
	file, err := os.Open("Tarzan-triumphant.txt") // Replace with your input file
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	wordFrequency := countWord(file)
	// < Print word frequency #map# >
	/*
		for word, freq := range wordFrequency {
			fmt.Printf("%s: %d\n", word, freq)
		}
	*/

	totalCountedWords := len(wordFrequency)
	defer fmt.Println("Total words: ", totalCountedWords)

	// < Print words only. alphabetical-order #map# >
		/*
		sortedWords := alphabetical(wordFrequency)
		for _,w := range sortedWords {
			fmt.Println(w)
		}
		*/


	// < Convert from map to slice, for counting-order >
	wordcounters := converter(wordFrequency)
	// < Print by count $slice$ >
	
		sort.Slice(wordcounters, func(i,j int)bool{
			return wordcounters[i].Counter < wordcounters[j].Counter
		})
		fmt.Println("By count (Ascending): ", wordcounters)
	
	// < Print alphabetical-order $slice$ >
	/*
	sort.Slice(wordcounters, func(i, j int) bool {
		return wordcounters[i].Word < wordcounters[j].Word
	})
	fmt.Println("By word: ", wordcounters)
	*/

}

func countWord(file io.Reader) map[string]int {
	// Create a map to store word frequency
	wordFrequency := make(map[string]int)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // Split line into words

		// Count word frequency
		for _, word := range words {
			// Remove symbols as much as possible and convert to lowercase for better counting
			word = strings.Trim(word,".\";,?-_‘ “•'()* ")
			word = strings.ToLower(word)
			wordFrequency[word]++
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading file:", err)
		}
	}
	return wordFrequency
}

func alphabetical(wordFrequency map[string]int) []string {
	sortedWords := make([]string, 0, len(wordFrequency))
	for word := range wordFrequency {
		sortedWords = append(sortedWords, word)
	}
	sort.Strings(sortedWords)
	return sortedWords
}

type WordCounter struct {
	Word    string
	Counter int
}

type SliceCounter []WordCounter

func converter(m map[string]int) SliceCounter {
	wordcounters := make(SliceCounter, len(m))
	i := 0
	for word, counter := range m {
		wordcounters[i] = WordCounter{word, counter}
		i++
	}
	return wordcounters
}
