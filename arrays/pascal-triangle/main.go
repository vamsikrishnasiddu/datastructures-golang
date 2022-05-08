package main

import "fmt"

func generate(numRows int) [][]int {

	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {

		arr := make([]int, i+1)
		triangle[i] = arr
	}

	for i := 0; i < len(triangle); i++ {

		triangle[i][0] = 1
		triangle[i][len(triangle[i])-1] = 1

		for j := 1; j < len(triangle[i])-1; j++ {

			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]

		}

	}

	return triangle

}

func generateBetter(numRows int) [][]int {

	triangle := make([][]int, numRows)

	for i := 0; i < len(triangle); i++ {

		triangle[i] = make([]int, i+1)

		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {

			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]

		}

	}

	return triangle

}

func getRow(rowIndex int) []int {

	n := rowIndex

	var arr = []int{}

	for j := 0; j < rowIndex+1; j++ {

		res := 1

		for i := 0; i < j; i++ {
			res = res * (n - i)
			res = res / (i + 1)
		}

		arr = append(arr, res)

	}

	return arr

}

func main() {

	numRows := 5

	arr := generate(numRows)

	fmt.Println(arr)

	brr := getRow(4)

	fmt.Println(brr)

}
