package main

import (
	"fmt"

	"linear/utils"
)

func main() {
	// fmt.Println("Usage : go run . <fileName>\nExample : go run . data.txt")

	// fileName := os.Args[1]
	// Sample Data
	xs, ys, err := utils.Reader()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	utils.Output(xs, ys)
}
