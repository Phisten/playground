/*
 * @lc app=leetcode id=109 lang=typescript
 */

// @lc code=start

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
