package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

func countWords(filename string, wg *sync.WaitGroup, results chan<- map[string]int) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file info for %s: %v\n", filename, err)
		return
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	wordCount := make(map[string]int)
	words := regexp.MustCompile(`[a-zA-Z]+`).FindAllString(string(data), -1)

	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	results <- wordCount
}

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	var wg sync.WaitGroup
	results := make(chan map[string]int)

	for _, file := range files {
		wg.Add(1)
		go countWords(file, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalWordCount := make(map[string]int)
	for wordCount := range results {
		for word, count := range wordCount {
			totalWordCount[word] += count
		}
	}

	for word, count := range totalWordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
