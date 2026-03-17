package main

import "fmt"

func konversiMeterMil(meter float64) float64{
	return meter *0.00062137119 
}
func main() {
	var meter float64

	fmt.Scan(&meter)

	fmt.Printf("%.2f\n", konversiMeterMil(meter))

	return
	
}