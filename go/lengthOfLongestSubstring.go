/*
Medium
Longest Substring Without Repeating Characters
Topics
Companies
Hint
Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package main

import (
	"strings"
)

func lengthOfLongestSubstringO3n(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		count := 0
		for j := i; j < len(s); j++ {
			if !strings.ContainsAny(s[i:j], s[j:j+1]) {
				count++
				if count > longest {
					longest = count
				}
			} else {
				if count > longest {
					longest = count
				}
				break
			}
		}
	}
	return longest
}

// try to do it with O(n^2)
func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		clear(charIndexMap)
		for start = i; start < len(s); start++ {
			_, found := charIndexMap[s[start]]
			if found {
				length := len(charIndexMap)
				if length > maxLength {
					maxLength = length
				}
				break
			}
			charIndexMap[s[start]] = start
			if start == len(s)-1 {
				length := len(charIndexMap)
				if length > maxLength {
					maxLength = length
				}
			}
		}
	}

	return maxLength
}

func LengthOfLongestSubstring2(s string) int {
	// Create a map to store the last index of each character
	charIndexMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		// If the character is found in the map and is within the current substring
		if index, found := charIndexMap[s[i]]; found && index >= start {
			start = index + 1
		}

		// Update the last index of the character in the map
		charIndexMap[s[i]] = i

		// Calculate the maximum length of the substring
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
