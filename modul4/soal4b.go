package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 { 
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenangNama string
	var soalSelesai, totalSkor int
	var maxSoal, minSkor int

	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&soalSelesai, &totalSkor)

		if soalSelesai > maxSoal || (soalSelesai == maxSoal && totalSkor < minSkor) {
			maxSoal = soalSelesai
			minSkor = totalSkor
			pemenangNama = nama
		}
	}

	if pemenangNama != "" {
		fmt.Printf("%s %d %d\n", pemenangNama, maxSoal, minSkor)
	}
}