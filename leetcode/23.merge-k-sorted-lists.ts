/*
 * @lc app=leetcode id=23 lang=typescript
 *
 * [23] Merge k Sorted Lists
 */

class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}

// @lc code=start
function mergeKLists(lists: Array<ListNode | null>): ListNode | null {
  const arr: ListNode[] = [];
  for (let i = 0; i < lists.length; i++) {
    let item = lists[i];
    while (item !== null) {
      arr.push(item);
      item = item.next;
    }
  }

  arr.sort((a, b) => a.val - b.val);
  for (let i = 0; i < arr.length - 1; i++)
    arr[i].next = arr[i + 1];

  return arr.length ? arr[0] : null;
};
// @lc code=end
