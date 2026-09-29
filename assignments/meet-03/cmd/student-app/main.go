package main

import (
	"bufio"
	"fmt"
	"os"

	"data-manipulation-packages/internal/input"
	"data-manipulation-packages/internal/student"
	"data-manipulation-packages/internal/ui"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	students := student.InitialData()
	nextID := student.NextID(students)

	var dummy string

	ui.PrintLanding()
	fmt.Scanln(&dummy)


	for {
		ui.PrintTable()
		ui.PrintMenu()
		ui.PrintStudents(students)

		choice := input.ReadInt(scanner, "\x1b[9;56H")
		switch choice {
		case 1:
			student := ui.AddStudentMenu(scanner, nextID)
			if (student.Name != "") {
				students = append(students,student)
				nextID++
			} 
			
		case 2:
			update := ui.UpdateStudentMenuSelection(scanner)
			switch update {
			case 1:

			case 2:
				
			}


		case 3:
			id := input.ReadInt(scanner, "ID student: ")
			score := input.ReadInt(scanner, "Nilai baru: ")
			if student.UpdateScore(students, id, score) {
				fmt.Println("Nilai berhasil diupdate.")
			} else {
				fmt.Println("Student tidak ditemukan.")
			}
		case 4:
			before := len(students)
			students = student.RemoveInactive(students)
			fmt.Printf("%d student inactive dihapus.\n", before-len(students))
		case 5:
			//major := input.ReadString(scanner, "Jurusan: ")
			//ui.PrintStudents(student.FilterByMajor(students, major))
		case 6:
			//id := input.ReadInt(scanner, "ID student: ")
			//ui.PrintStudentByID(students, id)
		case 7:
			fmt.Printf("Rata-rata nilai: %.2f\n", student.AverageScore(students))
		case 0:
			fmt.Println("Selesai.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func addStudent(scanner *bufio.Scanner, students []student.Student, id int) []student.Student {
	name := input.ReadString(scanner, "Nama: ")
	major := input.ReadString(scanner, "Jurusan: ")
	score := input.ReadInt(scanner, "Nilai: ")

	fmt.Println("Student berhasil ditambahkan.")
	return append(students, student.Student{
		ID:     id,
		Name:   name,
		Major:  major,
		Score:  score,
		Active: true,
	})
}
