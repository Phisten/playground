/*
 * @lc app=leetcode id=235 lang=typescript
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function lowestCommonAncestor(root: TreeNode | null, p: TreeNode | null, q: TreeNode | null): TreeNode | null {
  if (root === null) return null;
  if (root.val === q.val || root.val === p.val) return root;

const left = lowestCommonAncestor(root.left, p, q);
const right = lowestCommonAncestor(root.right, p, q);

  if (left === null && right === null) {
      return null;
  } else if (left && right === null) {
      return left;
  } else if (left === null && right) {
      return right;
  } else {
      return root;
  }
};
// @lc code=end

