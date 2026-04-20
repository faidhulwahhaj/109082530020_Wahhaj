package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, idx, target int
	fmt.Scan(&n)

	data := make([]int, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
		sum += float64(data[i])
	}

	fmt.Println("Isi array:", data)

	fmt.Print("Indeks Ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", data[i])
	}
	fmt.Print("\nIndeks Genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", data[i])
	}

	fmt.Scan(&x, &target)
	count := 0
	fmt.Printf("\nIndeks Kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if x > 0 && i%x == 0 {
			fmt.Printf("%d ", data[i])
		}
		if data[i] == target {
			count++
		}
	}

	mean := sum / float64(n)
	var varSum float64
	for i := 0; i < n; i++ {
		selisih := float64(data[i]) - mean
		varSum = varSum + (selisih * selisih)
	}
	sd := math.Sqrt(varSum / float64(n))

	fmt.Printf("\nMean: %.2f\nSD: %.2f\n", mean, sd)
	fmt.Printf("Frekuensi %d: %d\n", target, count)

	fmt.Scan(&idx)
	for i := idx; i < n-1; i++ {
		data[i] = data[i+1]
	}
	data = data[:n-1]
	
	fmt.Println("Data Akhir:", data)
}