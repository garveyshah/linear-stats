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
	// fmt.Println("xs := ", xs)
	return slope, intercept
}

// Function to calculate the range the next number may fall and predic the next number employing the linear regression model.
func Prectidor(x, slope, intercept float64) (float64, float64, float64) {
	guess := (slope * x) + intercept

	residual := x - guess

	return x, guess, residual
}

func CalculateRange(guessArray []float64, stdError float64) (float64, float64) {
	var upperB, lowerB float64

	for _, guess := range guessArray {
		upperB = guess - (stdError * 2)
		lowerB = guess + (stdError * 2)
	}
	return upperB, lowerB
}

// Function to calculate Variance of the residuals
func StandardError(residuals []float64) float64 {
	n := len(residuals)
	var ssd, stdError float64

	for _, num := range residuals {
		ssd += num * num
		stdError = ssd / float64(n-2)
	}
	return float64(math.Sqrt(stdError))
}
