package main

import "fmt"

func main() {
	namaBarang := [3]string{"Buku", "Pulpen", "Tas"}
	//namaBarang = append(namaBarang, "Penghapus")
	//namaBarang = append(namaBarang, "123")

	for _, barang := range namaBarang {
		fmt.Println("Barang:", barang)
	}

	stok := map[string]int{
		"Buku":   10,
		"Pulpen": 25,
		"Tas":    5,
	}
	stok2 := map[string][]string{
		"Buku":   {"judul 1", "judul 2"},
		"Pulpen": {"kenko", "joyco"},
	}
	fmt.Println("Stok Buku:", stok["Buku"])
	stok["Pulpen"] = 20
	fmt.Println("Stok Pulpen:", stok["Pulpen"])

	fmt.Println(stok2)
}
