package main

import (
	"fmt"
	"os"

	"guess/stats"
	"guess/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <fileName>\nExample : go run . data.txt")
		return
	}

	fileName := os.Args[1]
	// Sample Data
	data, err := utils.Reader(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Data :- ", data, "\n")
	// Perform linear regression
	slope, intercept := stats.LinearRegression(data)

	// Output results
	fmt.Printf("Slope: %.2f\n", slope)
	fmt.Printf("Intercept: %.2f\n", intercept)
}
