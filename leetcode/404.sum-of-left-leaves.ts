/*
 * @lc app=leetcode id=404 lang=typescript
 *
 * [404] Sum of Left Leaves
 */

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
// @lc code=start
function sumOfLeftLeaves(root: TreeNode | null): number {
  let ans = 0;

  const dfs = (node: TreeNode | null, isLeft: boolean): boolean => {
    if (node === null) return true;

    const left = dfs(node.left, true);
    const right = dfs(node.right, false);

    if (left && right && isLeft) ans += node.val;
  }
  dfs(root, false);

  return ans;
};
// @lc code=end
