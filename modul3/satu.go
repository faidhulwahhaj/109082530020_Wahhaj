package main

import (
	"fmt"
)

func factorial(n int) int64 {
	var res int64 = 1
	for i := 1; i <= n; i++ {
		res *= int64(i)
	}
	return res
}

func permutation(n, r int) int64 {

	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int64 {
	
	return permutation(n, r) / factorial(r)
}

func main() {
	var a, b, c, d int

	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		return
	}


	fmt.Printf("%d %d\n", permutation(a, c), combination(a, c))

	fmt.Printf("%d %d\n", permutation(b, d), combination(b, d))
}