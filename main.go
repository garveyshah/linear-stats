package main

import (
	"fmt"
	"guess/functions"
	"log"
)

func main() {
	filename := "/home/oumouma/guess-it-1/data.txt"
	lines, err := functions.ReadFileLines(filename)
	if err != nil {
		log.Fatalf("Error reading %s: %v", filename, err)
	}
	fmt.Println("Finished reading the file.")
	fmt.Println("LInes read from file:", lines)

}
