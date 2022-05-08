package main

import "fmt"

func mergeSort(arr []int) {

	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2

	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	for i := 0; i < mid; i++ {
		left[i] = arr[i]
	}

	for i := mid; i < len(arr); i++ {
		right[i-mid] = arr[i]
	}

	mergeSort(left)
	mergeSort(right)
	merge(arr, left, right)

}

func merge(arr []int, left []int, right []int) {

	i := 0
	j := 0
	k := 0
	nL := len(left)
	nR := len(right)

	for i < nL && j < nR {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < nL {
		arr[k] = left[i]
		i++
		k++
	}

	for j < nR {
		arr[k] = right[j]
		j++
		k++
	}
}

func main() {

	arr := []int{5, 4, 3, 6, 7}

	mergeSort(arr)

	fmt.Println(arr)

}
