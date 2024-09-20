package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reader() ([]float64, []float64, error) {
	fileScan := bufio.NewScanner(os.Stdin)
	fileScan.Split(bufio.ScanLines)
	var ys, xs []float64

	fmt.Println("Enter value (or 'exit' to finish): ")
	for {
		if !fileScan.Scan() {
			break
		}

		line := fileScan.Text()

		// if line == "exit" {
		// 	break
		// }

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing float from line '%v': %v", line, err)
		}

		ys = append(ys, value)

		if len(ys) == 2 {
			break
		}

		// Check if any error occurred during scanning
		if err := fileScan.Err(); err != nil {
			return nil, nil, fmt.Errorf("error readering input: %v", err)
		}
	}

	for i := 0; i < len(ys); i++ {
		xs = append(xs, float64(i))
	}

	return xs, ys, nil
}
