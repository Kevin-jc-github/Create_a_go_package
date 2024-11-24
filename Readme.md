# GO_Creating_A_Package

## Overview
This repository contains a Go program that demonstrates how to use the `TrimmedMean` package, which calculates the trimmed mean of a given slice of numeric values. The program imports the `TrimmedMean` package from a separate repository and applies it to a data set of 100 numbers to illustrate its functionality.

The main program includes the following file:

- **`main.go`**: A Go program that imports the `TrimmedMean` package and uses it to calculate the trimmed mean for an example data set.

## File Descriptions

### `main.go`
This file contains a simple Go program that imports and uses the `TrimmedMean` function from the `GO_trimmedmean_Package_By_Jiancong_Zhu` repository. It demonstrates how to compute the trimmed mean for a given slice of numeric values with specified trimming proportions.

#### Functionality
- The `main.go` script uses the `TrimmedMean` function from the package to calculate the trimmed mean for a set of numeric values.
- It provides a practical example to show how the package can be used in real scenarios.

## Installation and Setup

### Step 1: Clone the Main Program Repository
First, clone the repository containing the main program:

```bash
git clone https://github.com/Kevin-jc-github/GO_MainProgram_Using_TrimmedMean_Package.git
```

### Step 2: Install the TrimmedMean Package
To use the `TrimmedMean` package in `main.go`, you need to install the package from its GitHub repository. Run the following command:

```bash
go get github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu
```

### Step 3: Initialize the Module
Navigate to the main program directory and initialize the module:

```bash
cd GO_MainProgram_Using_TrimmedMean_Package
go mod init github.com/Kevin-jc-github/GO_MainProgram_Using_TrimmedMean_Package
```

### Step 4: Import the TrimmedMean Package
In `main.go`, you will need to import the `TrimmedMean` package to use its function. The import statement should look like this:

```go
import (
    "fmt"
    "github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu"
)
```

## Example Usage
Here's an example of how the `TrimmedMean` function is used in `main.go`:

```go
package main

import (
    "fmt"
    "github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu"
)

func main() {
    // Create a slice of 100 float64 numbers from 1 to 100
    data := make([]float64, 100)
    for i := 0; i < 100; i++ {
        data[i] = float64(i + 1)
    }

    // Calculate the trimmed mean with a 0.05 trimming proportion
    mean, err := trimmedmean.TrimmedMean(data, 0.05)
    if err != nil {
        fmt.Printf("Error calculating trimmed mean: %v\n", err)
        return
    }

    fmt.Printf("Trimmed mean: %.2f\n", mean)
}
```

## Running the Main Program
To run the `main.go` program, use the following command:

```bash
go run main.go
```

This will compute the trimmed mean of the data set and display it in the terminal.
