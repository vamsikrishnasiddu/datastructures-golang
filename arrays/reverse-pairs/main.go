package main

import "fmt"

func main() {

	nums := []int{2, 4, 3, 5, 1}

	fmt.Println(reversePairs(nums))

}

func reversePairs(nums []int) int {

	count := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if i >= 0 && i < j && j < len(nums) && nums[i] > 2*nums[j] {
				count++
			}
		}
	}

	return count

}
