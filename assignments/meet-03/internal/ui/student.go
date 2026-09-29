package ui

import (
	//"fmt"
	//"sort"

	//"data-manipulation-packages/internal/student"
)
/*
func PrintStudents(students []student.Student) {
	if len(students) == 0 {
		fmt.Println("Data student kosong.")
		return
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Score > students[j].Score
	})

	fmt.Println("Daftar student:")
	for _, current := range students {
		status := "inactive"
		if current.Active {
			status = "active"
		}
		fmt.Printf("- ID %d | %s | %s | nilai %d | grade %s | %s\n",
			current.ID, current.Name, current.Major, current.Score, student.Grade(current.Score), status)
	}
}
*/
/*
func PrintStudentByID(students []student.Student, id int) {
	studentByID := student.IndexByID(students)
	if current, ok := studentByID[id]; ok {
		PrintStudents([]student.Student{current})
		return
	}
	fmt.Println("Student tidak ditemukan.")
}
*/
