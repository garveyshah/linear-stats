package stats

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
