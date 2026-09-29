package student

import "strings"

func UpdateScore(students []Student, id int, score int) bool {
	for i := range students {
		if students[i].ID == id {
			students[i].Score = score
			return true
		}
	}
	return false
}

func RemoveInactive(students []Student) []Student {
	activeStudents := students[:0]
	for _, student := range students {
		if student.Active {
			activeStudents = append(activeStudents, student)
		}
	}
	return activeStudents
}

func FilterByMajor(students []Student, major string) []Student {
	var result []Student
	for _, student := range students {
		if strings.EqualFold(student.Major, major) {
			result = append(result, student)
		}
	}
	return result
}

func IndexByID(students []Student) map[int]Student {
	studentByID := make(map[int]Student, len(students))
	for _, student := range students {
		studentByID[student.ID] = student
	}
	return studentByID
}

func AverageScore(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}

	total := 0
	for _, student := range students {
		total += student.Score
	}
	return float64(total) / float64(len(students))
}

func Grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "D"
	}
}
