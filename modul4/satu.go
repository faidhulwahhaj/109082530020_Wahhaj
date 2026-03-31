package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nrFact int
	factorial(n, &nFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / nrFact
}

func combination(n, r int, hasil *int) {
	var p, rFact int
	permutation(n, r, &p)
	factorial(r, &rFact)
	*hasil = p / rFact
}

func main() {
	var a, b, c, d int
	var resP, resC int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &resP)
	combination(a, c, &resC)
	fmt.Println(resP, resC)

	permutation(b, d, &resP)
	combination(b, d, &resC)
	fmt.Println(resP, resC)
}