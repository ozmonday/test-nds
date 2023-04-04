package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution(10, 350))
}

func Solution(n int, double int) string {

	ds := fmt.Sprintf("%d00", double)
	var result []string

	for i := 0; i < n-len(ds); i++ {
		result = append(result, "0")
	}

	result = append(result, ds)

	return strings.Join(result, "")
}
