package main

import "fmt"

func main() {
	Solution(9, 12)
}

func Solution(jamMasuk int, jamKeluar int) {

	jam := jamKeluar - jamMasuk

	if jam == 1 {
		fmt.Printf("Jam Masuk %d Jam Keluar %d Biaya Sewa %d\n", jamMasuk, jamKeluar, 350000)
	} else if jam == 2 {
		fmt.Printf("Jam Masuk %d Jam Keluar %d Biaya Sewa %d\n", jamMasuk, jamKeluar, 500000)
	} else if jam > 2 && jam <= 8 {
		fmt.Printf("Jam Masuk %d Jam Keluar %d Biaya Sewa %d\n", jamMasuk, jamKeluar, 500000+(jam-2)*75000)
	} else {
		fmt.Printf("Jam Masuk %d Jam Keluar %d Biaya Sewa %d\n", jamMasuk, jamKeluar, 500000+(jam-2)*100000)
	}
}
