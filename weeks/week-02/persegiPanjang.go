// start
// var luas, keliling, p, l dalam float32

// input dan rekam p, l
// operasikan p, l dalam operasi yang dinamakan luas
// operasikan p, l dalam operasi yang dinamakan keliling
// tampilkan hasil luas dan keliling

// end

package main

import (
	"fmt"
)

func main() {
	var luas, keliling, p, l float32

	fmt.Print("Masukkan inputan p dan l persegi panjang anda:")
	fmt.Scan(&p, &l)

	luas = p * l
	fmt.Println("Luas persegi panjang adalah = ", luas)
	keliling = (2 * p) + (2 * l)
	fmt.Println("Luas persegi panjang adalah = ", keliling)
}
