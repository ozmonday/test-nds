package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution(13, 18, 10))
}

func Solution(nilai1 int, nilai2 int, deret int) string {
	var res []string
	var current int
	for i := 0; i < deret; i++ {
		if i == 0 {
			current = nilai1
			res = append(res, fmt.Sprintf("%d", nilai1))
			continue
		}

		if i == 1 {
			current = nilai2
			res = append(res, fmt.Sprintf("%d", nilai2))
			continue
		}

		val := current + nilai2 - nilai1

		res = append(res, fmt.Sprintf("%d", val))
		current = val
	}

	return strings.Join(res, " ")
}
