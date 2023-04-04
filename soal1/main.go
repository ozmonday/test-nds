package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution(4))
}

// soal 1
func Solution(tinggi int) string {
	current := float32(tinggi)
	var result []string

	for {
		result = append(result, fmt.Sprintf("%.2f", current))
		current = current / 2
		if current <= 0.5 {
			result = append(result, fmt.Sprintf("%.2f", current))
			break
		}
	}
	return strings.Join(result, " ")
}

//soal 2
