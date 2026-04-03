package main

import "fmt"

func fibonacci(n int) int {

	if n <= 1 {
		return n
	}
	// Langkah rekursif: Sn = Sn-1 + Sn-2
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var limit int = 10

	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", limit)
	fmt.Println("n  |  Sn")
	fmt.Println("----------")

	for i := 0; i <= limit; i++ {
		fmt.Printf("%-2d |  %d\n", i, fibonacci(i))
	}
}