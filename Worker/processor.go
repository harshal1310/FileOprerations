package worker

import (
	"TrueCallerAssignment/counter"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Job struct {
	FilePath string
}

// Result contains file name ,line count,wordfreq for reach file,error
type Result struct {
	FilePath  string
	LineCount int
	wordFreq  map[string]int
	Err       error
}

var wordFreqForAllFiles = make(map[string]int)

func workerFunc(fileChan <-chan string, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for filePath := range fileChan {
		lineCount, wordFreq, err := processFile(filePath)
		result <- Result{
			FilePath:  filePath,
			LineCount: lineCount,
			wordFreq:  wordFreq,
			Err:       err,
		}
	}
}

func processFile(filePath string) (int, map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()
	lineCount := 0
	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		words := strings.Fields(line)
		for _, word := range words {
			clean := strings.ToLower(word)
			wordFreq[clean]++
		}
	}
	return lineCount, wordFreq, scanner.Err()
}
func StartWorker(files []string, workerPool int) {
	fileChan := make(chan string, workerPool)
	results := make(chan Result, workerPool)
	var wg sync.WaitGroup

	for worker := 0; worker < workerPool; worker++ {
		wg.Add(1)
		go workerFunc(fileChan, results, &wg)
	}

	for _, file := range files {
		fileChan <- file
	}
	close(fileChan)
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Println("Error while reading the File", result.FilePath)
			continue
		}
		fmt.Printf("File name %s has %d lines\n", result.FilePath, result.LineCount)
		counter.MergerWordCount(wordFreqForAllFiles, result.wordFreq)

	}
	topN := 100
	topFrequentWords := counter.PrintTopWords(wordFreqForAllFiles, topN)
	size := len(topFrequentWords)
	if size < topN {
		fmt.Printf("List has only %d unique words\n", size)
	}
	fmt.Printf("Top %d freq words are:\n", topN)
	for _, wordCount := range topFrequentWords {
		fmt.Printf("%s: %d\n", wordCount.Word, wordCount.WordFreq)
	}
}
