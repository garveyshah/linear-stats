package stats

import (
	"fmt"
	"math"
)

// LinearRegression calculates the linear regression parameters (slope and intercept)
// and the Pearson correlation coefficient for a given set of y-values.
// It returns the correlation coefficient (r), the slope (b1), the intercept (b0),
// and any error encountered during the calculation.
func LinearRegression(ys []float64) (r, slope, intercept float64, err error) {
	// Number of data points (n)
	n := float64(len(ys))

	// Initialize sums required for the slope and intercept formulas
	var sumX, sumY, sumXY, sumX2, sumY2 float64
	var xs []float64 // Slice to hold x-values corresponding to y-values

	// Calculate sums and prepare x-values
	for i := 0; i < len(ys); i++ {
		xs = append(xs, float64(i)) // Assign x-values (0, 1, 2, ...)
		sumX += xs[i]               // Sum of x-values
		sumY += ys[i]               // Sum of y-values
		sumXY += xs[i] * ys[i]      // Sum of cross products (x * y)
		sumX2 += xs[i] * xs[i]      // Sum of squares of x-values
		sumY2 += ys[i] * ys[i]      // Sum of squares of y-values
	}

	// Calculate the numerator and denominator for the Pearson correlation coefficient
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))

	// Check for undefined correlation (denominator is zero)
	if denominator == 0.0 {
		return 0.0, slope, intercept, fmt.Errorf("undefined (denominator is zero)")
	}
	r = numerator / denominator // Calculate the Pearson correlation coefficient

	// Calculate the slope (b1) using the formula
	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)

	// Calculate the intercept (b0) using the formula
	intercept = (sumY - slope*sumX) / n

	// Return the correlation coefficient, slope, intercept, and no error
	return r, slope, intercept, nil
}
