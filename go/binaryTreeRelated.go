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