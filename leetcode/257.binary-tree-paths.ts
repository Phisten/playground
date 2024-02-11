/*
 * @lc app=leetcode id=257 lang=typescript
 *
 * [257] Binary Tree Paths
 */

// @lc code=start
class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
  }
}

function binaryTreePaths(root: TreeNode | null): string[] {
  if (root == null) return [];

  const ans: string[] = [];


  const dsf = (node: TreeNode | null, str) => {
    if (node === null) return;

    str = str + '->' + node.val.toString()

    if (!node.left && !node.right) {
      ans.push(str);
      return;
    }

    dsf(node.left, str);
    dsf(node.right, str);
  };

  if (!root.left && !root.right) {
    ans.push(root.val.toString());
  } else {
    dsf(root.left, root.val.toString());
    dsf(root.right, root.val.toString());
  }

  return ans;
};
// @lc code=end
