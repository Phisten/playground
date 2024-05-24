/*
 * @lc app=leetcode id=21 lang=kotlin
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Example:
 * var li = ListNode(5)
 * var v = li.`val`
 * Definition for singly-linked list.
 * class ListNode(var `val`: Int) {
 *     var next: ListNode? = null
 * }
 */
class Solution {
    fun mergeTwoLists(list1: ListNode?, list2: ListNode?): ListNode? {
        val head: ListNode = ListNode(0, null);
        var tail: ListNode? = head;
        var _list1 = list1;
        var _list2 = list2;

        while (true) {
            if (_list1 == null && _list2 == null) break;
            if (tail == null) break;

            if (_list1 == null) {
                tail.next = _list2
                _list2 = _list2?.next
            }
            else if (_list2 == null) {
                tail.next = _list1
                _list1 = _list1?.next
            }
            else if (_list1 != null && _list2 != null && _list1.`val` < _list2.`val`) {
                tail.next = _list1
                _list1 = _list1?.next
            } else {
                tail.next = _list2
                _list2 = _list2?.next
            }

            tail = tail?.next
        }

        return head.next;
    }
}
// @lc code=end
