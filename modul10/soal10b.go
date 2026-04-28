package main

import (
	"fmt"
)

func main() {
	var x, y int
	var beratIkan [1000]float64
	
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalPerWadah []float64
	var rataRataPerWadah []float64

	currentTotal := 0.0
	countInWadah := 0

	for i := 0; i < x; i++ {
		currentTotal += beratIkan[i]
		countInWadah++

		if countInWadah == y || i == x-1 {
			totalPerWadah = append(totalPerWadah, currentTotal)
			rataRataPerWadah = append(rataRataPerWadah, currentTotal/float64(countInWadah))
			
			currentTotal = 0
			countInWadah = 0
		}
	}

	for i, val := range totalPerWadah {
		fmt.Printf("%.2f", val)
		if i < len(totalPerWadah)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	for i, val := range rataRataPerWadah {
		fmt.Printf("%.2f", val)
		if i < len(rataRataPerWadah)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}