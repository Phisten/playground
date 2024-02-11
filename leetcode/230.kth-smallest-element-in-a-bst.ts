/*
 * @lc app=leetcode id=230 lang=typescript
 *
 * [230] Kth Smallest Element in a BST
 */

import { TreeNode } from './LeetcodeDataStruct';
// @lc code=start
function kthSmallest(root: TreeNode | null, k: number): number {
  let ans: number | null = null;

  // inorder traversal recursion
  const dfs = (node: TreeNode | null) => {
    // If the node is null or the result is already found, return
    if (node == null || ans !== null) return;

    dfs(node.left);

    // Decrement the count of remaining nodes to visit
    k--;

    // If k becomes 0, we found the kth smallest element
    if (k === 0) {
      ans = node.val; // Store the value in the result variable
      return;
    }

    dfs(node.right);
  }

  dfs(root);

  return ans ?? 0;
};
// @lc code=end
