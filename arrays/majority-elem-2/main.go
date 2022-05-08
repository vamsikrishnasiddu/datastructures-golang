package main

import (
	"fmt"
	"sort"
)

func majorityElementBrute(nums []int) int {

	var pos int = -1
	var prev int = -9999

	for i := 0; i < len(nums); i++ {
		count := 0

		if nums[i] == prev {
			continue
		}

		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}

		if count > len(nums)/3 {
			pos = i
			break
		}
		prev = nums[i]
	}

	return nums[pos]

}

func majorityElement(nums []int) []int {

	sort.Slice(nums, func(x, y int) bool {
		return nums[x] > nums[y]
	})

	m := make(map[int]int)
	var result []int = []int{}

	var prev int = -9999

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++

		if m[nums[i]] > len(nums)/3 && nums[i] != prev {

			result = append(result, nums[i])

			prev = nums[i]

		}
	}

	return result
}

func majorityElementMoore(nums []int) []int {

	var result []int = []int{}

	number1 := -1
	number2 := -1

	count1 := 0
	count2 := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == number1 {
			count1++
		} else if nums[i] == number2 {
			count2++
		} else if count1 == 0 {
			number1 = nums[i]
			count1 = 1
		} else if count2 == 0 {
			number2 = nums[i]
			count2 = 1
		} else {
			count1--
			count2--
		}

	}

	count1 = 0
	count2 = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == number1 {
			count1++
		} else if nums[i] == number2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		result = append(result, number1)
	}

	if count2 > len(nums)/3 {
		result = append(result, number2)
	}

	return result
}

func main() {

	nums := []int{1, 1, 2, 2, 7, 7, 8, 8, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3}
	//nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(len(nums))

	fmt.Println(majorityElementMoore(nums))

}
