package utils

import (
	"fmt"

	"guess/stats"
)

func Output(xs, ys []float64) {
	var guess, upperB, lowerB, residual, x, stdError float64
	slope, intercept := stats.LinearRegression( ys)
	var residsArray, guessArray []float64

	fmt.Printf("Slope:     %.2f\n", slope)
	fmt.Printf("Intercept: %.2f\n\n", intercept)

	for i, num := range xs {
		x, guess, residual = stats.Prectidor(num, slope, intercept)

		residsArray = append(residsArray, residual)
		guessArray = append(guessArray, guess)

		stdError = stats.StandardError(residsArray)
		upperB, lowerB = stats.CalculateRange(guessArray, stdError)

		fmt.Printf("y:	%.2f\n", ys[i])
		fmt.Printf("Guess:	%.2f\n", guess)
		fmt.Printf("\nx: %.2f\nUpperB: %.2f\nLowerB: %.2f\nResidual: %.2f\n\n", x, upperB, lowerB, residual)

	}

	
}
