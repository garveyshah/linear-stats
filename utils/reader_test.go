package utils

import (
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	tt := []struct {
		input string
		want1 []float64
		want2 []float64
		err   error
	}{
		{"../test_data/data.txt", []float64{0, 1, 2, 3, 4, 5}, []float64{0.000000, 1.000000, 2.000000, 3.000000, 4.000000, 6.000000}, nil},
	}

	for _, tc := range tt {
	ys, _ := Reader(tc.input)
	var xs []float64

		// if err != tc.err {
		// 	t.Fatalf("Test for %q Failed - error type mismatch, got= %q want %q", tc.input, err, tc.err)
		// }

		if !reflect.DeepEqual(xs, tc.want1) {
			t.Fatalf("Test for %q Failed - error type mismatch, got= \"%f\" want \"%f\"", tc.input, xs, tc.want1)
		}

		if !reflect.DeepEqual(ys, tc.want2) {
			t.Fatalf("Test for %q Failed - error type mismatch, got= \"%f\" want \"%f\"", tc.input, ys, tc.want2)
		}
	}
}
