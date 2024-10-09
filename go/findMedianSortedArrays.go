package main

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	count := 0
	isTotalEven := true

	if (len(nums1)+len(nums2))%2 == 0 {
		count = (len(nums1)+len(nums2))/2 - 1
	} else {
		count = (len(nums1) + len(nums2) - 1) / 2
		isTotalEven = false
	}

	pointer1, pointer2 := 0, 0
	for pointer1 < len(nums1) && pointer2 < len(nums2) && count > 0 {
		if nums1[pointer1] <= nums2[pointer2] {
			pointer1++
		} else {
			pointer2++
		}
		count--
	}

	if pointer1 == len(nums1) {
		pointer2 = pointer2 + count
	} else if pointer2 == len(nums2) {
		pointer1 = pointer1 + count
	}

	valuesForMedian := make([]int, 2)
	count = 0

	if pointer1 == len(nums1) {
		valuesForMedian[0] = nums2[pointer2]
		if pointer2+1 < len(nums2) {
			valuesForMedian[1] = nums2[pointer2+1]
		}
	} else if pointer2 == len(nums2) {
		valuesForMedian[0] = nums1[pointer1]
		if pointer1+1 < len(nums1) {
			valuesForMedian[1] = nums1[pointer1+1]
		}
	} else {
		for pointer1 < len(nums1) && pointer2 < len(nums2) && count < 2 {
			if nums1[pointer1] <= nums2[pointer2] {
				valuesForMedian[count] = nums1[pointer1]
				pointer1++
			} else {
				valuesForMedian[count] = nums2[pointer2]
				pointer2++
			}
			count++
		}
		if count < 2 {
			if pointer1 == len(nums1) {
				valuesForMedian[1] = nums2[pointer2]
			} else if pointer2 == len(nums2) {
				valuesForMedian[1] = nums1[pointer1]
			}
		}
	}
	if isTotalEven {
		return (float64(valuesForMedian[0]) + float64(valuesForMedian[1])) / 2
	} else {
		return float64(valuesForMedian[0])
	}

}
