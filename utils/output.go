package utils

import (
	"fmt"

	"linear/stats"
)

func Output(xs, ys []float64) {
	var guess, upperB, lowerB, residual, x, stdError float64
	slope, intercept := stats.LinearRegression(ys)
	var residsArray, guessArray []float64

	fmt.Printf("\033[032mSlope:\033[0m     %.2f\n", slope)
	fmt.Printf("\033[032mIntercept:\033[0m %.2f\n\n", intercept)

	for i, num := range xs {
		x, guess, residual = stats.Prectidor(num, ys[i], slope, intercept)

		residsArray = append(residsArray, residual)
		guessArray = append(guessArray, guess)

		stdError = stats.StandardError(residsArray)
		upperB, lowerB = stats.CalculateRange(guessArray, stdError)

		fmt.Printf("\033[032my:\033[0m	%.2f\n", ys[i])
		fmt.Printf("\033[032mGuess:\033[0m	%.2f\n", guess)
		fmt.Printf("\n\033[032mx:	\033[0m %.2f\n\033[032mUpperB:\033[0m %.2f\n\033[032mLowerB:\033[0m %.2f\n\033[032mResidual:\033[0m %.2f\n\n", x, upperB, lowerB, residual)

	}
}
