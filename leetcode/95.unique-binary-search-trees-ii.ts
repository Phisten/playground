/*
 * @lc app=leetcode id=95 lang=typescript
 *
 * [95] Unique Binary Search Trees II
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

function generateTrees(n: number): Array<TreeNode | null> {
  const dsf = (lBound, rBound): (TreeNode | null)[] => {
    if (lBound > rBound) return [null];

    const dsfRes: (TreeNode | null)[] = [];

    for (let i = lBound; i <= rBound; i++) {
      const leftRes = dsf(lBound, i - 1)
      const rightRes = dsf(i + 1, rBound)

      for (let leftItem of leftRes) {
        for (let rightItem of rightRes) {
          dsfRes.push(new TreeNode(i, leftItem, rightItem));
        }
      }
    }

    return dsfRes;
  }

  return dsf(1, n);
};
// @lc code=end
