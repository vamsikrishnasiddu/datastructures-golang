package main

import "fmt"

func sortColors(arr []int) {
	brr := []int{}

	for i := 0; i < len(arr); i++ {

		if arr[i] == 0 {
			brr = append(brr, arr[i])
		}
	}

	for i := 0; i < len(arr); i++ {

		if arr[i] == 1 {
			brr = append(brr, arr[i])
		}
	}

	for i := 0; i < len(arr); i++ {

		if arr[i] == 2 {
			brr = append(brr, arr[i])
		}
	}
	copy(arr, brr)
}

func sortColorsBubbleSort(arr []int) {
	// bubble sort

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

// Dutch NationalFlag Algorithm
func sortColorsOptimal(arr []int) {

	low := 0
	mid := 0
	high := len(arr) - 1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else if arr[mid] == 2 {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

func main() {
	var arr = []int{2, 0, 1}

	sortColorsOptimal(arr)

	fmt.Println(arr)
}
