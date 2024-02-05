/*
 * @lc app=leetcode id=114 lang=typescript
 *
 * [114] Flatten Binary Tree to Linked List
 */

// @lc code=start

/**
 Do not return anything, modify root in-place instead.
 */
function flatten(root: TreeNode | null): void {
  const preOrder = (node: TreeNode | null): TreeNode | null => {
    if (!node) return null;

    const leftLast = preOrder(node.left)
    const rightLast = preOrder(node.right)

    const right = node.right;
    if (node.left)
      node.right = node.left;
    node.left = null;

    if (leftLast) {
      leftLast.right = right
    }

    return rightLast ?? leftLast ?? node;
  }

  preOrder(root);
};
// @lc code=end
