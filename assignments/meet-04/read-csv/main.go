package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
const ( //color
        RESET   = "\x1b[0m"
        BOLD   = "\x1b[1m"
        BLACK = "\x1b[30m"
        RED     = "\x1b[31m"
        GREEN  = "\x1b[32m"
        YELLOW = "\x1b[33m"
        BLUE = "\x1b[34m"
        MAGENTA = "\x1b[35m"
        CYAN = "\x1b[36m"
        WHITE = "\x1b[37m"
        DEFAULT = "\x1b[39m"
        CLEAR = "\x1b[2J"
        HOME = "\x1b[H"

        BG_BLACK = "\x1b[40m"
        BG_RED = "\x1b[41m"
        BG_GREEN = "\x1b[42m"
        BG_YELLOW = "\x1b[43m"
        BG_BLUE = "\x1b[44m"
        BG_MAGENTA = "\x1b[45m"
        BG_CYAN = "\x1b[46m"
        BG_WHITE = "\x1b[47m"

)
const csvFile = "file.csv"

var reader = bufio.NewReader(os.Stdin)
var dummy string


type Inventory struct {
	id int
	Name string
	Condition string
	Status string
}

func main() {
	for {
		fmt.Print(CLEAR+HOME)
		fmt.Println(BG_BLUE+BLACK+" MENU INVENTARIS "+RESET)
		fmt.Println("1. Tampilkan barang")
		fmt.Println("2. Tambah barang")
		fmt.Println("3. Edit barang")
		fmt.Println("4. Hapus barang")
		fmt.Println("0. Keluar")
		fmt.Print("\n -> ")

		pilihan := inputInt("")

		switch pilihan {
		case 1:
			showInventory()
			fmt.Println("Ketuk enter untuk melanjutkan...")
			fmt.Scanln(&dummy)
		case 2:
			addInventory()
		case 3:
			editInventory()
		case 4:
			deleteInventory()
		case 0:
			return
		default:
			fmt.Println(RED+"MENU TIDAK VALID"+RESET)
			fmt.Println("ketuk enter untuk melanjutkan")
			fmt.Scanln(&dummy)
		}

	}
}

/////// helper
func input(label string) string {
	// Catatan: helper ini dipakai supaya input seperti nama bisa memakai spasi.
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputInt(label string) int {
	value, _ := strconv.Atoi(input(label))
	return value
}
///////
func showInventory() {
	inventories := readCsv()

	fmt.Print(CLEAR + HOME)
	fmt.Printf("%-5s", "ID")
	fmt.Printf("%-25s", "Nama")
	fmt.Printf("%-10s", "Kondisi")
	fmt.Printf("%-10s", "Status")
	fmt.Print("\n")

	for _, inventory := range inventories {
		fmt.Printf("%-5d", inventory.id)
		fmt.Printf("%-25s", inventory.Name)
		fmt.Printf("%-10s", inventory.Condition)
		fmt.Printf("%-10s", inventory.Status)
		fmt.Print("\n")
	}


}

func addInventory() {
	inventories := readCsv()


	name := input("Nama: ")
	kondisi := input("Kondisi: ")
	status := input("status: ")

	id := 1
	if len(inventories) > 0 {
		id = inventories[len(inventories)-1].id + 1
	}

	inventories = append(inventories, Inventory{id, name, kondisi, status})
	writeCsv(inventories)
	fmt.Println("Data ditambahkan")

}
func editInventory() {
	inventories := readCsv()


	showInventory()
	for {
		id := inputInt("ID yang diubah (0 keluar): ")

		if (id == 0) {return
		}else if id < 1 || id > len(inventories){
			fmt.Println("Id tidak valid...\nenter untuk melanjutkan...")
			fmt.Scanln(&dummy)
		}

		fmt.Printf("Apakah data: %s", inventories[id-1].Name)
		fmt.Printf("\nY/n: ")
		a := input("")
		if a == "n" || a == "N" {
			return
		}

		name := input("Nama baru: ")
		kondisi := input("Kondisi: ")
		status := input("Status: ")

		for i, inventory := range inventories {
			if inventory.id == id {
				inventories[i] = Inventory{id, name, kondisi, status}
				writeCsv(inventories)
				fmt.Println("Data diubah")
				return
			}
	}
	
		fmt.Println("Data tidak ditemukan")


	}

}
func deleteInventory() {
	inventories := readCsv()

	for {
		id := inputInt("ID yang diubah (0 keluar): ")

		if (id == 0) {return
		}else if id < 1 || id > len(inventories){
			fmt.Println("Id tidak valid...\nenter untuk melanjutkan...")
			fmt.Scanln(&dummy)
		}

		fmt.Printf("Apakah data: %s", inventories[id-1].Name)
		fmt.Printf("\nY/n: ")
		a := input("")
		if a == "n" || a == "N" {
			return
		}

		for i, inventory := range inventories {
			if inventory.id == id {
				inventories = append(inventories[:i], inventories[i+1:]...)
				writeCsv(inventories)
				fmt.Println("Data dihapus")
				return
			}
		}
	}
		fmt.Println("Data tidak ditemukan")

}
/**************************************/
func readCsv() []Inventory{

	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var inventories []Inventory
	for _, row := range rows[1:] {
		id, _ := strconv.Atoi(row[0])
		inventories = append(inventories, Inventory{id, row[1], row[2], row[3]})
	}
	return inventories


}
func writeCsv(inventories []Inventory) {
	file, err := os.Create(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"id", "Name", "Condition", "Status"})
	for _, inventory := range inventories {
		writer.Write([]string{
			strconv.Itoa(inventory.id),
			inventory.Name,
			inventory.Condition,
			inventory.Status,
		})
	}

}

