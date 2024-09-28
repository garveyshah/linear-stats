package stats

import (
	"math"
	"testing"

	"linear/common"
)

// Helper function for comparing floating-point numbers
func isAlmostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestLinearRegression(t *testing.T) {
	// Expected slope and intercept values for each dataset
	tt := []struct {
		input     string
		slope     float64
		intercept float64
		r         float64
		err       error
	}{
		{"../test_data/data1.txt", -8.742857, 153.857143, -0.5330331012, nil},
		{"../test_data/data2.txt", 17.500000, 68.500000, 0.9891584832, nil},
		{"../test_data/data3.txt", -2.151515, 150.781818, -0.1898733212, nil},
	}
	for i, tc := range tt {
		ys, err := common.Reader(tc.input)
		if err != nil {
			t.Fatalf("Failed to read data from %q: %v", tc.input, err)
		}
		// Linear regression calculation for the current dataset
		r, slope, intercept, err := LinearRegression(ys)
		if err != tc.err {
			t.Fatalf("Test for %q Failed (case %d) - wrong error type, got= %s, want= %s", tc.input, i, err, tc.err)
		}

		tolerance := 0.0001 // Defination of a tolerance for floating-point comparison
		if !isAlmostEqual(r, tc.r, tolerance) {
			t.Fatalf("Test for %q Failed  (case %d) - wrong coeffient, got= %f, want= %f", tc.input, i, r, tc.r)
		}
		if !isAlmostEqual(slope, tc.slope, tolerance) {
			t.Fatalf("Test for %q Failed  (case %d) - wrong slope, got= %f, want= %f", tc.input, i, slope, tc.slope)
		}

		if !isAlmostEqual(intercept, tc.intercept, tolerance) {
			t.Fatalf("Test for %q Failed (case %d) - wrong intercept, got= %f, want= %f", tc.input, i, intercept, tc.intercept)
		}
	}
}
