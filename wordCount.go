package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("What file should we open:")
	var fileName string
	fmt.Scan(&fileName)

	byteArray, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Couldn't open file:", err)
		return
	}

	allText := string(byteArray)
	allLines := strings.Split(allText, "\n")
	wordCounts := countWords(allLines)
	reportResults(wordCounts)
}

func countWords(lines []string) map[string]int {
	wordCounts := make(map[string]int)
	regEx, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		line = regEx.ReplaceAllString(line, " ")
		words := strings.Split(line, " ")
		for _, currentWord := range words {
			count, isThere := wordCounts[currentWord]
			if !isThere {
				wordCounts[currentWord] = 1
			} else {
				wordCounts[currentWord] = count + 1
			}
		}
	}
	return wordCounts
}

func reportResults(wordCounts map[string]int) {
	for word, count := range wordCounts {
		fmt.Println(word, ": ", count)
	}

}
