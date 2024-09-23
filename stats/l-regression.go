package stats

import "math"

func LinearRegression(ys []float64) (slope, intercept float64) {
	// Number of data points(len(xs))
	n := float64(len(ys))

	// Sums required for slope and intercept formukars
	var sumX, sumY, sumXY, sumX2 float64
	var xs []float64
	for i := 0; i < len(ys); i++ {
		xs = append(xs, float64(i))
		sumX += xs[i]
		sumY += ys[i]
		sumXY += xs[i] * ys[i]
		sumX2 += xs[i] * xs[i]
	}

	// calculate slope (b1)
	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)

	// Calculate intercept (b0)
	intercept = (sumY - slope*sumX) / n
	return slope, intercept
}

// Predictor makes predictions and calculates residuals
func Prectidor(x, y, slope, intercept float64) (float64, float64, float64) {
	guess := (slope * x) + intercept
	residual := y - guess
	return x, guess, residual
}

// Calculate returns the upper and lower bounds for predictions
func CalculateRange(guessArray []float64, stdError float64) (float64, float64) {
	var upperB, lowerB float64

	for _, guess := range guessArray {
		upperB = guess + (stdError * 3)
		lowerB = guess - (stdError * 3)
	}
	return upperB, lowerB
}

// StandardError calculates the standard error of the residuals
func StandardError(residuals []float64) float64 {
	n := len(residuals)
	var ssd float64

	if n <= 2 {
		return math.NaN()
	}

	for _, num := range residuals {
		ssd += num * num
	}
	stdError := ssd / float64(n-2)
	return float64(math.Sqrt(stdError))
}
