package main

import (
	"fmt"
	"time"
)

func main() {
	var banyak, sisi int64

	fmt.Println("PROGRAM BIDANG PERSEGI")
	fmt.Println("======================")
	time.Sleep(1 * time.Second)

	fmt.Println("Masukkan jumlah persegi")
	fmt.Scan(&banyak)
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------")
	time.Sleep(1 * time.Second)
	fmt.Printf("\n")
	time.Sleep(1 * time.Second)

	//berbagai perintah prints untuk menampilkan hasil
	fmt.Println("RESULT!")

	for i := int64(0); i < banyak; i++ {
		fmt.Printf("Masukkan panjang sisi untuk persegi %d: ", i+1)
		fmt.Scan(&sisi) // Meminta input sisi untuk setiap persegi

		// Menghitung keliling dan luas
		hasilKeliling := 4 * sisi
		hasilLuas := sisi * sisi

		// Menampilkan hasil
		fmt.Printf("Persegi %d: Keliling = %d, Luas = %d\n", i+1, hasilKeliling, hasilLuas)
		time.Sleep(500 * time.Millisecond)
	}

}
