/*
 * @lc app=leetcode id=19 lang=cpp
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        int len = 0;
        function<int(ListNode*)> dsf = [&](ListNode* node) -> int {
            if (node == NULL) {
                return 1;
            } else {
                len++;
                int curNum = dsf(node->next);
                
                if (curNum == n) return -1;
                else if (curNum == -1) {
                    node->next = node->next->next;
                    return -2;
                }
                else if (curNum == -2) return -2;

                return curNum + 1;
            }

        };
        dsf(head);

        if (len-n > 0) return head;
        return head->next;
    }
};
// @lc code=end

