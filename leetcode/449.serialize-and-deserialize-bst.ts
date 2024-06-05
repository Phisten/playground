/*
 * @lc app=leetcode id=449 lang=typescript
 *
 * [449] Serialize and Deserialize BST
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
/*
 * Encodes a tree to a single string.
 */
function serialize(root: TreeNode | null): string {
  let ss = [];
  const dfs = (cur: TreeNode | null) => {
      if (cur === null) {
          ss.push("_");
          return;
      }
      ss.push(cur.val);
      dfs(cur.left);
      dfs(cur.right);
  };
  dfs(root);
  const ans = ss.join(',');

  return ans;
};

/*
* Decodes your encoded data to tree.
*/
function deserialize(data: string): TreeNode | null {
  const arr = data.split(',');
  const n = arr.length;
  if (n === 0) return null;

  let nextDirect = -1;

  let i = 0;
  const dfs = () => {
      if (i >= n) return null;

      const curValue = arr[i] === "_" ? null : Number(arr[i]);
      i++;
      if (curValue === null) return null;
      return new TreeNode(curValue, dfs(), dfs());
  }

  return dfs();
};
/**
* Your functions will be called as such:
* deserialize(serialize(root));
*/
// @lc code=end

