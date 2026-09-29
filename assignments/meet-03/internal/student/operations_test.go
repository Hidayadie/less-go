package student

import "testing"

func TestOperations(t *testing.T) {
	students := InitialData()

	if NextID(students) != 4 {
		t.Fatal("NextID harus mengembalikan ID setelah ID terbesar")
	}
	if !UpdateScore(students, 1, 95) || students[0].Score != 95 {
		t.Fatal("UpdateScore harus mengubah nilai student")
	}
	if got := len(RemoveInactive(students)); got != 2 {
		t.Fatalf("RemoveInactive len = %d, want 2", got)
	}
	if got := len(FilterByMajor(students, "backend")); got != 2 {
		t.Fatalf("FilterByMajor len = %d, want 2", got)
	}
	if AverageScore(nil) != 0 {
		t.Fatal("AverageScore data kosong harus 0")
	}
	if Grade(91) != "A" || Grade(82) != "B" || Grade(74) != "C" || Grade(60) != "D" {
		t.Fatal("Grade tidak sesuai batas nilai")
	}
}
