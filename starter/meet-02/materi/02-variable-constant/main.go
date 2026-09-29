package main

import "fmt"
import "strconv"
func main() {
	// var nama string = "Siti"
	// var umur int = 21
	const namaKelas = "Go Programming"
	nama := "Siti"
	umur := 21
	tinggi := 165.5
	aktif := true

	umur,_ = strconv.Atoi("44")	
	fmt.Println("Kelas:", namaKelas)
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Tinggi:", tinggi)
	fmt.Println("Aktif:", aktif)
}
