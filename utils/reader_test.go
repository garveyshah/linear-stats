package utils

import (
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	tt := []struct {
		input string
		want  []float64
		err   error
	}{
		{"../test_data/data.txt", []float64{189.000000, 113.000000, 121.000000, 114.000000, 145.000000, 110.000000}, nil},
	}

	for _, tc := range tt {
		data, err := Reader(tc.input)

		if err != tc.err {
			t.Fatalf("Test for %q Failed - error type mismatch, got= %q want %q", tc.input, err, tc.err)
		}

		if !reflect.DeepEqual(data, tc.want) {
			t.Fatalf("Test for %q Failed - error type mismatch, got= \"%f\" want \"%f\"", tc.input, data, tc.want)
		}
	}
}
