# Linear Stats

**Linear Stats** is a Go project that calcultes linear linear regression and statistical metrics like the slope, intercept, and Pearson correlation coefficient from data files.

## Table of Contents
+ [Features](#features)
+ [Installation](#installation)
+ [Usage](#usage)
   + [Running the program](#running-the-program)
+ [Testing](#testing)
+ [Project Structure](#project-structure)
+ [License](#licence)

## Features
+ **Linear Regression**: Calculates the line of best fit for a set of data points.
+ **Pearson Correlation Coefficient**: Measure the strength of the linear relationship between two variables.
+ **Flexible Input**: Read data from text files to perform calculations.

## Installation
### Prerequisites
+ **Go** installed. You can download Go [here](https://golang.org/doc/install)
### Steps
1.Clone the project: 
```bash
git clone https://learn.zone01kisumu.ke/git/oumouma/linear-stats.git
```
2. Navigate to the project directory:
```bash
cd linear-stats
```
## Usage
### Running Linear Regression
You can use the `LinearRegression` function from the `stats` package to calculate the regression:
```go
package main

import (
    "fmt"
    "linear/common"
    "linear/stats"
)

func main() {
    ys, err := common.Reader("../test_data/data1.txt")
    if err != nil {
        fmt.Println("Error reading data:", err)
        return
    }

    r, slope, intercept, err := stats.LinearRegression(ys)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("y = %.6fx + %.6f\n", slope, intercept)
    fmt.Printf("Pearson Correlation: %.10f\n", r)
}
```
## Running the program
Use Go to run the program:
```bash
go run main.go
```

```bash
go run . data.txt
```
## Testing
Unit tests are included in the project. You can run the using:
```bash
go test ./...
```

## Project Structure
```bash
linear-stats/
│
├── stats/           # Linear regression code and tests
├── common/          # Helper functions for reading data
├── test_data/       # Sample datasets
├── main.go           # The entry point of the program
├── go.mod           # Go module file
└── README.md        # Project documentation
```

## Author

**Garvey Shah**

![Profile Picture](https://avatars.githubusercontent.com/u/162585009?v=4)

- GitHub: [garveyshah](https://github.com/garveyshah)

[Ouma Ouma]()

## Licence
This project is licensed under the [MIT License](/LICENSE).


