package main

import (
	"fmt"

	trimmedmean "github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu"
)

func main() {
	data := make([]float64, 100)
	for i := 0; i < 100; i++ {
		data[i] = float64(i + 1) // Fill with numbers 1 to 100
	}

	mean, err := trimmedmean.TrimmedMean(data, 0.05)
	if err != nil {
		fmt.Printf("Error calculating trimmed mean: %v\n", err)
		return
	}

	fmt.Printf("Trimmed mean: %.2f\n", mean)
}
