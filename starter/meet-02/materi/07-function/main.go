package main

import "fmt"

func sapa(nama string) {
	fmt.Println("Halo", nama)
}

func tambah(a int, b int) int {
	return a + b
}

func hitungDiskon(total int, persen int) (diskon, hasil int) {
	return persen, total * persen / 100
}

func main() {
	sapa("Andi")

	hasil := tambah(10, 5)
	fmt.Println("10 + 5 =", hasil)

	diskon, hasil := hitungDiskon(100000, 10)
	fmt.Println("Diskon:", diskon, "Hasil:", hasil)
}
