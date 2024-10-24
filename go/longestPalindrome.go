package main

/*
Given a string s, return the longest
palindromic

substring

	in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/
// func LongestPalindrome(s string) string {
// 	maxLength := 0
// 	left, right := 0
// 	for index := 0; index < len(s); index++ {
// 		count := 0
// 		for index-count >= 0 && index+count < len(s) {
// 			if s[index-count] != s[index+count] {
// 				break
// 			} else {
// 				count++
// 			}
// 		}
// 		if count > maxLength {
// 			maxLength = count
// 			palindrome = s[index-count : index+count]
// 		}
// 	}
// 	return palindrome
// }

// o(n^3)
func LongestPalindrome1(s string) string {
	maxLength := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if isPalindrome(s[i:j]) && j-i > maxLength {
				maxLength = j - i
				left = i
				right = j
			}
		}
	}
	return s[left:right]
}

// o(n^2)
// func LongestPalindrome(s string) string {
// 	maxLength := 0
// 	left, right := 0, 0
// 	for i := 0; i < len(s); i++ {
// 		if
// 	}
// 	return s[left:right]
// }

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	return s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1])
}

func findPalindromeLenth(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	// return s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1])
	count := 0
	for index := len(s)/2 - 1; index >= 0; index-- {
		if s[index] != s[len(s)-index-1] {
			break
		} else {
			count++
		}
	}
	return count
}
