package problem1051

//# [1051. Height Checker](https://leetcode.com/problems/height-checker/)
//
//Students are asked to stand in non-decreasing order of heights for an annual photo.
//
//Return the minimum number of students not standing in the right positions.  (This is the number of students that must move in order for all students to be standing in non-decreasing order of height.)
//
//Example 1:
//
//```text
//Input: [1,1,4,2,1,3]
//Output: 3
//Explanation:
//Students with heights 4, 3 and the last 1 are not standing in the right positions.
//```
//
//Note:
//
//- 1 <= heights.length <= 100
//- 1 <= heights[i] <= 100

// 这题其实就是看排序后有多少和排序前不一样

import "sort"

// 解法1
func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))

	sorted = append(sorted)

	count := 0
	copy(sorted, heights)
	sort.Ints(sorted)

	for i := range heights {
		if heights[i] != sorted[i] {
			count++
		}
	}

	return count
}

// 解法2

//func heightChecker(heights []int) int {
//	count := 0
//	hc := make([]int, 101)
//	for _, v := range heights {
//		hc[v]++
//	}
//	j := 0
//	for i := 1;i < len(hc); i++ {
//		for hc[i] > 0 {
//			if heights[j] != i {
//				count++
//			}
//			j++
//			hc[i]--
//		}
//	}
//	return count
//}
