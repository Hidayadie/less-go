package main

import ("bufio" 
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
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
)

const ( // char
    	horizontal   = "─"
    	vertical     = "│"
    	topLeft      = "┌"
    	topRight     = "┐"
    	bottomLeft   = "└"
    	bottomRight  = "┘"
    	teeLeft      = "├"
    	teeRight     = "┤"
    	teeTop       = "┬"
    	teeBottom    = "┴"
    	cross        = "┼"
)

//--------------------------

type Item struct {
	Nama     string
	Harga    int
	Stok     int
}
type Keranjang struct {
	Nama	 string
	Harga	 int
	Jumlah	 int
}

var stok = []Item{
	{
		Nama:     "Beras        ",
		Harga:    15000,
		Stok:     20,
	},
	{
		Nama:     "Gula Pasir   ",
		Harga:    17500,
		Stok:     10,
	},
	{
		Nama:     "Minyak Goreng",
		Harga:    18000,
		Stok:     2,
	},
	{
		Nama:     "Telur Ayam   ",
		Harga:    28000,
		Stok:     35,
	},
	{
		Nama:     "Tepung Terigu",
		Harga:    12000,
		Stok:     10,
	},
	{
		Nama:     "Garam        ",
		Harga:    5000,
		Stok:     25,
	},
	{
		Nama:     "Kopi Bubuk   ",
		Harga:    14000,
		Stok:     20,
	},
	{
		Nama:     "Teh Celup    ",
		Harga:    8000,
		Stok:     15,
	},
}
 
var keranjang = []Keranjang{}
var dummy string

func hitungSubtotal(harga int, jumlah int) int {
	return harga * jumlah
}

func hitungDiskon(total int) int {
	return total * 10 / 100
}

func hitungTotal(keranjang []Keranjang) int {
	total := 0
	for _, item := range keranjang {
		total += item.Harga * item.Jumlah
	}

	return total
}

func bacaTeks(scanner *bufio.Scanner, label string) (string, error) {
	fmt.Print(label)
	if !scanner.Scan() {
		fmt.Println("ketuk enter untuk melanjutkan")
		fmt.Scanln(&dummy)
		return "", fmt.Errorf(RED + "input teks tidak valid" + RESET)

	}

	value := strings.TrimSpace(scanner.Text())
	if value == "" {
		fmt.Println("ketuk enter untuk melanjutkan")
		fmt.Scanln(&dummy)
		return "", fmt.Errorf(RED + "input tidak boleh kosong" + RESET)
	}

	return value, nil
}

func bacaAngkaPositif(scanner *bufio.Scanner, label string) (int, error) {
	fmt.Print(label)
	if !scanner.Scan() {
		return 0, fmt.Errorf(RED + "input angka tidak valid" + RESET)
	}

	value, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0, fmt.Errorf(RED + "input angka tidak valid" + RESET)
	}
	if value < 0 {
		return 0, fmt.Errorf(RED + "input harus lebih dari 0" + RESET)
	}

	return value, nil
}
func inputItem(scanner *bufio.Scanner, nomor int) (Keranjang, error) {
	fmt.Printf("\nItem ke-%d\n", nomor)

	var kode int
	var jumlah int

	for {
		tampilkanKatalog()

		var err error

		kode, err = bacaAngkaPositif(scanner, "Masukkan kode barang, 0 untuk kembali: ")
		if err != nil {
			fmt.Println("Pastikan kode produk benar...")
			fmt.Scanln(&dummy)
			continue
		}

		if kode > 8 || kode < 0 {
			fmt.Println("Pastikan kode produk benar...")
			fmt.Scanln(&dummy)
			continue
		}

		if kode == 0 {return Keranjang{},nil}

		jumlah, err = bacaAngkaPositif(scanner, "Masukkan jumlah barang: ")
		if err != nil {
			fmt.Println("Error. Ketuk enter untuk melanjutkan...")
			fmt.Scanln(&dummy)
			continue
		}

		if jumlah > stok[kode-1].Stok {
			fmt.Println("Jumlah yang Anda inginkan melebihi stok...")
			fmt.Scanln(&dummy)
			continue
		}

		break
	}

	kode--

	fmt.Println("\nBarang berhasil dimasukkan, ketuk enter untuk melanjutkan")
	fmt.Scanln(&dummy)

	stok[kode].Stok -= jumlah

	return Keranjang{
		Nama:   stok[kode].Nama,
		Jumlah: jumlah,
		Harga:  stok[kode].Harga,
	}, nil
}
func tampilkanKeranjang(keranjang []Keranjang) {
	fmt.Println("\x1b[2J")
	fmt.Println("\x1b[H")


	if len(keranjang) == 0 {
		fmt.Println(RED+"Keranjang masih kosong"+RESET)
		fmt.Println("ketuk enter untuk melanjutkan")
		fmt.Scanln(&dummy)
		return
		
	}

	fmt.Println("---=== Keranjang Anda ===---")
	fmt.Println("no\tNama\t\tHarga\tJumlah\tTotal")
	for i, item := range keranjang {
		fmt.Println(i+1," \t"+item.Nama,"\t",item.Harga,"\t",item.Jumlah,"\tRp",item.Jumlah * item.Harga)

	}

	fmt.Println("\nApakah anda ingin checkout?")
	fmt.Print("(Y/n): ")
	var input string
	if fmt.Scanln(&input); input == "N" || input == "n" {
		return	
	}

	total := hitungTotal(keranjang)
	if total >= 100000 {
		fmt.Println(GREEN+"Selamat!! Anda mendapatkan potongan 10%"+RESET)
		fmt.Println("Total harga: ", total, "diskon:", hitungDiskon(total))
		fmt.Println("hasil akhir:", total - hitungDiskon(total))
	} else {
		fmt.Println("Total harga: ", total)

	}
	fmt.Println("ketuk enter untuk melanjutkan ke pembayaran")
	fmt.Scanln(&dummy)

	fmt.Println(BLUE+"Beep Boop Beep Boop transaksi selesai..."+RESET)
	fmt.Scanln(&dummy)

	keranjang = slices.Delete(keranjang, 0, len(keranjang))

}


func tampilkanKatalog() {
	fmt.Println("\x1b[2J")
	fmt.Println("\x1b[H")
	fmt.Println("kode\tNama\t\t Harga\tStok")
	for i, Item := range stok {
		if Item.Stok == 0 {fmt.Print(RED)}
		fmt.Println("",i + 1,"", Item.Nama, "\t", Item.Harga, "\t", Item.Stok, RESET)
		
	}
}
func tampilkanPromo() {
	fmt.Println("\x1b[2J")
	fmt.Println("\x1b[H")

	fmt.Println(YELLOW+"SPESIAL HARI INI!!!"+RESET)
	fmt.Println("Minimal pembelian 100rb akan mendapatkan diskon 10%")
	fmt.Println("\nEnter untuk kembali ke menu utama...")
	fmt.Scanln(&dummy)
}
func tampilkanMenu() {
	fmt.Println("\x1b[2J")
	fmt.Println("\x1b[H")

	fmt.Println("\n=== CLI Kasir Sederhana ===")
	fmt.Println("1. Katalog barang")
	fmt.Println("2. Checkout keranjang")
	fmt.Println(MAGENTA+"3. Cek promo"+RESET)
	fmt.Println("4. Keluar")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		tampilkanMenu()

		pilihan, err := bacaAngkaPositif(scanner, "Pilih menu: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch pilihan {
		case 1:
			item, err := inputItem(scanner, len(keranjang)+1)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if item.Nama != "" {
				keranjang = append(keranjang, item)
			}

		case 2:
			tampilkanKeranjang(keranjang)
		case 3:
			tampilkanPromo()

		case 4:
			fmt.Println(GREEN + "Terima kasih" + RESET)
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
