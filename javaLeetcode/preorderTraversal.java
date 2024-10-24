package javaLeetcode;

// /*
// Example 1:

// Input: root = [1,null,2,3]

// Output: [1,2,3]

// Explanation:



// Example 2:

// Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

// Output: [1,2,4,5,6,7,3,8,9]

// */


//  Definition for a binary tree node.

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
}

