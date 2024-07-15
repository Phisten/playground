/*
 * @lc app=leetcode id=2196 lang=typescript
 *
 * [2196] Create Binary Tree From Descriptions
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
function createBinaryTree(descriptions: number[][]): TreeNode | null {
  const ht = {};
  const parentIs = {};
  const n = descriptions.length;

  for (let i = 0; i < n; i++) {
    const parent = descriptions[i][0];
    const child = descriptions[i][1];

    if (!ht[parent]) ht[parent] = new TreeNode(parent);
    if (!ht[child]) ht[child] = new TreeNode(child);

    if (descriptions[i][2]) ht[parent].left = ht[child];
    else ht[parent].right = ht[child];

    parentIs[child] = parent;
  }

  for (let i = 0; i < n; i++) {
    if (!parentIs[descriptions[i][0]] && parentIs[descriptions[i][0]] !== 0)
      return ht[descriptions[i][0]];
  }

  return null;
}
// @lc code=end
