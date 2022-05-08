package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rowIndex := -1

	for i := 0; i < len(matrix); i++ {

		if target < matrix[i][len(matrix[i])-1] {
			rowIndex = i
			break
		} else if target == matrix[i][len(matrix[i])-1] {
			return true
		}

	}

	if rowIndex == -1 {

		return false
	}

	return binarySearch(matrix[rowIndex], target)

}

func searchMatrixOptimized(matrix [][]int, target int) bool {

	var mid int

	left := 0
	right := len(matrix)*len(matrix[0]) - 1

	for left <= right {

		mid = (left + right) / 2

		rowIdx := mid / len(matrix[0])
		colIdx := mid % len(matrix[0])

		if matrix[rowIdx][colIdx] == target {

			return true

		}

		if target < matrix[rowIdx][colIdx] {

			right = mid - 1

		} else {

			left = mid + 1

		}

	}

	return false

}

func binarySearch(arr []int, target int) bool {

	var mid int
	var left int = 0
	var right int = len(arr) - 1

	for left <= right {

		mid = (left + right) / 2

		if arr[mid] == target {
			return true
		}

		if target < arr[mid] {
			right = mid - 1

		} else {
			left = mid + 1
		}

	}

	return false

}

func main() {
	var matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}

	fmt.Println(matrix)

	fmt.Println(searchMatrix(matrix, 3))
}
