package main

import "fmt"

func majorityElement(nums []int) int {

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

		if count > len(nums)/2 {
			pos = i
			break
		}
		prev = nums[i]
	}

	return nums[pos]

}

func majorityElementMap(nums []int) int {

	m := make(map[int]int)
	var result int

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++

		if m[nums[i]] > len(nums)/2 {

			result = nums[i]
			break

		}
	}

	return result
}

func majorityElementmoorealogo(nums []int) int {

	count := 0
	el := 0
	for i := 0; i < len(nums); i++ {

		if count == 0 {
			el = nums[i]
		}

		if nums[i] == el {
			count++
		} else {
			count--
		}

		fmt.Println("nums[i],el,count", nums[i], el, count)
	}

	return el
}

func main() {

	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElementmoorealogo(nums))

}
