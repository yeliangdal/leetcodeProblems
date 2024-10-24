package javaLeetcode;

import java.util.*;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        if (root == null) {
            return new ArrayList<>();
        }
        List<Integer> ret = new ArrayList<>();
        ret.add(root.val);
        ret.addAll(preorderTraversal(root.left));
        ret.addAll(preorderTraversal(root.right));
        return ret;
    }

    public List<Integer> postorderTraversal(TreeNode root) {
        if (root == null) {
            return new ArrayList<>();
        }
        List<Integer> ret = new ArrayList<>();
        ret.addAll(postorderTraversal(root.left));
        ret.addAll(postorderTraversal(root.right));
        ret.add(root.val);
        return ret;
    }

    public boolean isBalanced(TreeNode root) {
        if (root == null) 
            return true;
        if ( height(root) == -1) 
            return false;
        return true;
    }

    int height(TreeNode node) {
       if (node == null) {
        return 0;
       } 
       int leftHeight = height(node.left); 
       int rightHeight = height(node.right); 

       if (leftHeight == -1 || rightHeight == -1)
        return -1;
       if (Math.abs(leftHeight- rightHeight) > 1) 
        return -1;
       return Math.max(leftHeight, rightHeight) + 1;
    }

    public TreeNode invertTree(TreeNode root) {
        TreeNode tmp = root.left;
        root.left = root.right;
        root.right = tmp;
        invertTree(root.left);
        invertTree(root.right);
        return root;
    }

    public List<String> binaryTreePaths(TreeNode root) {
        List<String> paths = new ArrayList<>();
        if (root == null) {
            return paths;
        }
        if (root.left == null && root.right == null) {
            paths.add(String.valueOf(root.val));
            return paths;
        }
        List<String> leftPath = binaryTreePaths(root.left);
        List<String> rightPath = binaryTreePaths(root.right);
        for (String path:leftPath) {
            paths.add(String.valueOf(root.val) + "->" + path);
        }
        for (String path:rightPath) {
            paths.add(String.valueOf(root.val) + "->" + path);
        }
        return paths;
    }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}
