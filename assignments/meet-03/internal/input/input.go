package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadString(scanner *bufio.Scanner, label string) string {
	for {
		fmt.Print(label)
		if !scanner.Scan() {
			return ""
		}

		value := strings.TrimSpace(scanner.Text())
		if value != "" {
			return value
		}
		fmt.Println("Input tidak boleh kosong.")
	}
}

func ReadInt(scanner *bufio.Scanner, label string) int {
	for {
		value := ReadString(scanner, label)
		number, err := strconv.Atoi(value)
		if err == nil {
			return number
		}
		fmt.Println("Input harus angka.")
	}
}
