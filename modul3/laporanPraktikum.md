# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">Wahhaj - 109082530020</p>

## Unguided 

### 1. [Soal modul 2A]
#### satu.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Soal 1](output/outputsatu.png)
[penjelasan]
  Jadi kode itu digunakan untuk menggeser posisi nilai variabel

  ### 2. [Soal modul 2B]
#### soal2b.go

```go
package main

import "fmt"

func main() {
	

	var w1, w2, w3, w4 string

	hasilAkhir := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
	
		fmt.Scan(&w1, &w2, &w3, &w4)

	
		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
			hasilAkhir = false
		}
	}
	fmt.Printf("BERHASIL: %t\n", hasilAkhir)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Soal 2](output/outputsoal2b.png)
[penjelasan]
 Jadi kode itu digunakan untuk melakukan pengecekan string di dalam sebuah perulangan sebanyak 5 kali,jika ada satu yang tidak sama maka bisa dikatakan false.


### 3. [Soal modul 2C]
#### soal2c.go

```go
package main

import "fmt"

func main() {
	var beratGram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000


	biayaKg := kg * 10000


	var biayaSisa int
	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}

	totalBiaya := biayaKg + biayaSisa
	if kg > 10 {
		totalBiaya = biayaKg
	}

	
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Soal 3](output/outputsoal2c.png)
[penjelasan]
  Jadi kode tersebut digunakan untuk biaya pengiriman berdasarkan berat paket dalam sebuah gram.