package main

import (
	"fmt"
	"os"

	"guess/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <fileName>\nExample : go run . data.txt")
		return
	}

	fileName := os.Args[1]
	// Sample Data
	xs, ys, err := utils.Reader(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	utils.Output(xs, ys)
}
