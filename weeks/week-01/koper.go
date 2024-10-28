// If jumlahKoper more than zero
// REPEAT
// Put suitcase in the scale
// store result in scale1
// Then subtract 1 from jumlahKoper
// Then put the next suitcase in the scale
// store result in scale2
// add scale 1 with scale2 then
// store result in weight

// UNTIL jumlahKoper equals zero
// 	endResult is weight divided with jumlahKoper

// Else end the program.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var jumlahPenumpang int

	fmt.Println("Berapa jumlah koper penumpang yang akan anda hitung :")
	fmt.Scanln(&jumlahPenumpang)

	if jumlahPenumpang <= 0 {
		fmt.Println("Jumlah koper harus lebih dari nol.")
		return
	}

	numberOfRandoms := jumlahPenumpang

	time.Sleep(1 * time.Second)
	fmt.Println(".........")
	time.Sleep(1 * time.Second)
	fmt.Println("processing")
	time.Sleep(1 * time.Second)
	fmt.Println(".........")
	time.Sleep(1 * time.Second)
	fmt.Println(".........")
	time.Sleep(1 * time.Second)
	fmt.Println(".........")
	time.Sleep(1 * time.Second)

	randomNumbers := fungsiRandomInts(numberOfRandoms)
	fmt.Println("Menghitung berat koper-koper dalam timbangan.", randomNumbers)

	total, average := fungsiRataRata(randomNumbers)
	fmt.Printf("Total berat koper anda andalah: %d\n", total)
	fmt.Printf("Rata-rata berat koper anda adalah: %2.f\n", average)
}

func fungsiRandomInts(n int) []int {
	randoms := make([]int, n)
	for i := 0; i < n; i++ {
		randoms[i] = rand.Intn(100) + 1
	}
	return randoms
}

func fungsiRataRata(numbers []int) (int, float64) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))
	return sum, average
}
