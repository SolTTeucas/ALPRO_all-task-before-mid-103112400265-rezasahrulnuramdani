package main

import (
	"fmt"
)

func main() {
	var semester, eprt int16
	fmt.Println("Cek apakah anda akan lulus dengan cumlaude!")
	fmt.Println("Masukkan banyak semester lulus dan score eprt anda : ")
	fmt.Scan(&semester, &eprt)
	fmt.Println("Data yang anda input adalah ", semester, "dan", eprt)

	a := semester <= 8
	b := eprt >= 500
	negA := semester > 8
	negB := eprt < 500
	if a && b {
		fmt.Println("True | Mahasiswa lulus cumlaude dengan kuliah selama ")
		fmt.Println(semester, "semester dan EPrT", eprt)
	} else if negA && negB {
		fmt.Println("False | Tidak cumlaude karena anda huliah hingga, ", semester, "dengan score EPrT", eprt)
	} else if !a {
		fmt.Println("False | Tidak cumlaude karena kuliah hingga ", semester, "semester")
	} else if !b {
		fmt.Println("False | Tidak cumlaude karena score EPrT anda adalah ", eprt)
	}
}
