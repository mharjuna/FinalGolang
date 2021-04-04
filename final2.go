package main

import (
	"fmt"
)

type barang struct {
	NamaBrg  string
	HargaBrg int
}

type belanja struct {
	JumlahBrg int
}

func main() {
	fmt.Println("DAFTAR BELANJAAN")
	var a, b, c, d, e int
	fmt.Println("1. Sabun Cuci Piring = 50.000")
	brg1 := barang{
		HargaBrg: 50000}
	fmt.Print("Masukkan Jumlah Belanjaan = ")
	fmt.Scanln(&a)
	fmt.Println("2. Sabun Mandi = 25.000")
	brg2 := barang{
		HargaBrg: 25000}
	fmt.Print("Masukkan Jumlah Belanjaan = ")
	fmt.Scanln(&b)
	fmt.Println("3. Shampoo = 35.000")
	brg3 := barang{
		HargaBrg: 35000}
	fmt.Print("Masukkan Jumlah Belanjaan = ")
	fmt.Scanln(&c)
	fmt.Println("4. Pasta Gigi = 10.000")
	brg4 := barang{
		HargaBrg: 10000}
	fmt.Print("Masukkan Jumlah Belanjaan = ")
	fmt.Scanln(&d)
	fmt.Println("5. Tissue = 7.500")
	brg5 := barang{
		HargaBrg: 7500}
	fmt.Print("Masukkan Jumlah Belanjaan = ")
	fmt.Scanln(&e)

	fmt.Println("==================================================")
	fmt.Println("DETAIL BELANJAAN")

	fmt.Println("Sabun Cuci Piring = ", brg1.HargaBrg*a)
	fmt.Println("Sabun Mandi = ", brg2.HargaBrg*b)
	fmt.Println("Shampoo = ", brg3.HargaBrg*c)
	fmt.Println("Pasta Gigi = ", brg4.HargaBrg*d)
	fmt.Println("Tissue = ", brg5.HargaBrg*e)

	total := (brg1.HargaBrg * a) + (brg2.HargaBrg * b) + (brg3.HargaBrg * c) + (brg4.HargaBrg * d) + (brg5.HargaBrg * e)
	fmt.Println("Total keseluruhan belanjaan adalah = ", total)

}
