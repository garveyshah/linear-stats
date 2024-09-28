package common

import (
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	tt := []struct {
		input string
		want []float64
		err   error
	}{
		{"../test_data/data.txt",  []float64{0.000000, 1.000000, 2.000000, 3.000000, 4.000000, 5.000000}, nil},
	}

	for _, tc := range tt {
	ys, err := Reader(tc.input)

		if err != tc.err {
			t.Fatalf("Test for %q Failed - error type mismatch, got= %q want %q", tc.input, err, tc.err)
		}

		if !reflect.DeepEqual(ys, tc.want) {
			t.Fatalf("Test for %q Failed - mismatch, got= \"%f\" want \"%f\"", tc.input, ys, tc.want)
		}
	}
}
