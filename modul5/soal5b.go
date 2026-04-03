package main

import "fmt"

func cetakBintang(n int) {
	if n <= 0 {
		return
	}
	cetakBintang(n - 1)
	fmt.Print("*")
}

func buatPola(n int) {
	if n <= 0 {
		return
	}

	buatPola(n - 1)
	
	cetakBintang(n)
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	if n > 0 {
		buatPola(n)
	} else {
		fmt.Println("Masukan harus lebih besar dari 0")
	}
}