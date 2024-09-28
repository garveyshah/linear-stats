package main

import (
	"fmt"
	"os"

	"linear/common"
	"linear/stats"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <fileName>\nExample : go run . data.txt")
		return
	}

	fileName := os.Args[1]
	// Sample Data
	//var xs []float64
	ys, err := common.Reader(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	r, slope, intercept, err := stats.LinearRegression(ys)
	if err != nil {
		fmt.Printf("\033[033mLinear Regression Line:\033[0m		y = %.6fx, + %.6f\n", slope, intercept)
		fmt.Printf("\033[033mPearson Correlation Coefficient:	\033[0m %v\n", err)
	}
	fmt.Printf("\033[033mLinear Regression Line:\033[0m		y = %.6fx, + %.6f\n", slope, intercept)
	fmt.Printf("\033[033mPearson Correlation Coefficient:	\033[0m %10.10f\n", r)
}
