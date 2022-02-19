package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Add green colorcode around the text
func makeTextGreen(orig string) string {
	return fmt.Sprintf("\u001b[32m%s\u001b[00m", orig)
}

type mapper func(orig string) string

// create a mapper function that highlight the pattern as green text
func makePatternGreen(pat string) mapper {
	re := regexp.MustCompile(pat)

	return func(orig string) string {
		return fmt.Sprintf(re.ReplaceAllStringFunc(orig, makeTextGreen))
	}
}

func main() {
	var mappers []mapper

	if len(os.Args) == 1 { // no args, make the complete line green
		mappers = append(mappers, makeTextGreen)
	} else {
		for _, pat := range os.Args[1:] { // generate an array of mappers from the args
			mappers = append(mappers, makePatternGreen(pat))
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	// scan the lines and highligth with the mappers
	for scanner.Scan() {
		line := scanner.Text()
		for _, mapper := range mappers {
			line = mapper(line)
		}
		fmt.Println(line)
	}
}
