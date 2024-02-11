/*
 * @lc app=leetcode id=230 lang=typescript
 *
 * [230] Kth Smallest Element in a BST
 */

import { TreeNode } from './LeetcodeDataStruct';

// @lc code=start
function kthSmallest(root: TreeNode | null, k: number): number {

  let ans: number | null = null;
  const dsf = (node: TreeNode | null) => {
    if (node == null || ans !== null) return;

    dsf(node.left);

    k--;
    if (k === 0) {
      ans = node.val;
      return;
    }

    dsf(node.right);
  }
  dsf(root);

  return ans ?? 0;
};
// @lc code=end
