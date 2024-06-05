/*
 * @lc app=leetcode id=104 lang=typescript
 *
 * [104] Maximum Depth of Binary Tree
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

function maxDepth(root: TreeNode | null): number {
  const dfs = (cur: TreeNode | null, deep: number) => {
      if (cur === null) return deep - 1;

      return Math.max(
          dfs(cur.left, deep + 1),
          dfs(cur.right, deep + 1)
      );
  }

  return dfs(root, 1);
};
// @lc code=end

