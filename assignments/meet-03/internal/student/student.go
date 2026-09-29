package student

type Student struct {
	ID     int
	Name   string
	Major  string
	Score  int
	Active bool
}

func InitialData() []Student {
	return []Student{
		{ID: 1, Name: "Budi", Major: "Backend", Score: 82, Active: true},
		{ID: 2, Name: "Siti", Major: "Frontend", Score: 91, Active: true},
		{ID: 3, Name: "Andi", Major: "Backend", Score: 74, Active: false},
	}
}

func NextID(students []Student) int {
	maxID := 0
	for _, student := range students {
		if student.ID > maxID {
			maxID = student.ID
		}
	}
	return maxID + 1
}
