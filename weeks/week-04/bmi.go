// program BodyMassIndex
// kamus
// 	type Stats <
// 		Name : string
// 		Weight : float64
// 		Height : float64
// 		Bmi    : float64
// 	>
// 	stats: Stats
// algoritma
// 	input (stats.Name, stats.Weight, stats.Height)
//         output (stats.Name)
// 	output (stats.Weight)
// 	output (stats.Height)
// 	call output (stats.bmi)

// algoritma2
// 	stats.Wright / (stats.Height * stats.Height)
// endprogram

package main

import (
	"fmt"
	"time"
)

// penetapan tipe bentukan struct untuk menyimpan data dengan berbagai bentuk,
// in this case it's in the form of string, and float64
type Stats struct {
	Name   string
	Weight float64
	Height float64
	Bmi    float64
}

func main() {
	var stats Stats

	fmt.Println("BMI CALCULATOR PROGRAM")
	fmt.Println("======================")
	time.Sleep(1 * time.Second)

	fmt.Println("Input Your name, weight(kg), and height(m) below :")
	fmt.Scan(&stats.Name, &stats.Weight, &stats.Height)
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------")
	time.Sleep(1 * time.Second)
	fmt.Printf("\n")
	time.Sleep(1 * time.Second)

	//mengkalkulasi bmi dan hasilnya disimpan di stats.Bmi
	//memanggil hasil
	stats.Bmi = bmiOperation(stats)

	//berbagai perintah prints untuk menampilkan hasil
	fmt.Println("RESULT!")
	fmt.Println("======================")
	time.Sleep(1 * time.Second)
	fmt.Println("Informasi BMI: ")
	time.Sleep(1 * time.Second)
	fmt.Println("======================")
	fmt.Println("Name: ", stats.Name)
	fmt.Println("Weight (kg): ", stats.Weight)
	fmt.Println("Height (m): ", stats.Height)
	fmt.Printf("BMI: %.2f", stats.Bmi)
	fmt.Printf("\n")

}

// fungsi untuk kalkukasi dan penyimpanan data bmi
func bmiOperation(stats Stats) float64 {
	return stats.Weight / (stats.Height * stats.Height)

}
