package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const minimumDiskon = 100000

type Item struct {
	Nama     string
	Harga    int
	Jumlah   int
	Subtotal int
}

func hitungSubtotal(harga int, jumlah int) int {
	return harga * jumlah
}

func hitungDiskon(total int) int {
	if total >= minimumDiskon {
		return total * 10 / 100
	}

	return 0
}

func hitungTotal(items []Item) int {
	total := 0
	for _, item := range items {
		total += item.Subtotal
	}

	return total
}

func bacaTeks(scanner *bufio.Scanner, label string) (string, error) {
	fmt.Print(label)
	if !scanner.Scan() {
		return "", fmt.Errorf("input teks tidak valid")
	}

	value := strings.TrimSpace(scanner.Text())
	if value == "" {
		return "", fmt.Errorf("input tidak boleh kosong")
	}

	return value, nil
}

func bacaAngkaPositif(scanner *bufio.Scanner, label string) (int, error) {
	fmt.Print(label)
	if !scanner.Scan() {
		return 0, fmt.Errorf("input angka tidak valid")
	}

	value, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0, fmt.Errorf("input angka tidak valid")
	}
	if value <= 0 {
		return 0, fmt.Errorf("input harus lebih dari 0")
	}

	return value, nil
}

func inputItem(scanner *bufio.Scanner, nomor int) (Item, error) {
	fmt.Printf("\nItem ke-%d\n", nomor)

	nama, err := bacaTeks(scanner, "Nama barang: ")
	if err != nil {
		return Item{}, err
	}

	harga, err := bacaAngkaPositif(scanner, "Harga barang: ")
	if err != nil {
		return Item{}, err
	}

	jumlah, err := bacaAngkaPositif(scanner, "Jumlah barang: ")
	if err != nil {
		return Item{}, err
	}

	var dummy string
	fmt.Println("ketuk enter untuk melanjutkan")
	fmt.Scanln(&dummy)

	return Item{
		Nama:     nama,
		Harga:    harga,
		Jumlah:   jumlah,
		Subtotal: hitungSubtotal(harga, jumlah),
	}, nil
}

func tampilkanKeranjang(items []Item) {
	if len(items) == 0 {
		fmt.Println("Keranjang masih kosong")
		return
	}

	fmt.Println("\n=== Keranjang ===")
	for _, item := range items {
		fmt.Printf("%s x%d @Rp%d = Rp%d\n", item.Nama, item.Jumlah, item.Harga, item.Subtotal)
	}
}

func cetakStruk(items []Item) {
	if len(items) == 0 {
		fmt.Println("Keranjang masih kosong")
		return
	}

	total := hitungTotal(items)
	diskon := hitungDiskon(total)
	totalBayar := total - diskon

	tampilkanKeranjang(items)
	fmt.Println("---------------------")
	fmt.Println("Total:", total)
	fmt.Println("Diskon:", diskon)
	fmt.Println("Total Bayar:", totalBayar)
}

func tampilkanMenu() {
	fmt.Println("\x1b[2J")
	fmt.Println("\x1b[H")
	fmt.Println("\n=== CLI Kasir Sederhana ===")
	fmt.Println("1. Tambah barang")
	fmt.Println("2. Lihat keranjang")
	fmt.Println("3. Cetak struk")
	fmt.Println("4. Reset keranjang")
	fmt.Println("5. Keluar")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	items := []Item{}

	for {
		tampilkanMenu()

		pilihan, err := bacaAngkaPositif(scanner, "Pilih menu: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch pilihan {
		case 1:
			item, err := inputItem(scanner, len(items)+1)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			items = append(items, item)
			fmt.Println("Barang berhasil ditambahkan")
		case 2:
			tampilkanKeranjang(items)
		case 3:
			cetakStruk(items)
		case 4:
			items = []Item{}
			fmt.Println("Keranjang berhasil direset")
		case 5:
			fmt.Println("Terima kasih")
			return
		default:
			fmt.Println("Error: menu tidak tersedia")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
