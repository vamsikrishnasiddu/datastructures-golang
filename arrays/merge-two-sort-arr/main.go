package main

import (
	"fmt"
	"sort"
)

func mergeWithExtraSpace(nums1 []int, m int, nums2 []int, n int) {

	nums3 := make([]int, m+n)

	i := 0
	j := 0
	k := 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			nums3[k] = nums1[i]
			i++
		} else {
			nums3[k] = nums2[j]
			j++
		}

		k++
	}

	for i < m {
		nums3[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		nums3[k] = nums2[j]
		j++
		k++
	}

	fmt.Println(nums3)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0

	for i < m && j < n {

		if nums2[0] < nums1[i] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
			sort.Slice(nums2, func(x, y int) bool {
				return nums2[x] < nums2[y]
			})
		}
		i++
	}

	for j < n {
		nums1[i] = nums2[j]
		i++
		j++
	}

	fmt.Println(nums1)

}

func main() {
	var arr = []int{1, 2, 3, 0, 0, 0}
	var brr = []int{2, 5, 6}

	merge(arr, 3, brr, 3)

}
