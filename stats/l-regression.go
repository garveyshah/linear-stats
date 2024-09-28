package stats

import (
	"fmt"
	"math"
)

func LinearRegression(ys []float64) (r, slope, intercept float64, err error) {
	// Number of data points(len(xs))
	n := float64(len(ys))

	// Sums required for slope and intercept formulars
	var sumX, sumY, sumXY, sumX2, sumY2 float64
	var xs []float64

	for i := 0; i < len(ys); i++ {
		xs = append(xs, float64(i))
		sumX += xs[i]
		sumY += ys[i]
		sumXY += xs[i] * ys[i] // cross product
		sumX2 += xs[i] * xs[i]
		sumY2 += ys[i] * ys[i]
	}

	// calculate the Pearson correlation coefficient
	numerator := n*(sumXY) - (sumX)*(sumY)
	denominator := math.Sqrt(((n * sumX2) - sumX*sumX) * (n*sumY2 - sumY*sumY))

	if denominator == 0.0 {
		return 0.0, slope, intercept, fmt.Errorf("undefined (denominator is zero)")
	}
	r = numerator / denominator

	// calculate slope (b1)
	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)

	// Calculate intercept (b0)
	intercept = (sumY - slope*sumX) / n
	return r, slope, intercept, nil
}
