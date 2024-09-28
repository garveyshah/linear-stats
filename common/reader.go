package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reader(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opeing file %v\n: %v", fileName, err)
	}
	defer file.Close()

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	data := []float64{}

	for fileScan.Scan() {
		line := fileScan.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing float from line '%v': %v", line, err)
		}

		data = append(data, value)

		// Check if any error occurred during scanning
		if err := fileScan.Err(); err != nil {
			return nil, fmt.Errorf("error reader file %v: %v", fileName, err)
		}
	}
	return data, nil
}
