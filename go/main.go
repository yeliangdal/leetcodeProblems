package main

import (
	"fmt"
)

func main() {
	// str1 := "aust"
	// res := LengthOfLongestSubstring2(str1)
	// fmt.Println("Result: ", res)

	// nums1 := []int{}
	// nums2 := []int{1}
	// median := FindMedianSortedArrays(nums1, nums2)
	// fmt.Println("median: ", median)

	// s := ""
	// subS := LongestPalindrome1(s)
	input := []int{3, 2, 1}
	nodes := constructMaximumBinaryTree(input)
	fmt.Println("results: ", nodes)
}
