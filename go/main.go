package main

import (
	"fmt"
)

func main() {
	// str1 := "aust"
	// res := LengthOfLongestSubstring2(str1)
	// fmt.Println("Result: ", res)

	nums1 := []int{}
	nums2 := []int{1}
	median := FindMedianSortedArrays(nums1, nums2)

	// Displaying the result
	fmt.Println("median: ", median)
}