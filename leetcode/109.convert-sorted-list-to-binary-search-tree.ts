/*
 * @lc app=leetcode id=109 lang=typescript
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start

class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}

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

function sortedListToBST(head: ListNode | null): TreeNode | null {
  if (!head) return null;

  const arr: TreeNode[] = [];

  for (let item = head; item; item = item.next) {
    arr.push(new TreeNode(item.val, null, null))
    if (!item.next) break;
  }
  const len = arr.length;

  const dsf = (l: number, r: number): TreeNode | null => {
    if (l < 0 || r >= len || l > r) return null;

    const mid = Math.floor((l + r) / 2);
    arr[mid].left = dsf(l, mid - 1);
    arr[mid].right = dsf(mid + 1, r);

    return arr[mid];
  }

  return dsf(0, len - 1)
};
// @lc code=end
