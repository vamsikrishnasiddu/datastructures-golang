package main

import "fmt"

func countInversions(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[i] {
				count++
			}

		}
	}

	return count
}

func main() {
	var arr = []int{5, 3, 2, 1, 4}

	fmt.Println(countInversions(arr))
}
