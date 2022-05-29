package main

import "fmt"

func countNOOfSubArrWithGivenXorK(nums []int, b int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		current_xor := 0
		for j := i; j < len(nums); j++ {

			current_xor = current_xor ^ nums[j]

			if current_xor == b {
				count++
			}

		}
	}

	return count

}

func countNoOfSubArraysWithGivenXorKOn(nums []int, k int) int {

	count := 0
	xor := 0
	freq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]

		if val, ok := freq[xor^k]; ok {
			count = count + val
		}

		if xor == k {
			count++
		}

		if _, ok := freq[xor]; ok {
			freq[xor]++
		} else {
			freq[xor] = 1
		}

	}

	return count
}

func main() {

	var nums []int = []int{4, 2, 2, 6, 4}
	var b int = 6

	result := countNoOfSubArraysWithGivenXorKOn(nums, b)
	fmt.Println(result)

}
