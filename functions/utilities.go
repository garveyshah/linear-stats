package functions

import (
	"bufio"
	"os"
)

// Function to check if a file exists at the specified 'filename'.
func CheckFileExits(filename string) bool {
	_,  err := os.Stat(filename)
	return !os.IsNotExist(err)
}


// ReadFileLines reads all lines from a file and returns them as a slice of string.
func ReadFileLines(filename string) ([]string, error) {
	
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// // Function prints the output data in a the specific desirable format.
// func PrintResults() {
// 	fmt.Printf("%d --> the standard input\n",lines[i])
// 	fmt.Printf("%d %d  --> the range for the next, input in this case for the number %d",lines[i])

// }