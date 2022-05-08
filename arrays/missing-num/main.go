package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != num {
			break
		}

		num++
	}

	return num

}

func main() {

	var arr = []int{3, 0, 1}

	res := missingNumber(arr)

	fmt.Println(res)

}
