package main

import (
	"fmt"
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {

	var result = [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	pair := intervals[0]

	for i := 0; i < len(intervals); i++ {

		if intervals[i][0] <= pair[1] {
			pair[1] = int(math.Max(float64(intervals[i][1]), float64(pair[1])))
		} else {
			result = append(result, pair)

			pair = intervals[i]
		}
	}

	result = append(result, pair)

	return result

}

func main() {

	var intervals = [][]int{{1, 3}, {2, 4}, {2, 6}, {8, 9}, {8, 10}, {9, 11}, {15, 18}, {16, 17}}

	fmt.Println(merge(intervals))

}
