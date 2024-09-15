package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reader(fileName string) ([]float64, []float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opeing file %v\n: %v", fileName, err)
	}
	defer file.Close()

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	var ys, xs []float64

	for fileScan.Scan() {
		line := fileScan.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing float from line '%v': %v", line, err)
		}

		ys = append(ys, value)

		// Check if any error occurred during scanning
		if err := fileScan.Err(); err != nil {
			return nil, nil, fmt.Errorf("error reader file %v: %v", fileName, err)
		}
	}

	for i := 0; i < len(ys); i++ {
		xs = append(xs, float64(i))
	}

	return xs, ys, nil
}
