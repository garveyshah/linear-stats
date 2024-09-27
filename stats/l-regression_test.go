package stats

import (
	"math"
	"testing"

	"linear/utils"
)

// Helper function for comparing floating-point numbers
func isAlmostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestLinearRegression(t *testing.T) {
	// Files containing test data
	testData := []string{"../test_data/data1.txt", "../test_data/data2.txt", "../test_data/data3.txt"}

	for _, dataFile := range testData {
		ys, err := utils.Reader(dataFile)
		if err != nil {
			t.Fatalf("Failed to read data from %q: %v", dataFile, err)
		}

		// Expected slope and intercept values for each dataset
		tt := []struct {
			slope     float64
			intercept float64
			r         float64
			err       error
		}{
			{-0.999247, 552.465663, -0.9986787077, nil},
			// {0.8, 1.2},
			// {-0.4, 3.5},
		}

		// Linear regression calculation for the current dataset
		r, slope, intercept, err := LinearRegression(ys)

		for i, tc := range tt {
			tolerance := 0.0001 // Defination of a tolerance for floating-point comparison

			if !isAlmostEqual(r, tc.slope, tolerance) {
				t.Fatalf("Test for %q")
			}

			if !isAlmostEqual(slope, tc.slope, tolerance) {
				t.Fatalf("Test for %q Failed  (case %d) - wrong slope, got= %f, want= %f", dataFile, i, slope, tc.slope)
			}

			if intercept != tc.intercept {
				t.Fatalf("Test for %q Failed (case %d) - wrong intercept, got= %f, want= %f", dataFile, i, intercept, tc.intercept)
			}
		}
	}
}
