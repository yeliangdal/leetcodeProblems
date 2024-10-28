package main

import (
	"math"
)

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(findSum(root.Left)-findSum(root.Right)) + findTilt(root.Left) + findTilt(root.Right)
}

func findSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Val + findSum(node.Left) + findSum(node.Right)
}

func abs(num int) int {
	if num < 0 {
		return 0 - num
	}
	return num
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index, max := findLargest(nums)
	node := TreeNode{
		Val:   max,
		Left:  nil,
		Right: nil,
	}
	if len(nums) == 1 {
		return &node
	}
	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])
	return &node
}

func findLargest(nums []int) (index int, max int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index, max
}

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	m := height + 1
	n := int(math.Exp2(float64(height+1))) - 1
	ret := make([][]string, m)
	for i := range ret {
		ret[i] = make([]string, n)
	}
	placeToMatrix(root, ret, m, (n-1)/2, 0)
	return ret
}

func placeToMatrix(node *TreeNode, matrix [][]string, height int, c int, row int) {
	if node == nil {
		return
	}
	if row >= height {
		return
	}
	if row == 0 {
		matrix[row][c] = string(node.Val)
	}
	matrix[row+1][c-int(math.Exp2(float64(c-row-1)))] = string(node.Left.Val)
	matrix[row+1][c+int(math.Exp2(float64(c-row-1)))] = string(node.Right.Val)
	placeToMatrix(node.Left, matrix, height, c+int(math.Exp2(float64(c-row-1))), row+1)
	placeToMatrix(node.Right, matrix, height, c-int(math.Exp2(float64(c-row-1))), row+1)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if getHeight(root.Left) > getHeight(root.Right) {
		return 1 + getHeight(root.Left)
	}
	return 1 + getHeight(root.Right)
}
