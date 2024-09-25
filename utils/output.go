package utils

import (
	"fmt"

	"linear/stats"
)

func Output(xs, ys []float64) {
	// var guess, residual, x float64
	y,r, slope, intercept, err := stats.LinearRegression(ys)
	if err != nil {
		fmt.Printf("\033[033mLinear Regression Line:\033[0m		%.0f = %.6fx, + %.6f\n", y, slope, intercept)
		fmt.Printf("\033[033mPearson Correlation Coefficient:	\033[0m %v\n", err)
	}
	fmt.Printf("\033[033mLinear Regression Line:\033[0m		%.0f = %.6fx, + %.6f\n", y, slope, intercept)
	fmt.Printf("\033[033mPearson Correlation Coefficient:	\033[0m %10.10f\n", r)
}
