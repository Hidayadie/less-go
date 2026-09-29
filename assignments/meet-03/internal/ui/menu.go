package ui

import (
	"bufio"
	"fmt"
	"data-manipulation-packages/internal/student"
	"data-manipulation-packages/internal/input"
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

const ( // terminal layout
	max_width = 65
	max_height = 10
	control_width = 15
)

func PrintLanding(){
	
	fmt.Print(CLEAR+HOME)

	for i := 0 ; i < max_height; i++ {
        for j := 0; j < max_width; j++ {
			switch {
            // border
			case i == 0 && j == 0:
			fmt.Print(topLeft);

			case i == 0 && j == max_width -1:
			fmt.Print(topRight);

			case i == max_height-1 && j == 0:
			fmt.Print(bottomLeft);

			case i == max_height-1 && j == max_width-1:
			fmt.Print(bottomRight);

			// thee

			case i == 0 && j == max_width - control_width:
			fmt.Print(teeTop)

			case i == max_height-1 && j == max_width - control_width:
			fmt.Print(teeBottom)

			// line
			case i == 0 || i == max_height-1:
			fmt.Print(horizontal);

			case j == 0 || j == max_width -1 || j == max_width - control_width:
			fmt.Print(vertical);

			default:
			fmt.Print(" ");
            }
		}
		fmt.Print("\n");
    }
	fmt.Printf("\x1b[%d;%dH┐Student Management System┌", 0, max_width/2-20)
	fmt.Printf("\x1b[%d;%dH┐Menu┌", 0, max_width - control_width + 4)
	fmt.Printf("\x1b[%d;%dH SELAMAT DATANG", max_height/2, max_width/2-15)
	fmt.Printf("\x1b[%d;%dH ketuk enter untuk melanjutkan...", max_height/2 +1, max_width/2-25)

}

func PrintTable(){
	
	fmt.Print(CLEAR+HOME)

	for i := 0 ; i < max_height; i++ {
        for j := 0; j < max_width; j++ {
			switch {
            // border
			case i == 0 && j == 0:
			fmt.Print(topLeft);

			case i == 0 && j == max_width -1:
			fmt.Print(topRight);

			case i == max_height-1 && j == 0:
			fmt.Print(bottomLeft);

			case i == max_height-1 && j == max_width-1:
			fmt.Print(bottomRight);

			// thee

			case i == 0 && j == max_width - control_width:
			fmt.Print(teeTop)

			case i == max_height-1 && j == max_width - control_width:
			fmt.Print(teeBottom)

			// line
			case i == 0 || i == max_height-1:
			fmt.Print(horizontal);

			case j == 0 || j == max_width -1 || j == max_width - control_width:
			fmt.Print(vertical);

			default:
			fmt.Print(" ");
            }
		}
		fmt.Print("\n");
    }	
	fmt.Printf("\x1b[%d;%dH┐Student Management System┌", 0, max_width/2-20)
	fmt.Printf("\x1b[%d;%dH┐Menu┌", 0, max_width - control_width + 4)
	
}

func PrintMenu() {

	width_start := max_width - control_width + 2
	i := 2 //malas nulis height_start++ trus yahahaa
	
  	fmt.Printf("\x1b[%d;%dH1. Tambah", i, width_start)
	i++
	fmt.Printf("\x1b[%d;%dH2. Update", i, width_start)
	i++
	fmt.Printf("\x1b[%d;%dH3. Hapus", i, width_start)
	i++
	fmt.Printf("\x1b[%d;%dH4. Filter", i, width_start)
	i++
	fmt.Printf("\x1b[%d;%dH5. Cari", i, width_start)
	i++
	fmt.Printf("\x1b[%d;%dH6. Cek Nilai", i, width_start)
	i++
	fmt.Printf("\x1b[%d;%dH0. Keluar", i, width_start)
	i++
	fmt.Printf("\x1b[%d;%dH -> ", i, width_start)
	// posisi kursor input = 9, 52
}

func PrintStudents(students []student.Student) {
//func PrintStudents(){	
	width_start := 2
	i := 2 
	fmt.Print(BG_BLUE + BLACK)
	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Printf("%-4s", "ID")	
	fmt.Printf("%-17s", "Nama")	
	fmt.Printf("%-9s", "Major")	
	fmt.Printf("%-9s", "Nilai")	
	fmt.Printf("%-9s", "Aktif")
	fmt.Print(RESET)
	i++

	for _, current := range students {

		if !current.Active {fmt.Print(RED)}
		fmt.Printf("\x1b[%d;%dH", i, width_start)
		fmt.Printf("%-4d", current.ID)	
		fmt.Printf("%-17s",current.Name)	
		fmt.Printf("%-9s", current.Major)	
		fmt.Printf("%-9d", current.Score)
 		if current.Active {fmt.Printf("%-9s", "Aktif")
		} else {fmt.Printf("%-9s", "non-Aktif")}
		fmt.Print(RESET)
		i++
	}
	
}

func AddStudentMenu(scanner *bufio.Scanner, id int) student.Student{
	PrintTable()
	PrintMenu()

	width_start := 2
	i := 3

	//var student Student

	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Printf("Masukkan nama: ")
	nama := input.ReadString(scanner, "") 
	i++

	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Print("Masukkan major:")
	major := input.ReadString(scanner, "") 
	i++

	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Print("Masukkan nilai:")
	nilai := input.ReadInt(scanner, "") 

	return student.Student{
		ID: id,
		Name: nama,
		Major: major,
		Score: nilai,
		Active: true,
	}

}

func UpdateStudentMenuSelection(scanner *bufio.Scanner) int {
	for {
	PrintTable()
	PrintMenu()

	width_start := 2
	i := 3
	var dummy string
	//var student Student

	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Printf("Pilih mode edit...")
	i++

	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Print("1. Edit data siswa")
	i++

	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Print("2. hapus siswa non-aktif")
	i++
	
	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Print("0. Kembali")
	i++

	fmt.Printf("\x1b[%d;%dH", i, width_start)
	fmt.Print("-> ")
	i++

	input_angka:= input.ReadInt(scanner, "")
	if input_angka > 2 || input_angka < 0 {
		fmt.Printf("\x1b[%d;%dH", i, width_start)
		fmt.Print("Pastikan nilai benar...")
		fmt.Scanln(&dummy)
		continue
	} else {return input_angka}


	}


}

func UpdateStudentMenu(){
	  
}
