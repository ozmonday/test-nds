package main

import "fmt"

func main() {
	Solution(1200000, 12, 0.01)
}

func Solution(pokok float64, tenor int, bungaPerBulan float64) {

	fmt.Printf("Pokok : %.2f\n", pokok)
	fmt.Printf("Tenor : %d\n", tenor)
	fmt.Printf("Bunga: %.2f\n", bungaPerBulan*100)

	pb := pokok / float64(tenor)
	bunga := (pokok / float64(tenor)) * bungaPerBulan
	fmt.Printf("Pinjaman per bulan : %.2f\n", pokok/float64(tenor))
	fmt.Printf("Bunga per bulan: %.1f\n", bunga)
	fmt.Printf("Total pinjaman : %.1f\n", pb+bunga)
	fmt.Println("Ags # pinjaman # bunga # ciclan per bulan # sisa pinjaman")

	for i := 1; i <= tenor; i++ {
		fmt.Printf("%d # %d # %.1f # %.1f # %.1f \n", i, int(pb), bunga, pb+bunga, (pb+bunga)*float64(tenor)-(float64(i-1)*(pb+bunga)))
	}

}
