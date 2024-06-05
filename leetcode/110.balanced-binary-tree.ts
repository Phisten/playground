/*
 * @lc app=leetcode id=110 lang=typescript
 *
 * [110] Balanced Binary Tree
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

function isBalanced(root: TreeNode | null): boolean {
  let ans = true;

  const dfs = (cur: TreeNode | null, deep: number) => {
      if (cur === null) {
          return deep - 1;
      }

      const l = dfs(cur.left, deep+1);
      const r = dfs(cur.right, deep+1);
      
      if (l - r > 1 || l - r < -1) ans = false;

      return Math.max(l, r);
  };

  dfs(root, 1);

  return ans;
};
// @lc code=end

