package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(filename, oldWord, newWord string) ([]string, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		newLine := strings.ReplaceAll(line, oldWord, newWord)
		lines = append(lines, newLine)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}

func writeLines(filename string, lines []string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err = f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Filename, old word and new word are needed!")
		return
	}

	filename, oldWord, newWord := args[0], args[1], args[2]

	lines, err := readLines(filename, oldWord, newWord)
	check(err)

	writeLines(filename, lines)
	check(err)

	fmt.Println("Done!")
}
