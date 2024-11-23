## Repository: [TrimmedMeanPackage](https://github.com/BellowAverage/TrimmedMeanPackage)

### README.md

```markdown
# TrimmedMeanPackage

A Go package for computing the trimmed mean of slices of integers and floats with support for symmetric and asymmetric trimming.

## Overview

`TrimmedMeanPackage` provides functions to calculate the trimmed mean, a robust measure of central tendency, by removing a specified proportion of the smallest and largest values from a dataset before computing the mean.

### Features

- Compute trimmed means for slices of integers and floats.
- Support symmetric trimming (equal proportions from both ends).
- Support asymmetric trimming (different proportions from lower and upper ends).
- Error handling for invalid inputs.

## Installation

To install the package, run:

```bash
go get github.com/BellowAverage/TrimmedMeanPackage
```

This command will download the package and add it to your Go workspace.

## Usage

Import the package in your Go code:

```go
import "github.com/BellowAverage/TrimmedMeanPackage/trimmedmean"
```

### Functions

#### TrimmedMeanInt

Calculates the trimmed mean of a slice of integers.

```go
func TrimmedMeanInt(data []int, trimLower float64, trimUpper ...float64) (float64, error)
```

- `data`: Slice of integers.
- `trimLower`: Proportion to trim from the lower end (0 ≤ trimLower < 0.5).
- `trimUpper`: (Optional) Proportion to trim from the upper end. If omitted, trimming is symmetric using `trimLower`.

#### TrimmedMeanFloat

Calculates the trimmed mean of a slice of floats.

```go
func TrimmedMeanFloat(data []float64, trimLower float64, trimUpper ...float64) (float64, error)
```

- Parameters are the same as `TrimmedMeanInt`, but for floats.

### Examples

#### Symmetric Trimming

```go
package main

import (
    "fmt"
    "github.com/BellowAverage/TrimmedMeanPackage/trimmedmean"
)

func main() {
    data := []float64{2.5, 3.7, 4.1, 5.6, 100.0}
    mean, err := trimmedmean.TrimmedMeanFloat(data, 0.2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Symmetric Trimmed Mean: %.2f\n", mean)
}
```

#### Asymmetric Trimming

```go
package main

import (
    "fmt"
    "github.com/BellowAverage/TrimmedMeanPackage/trimmedmean"
)

func main() {
    data := []int{10, 20, 30, 40, 1000}
    mean, err := trimmedmean.TrimmedMeanInt(data, 0.1, 0.3)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Asymmetric Trimmed Mean: %.2f\n", mean)
}
```

## Error Handling

The functions return an error in the following cases:

- The data slice is empty.
- Trimming proportions are not within the range [0, 0.5).
- The sum of `trimLower` and `trimUpper` proportions removes all data points.

Always check for errors when calling these functions to ensure reliable results.