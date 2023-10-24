package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type kv struct {
	key   string
	value int
}

type kvSlice []kv

func main() {

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCounts[scanner.Text()]++
	}

	fmt.Println("There is", len(wordCounts), "unique words!")

	var wordsFreq kvSlice
	for k, v := range wordCounts {
		wordsFreq = append(wordsFreq, kv{k, v})
	}

	sort.SliceStable(wordsFreq, func(i, j int) bool {
		return wordsFreq[i].value > wordsFreq[j].value
	})

	for _, v := range wordsFreq[:5] {
		fmt.Printf("%v => %v\n", v.key, v.value)
	}
}
