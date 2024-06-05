/*
 * @lc app=leetcode id=61 lang=cpp
 *
 * [61] Rotate List
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
    ListNode* rotateRight(ListNode* head, int k) {
        ListNode* cur;
        ListNode* last = NULL;
        ListNode* target = NULL;
        ListNode* newLast = NULL;

        int n = 0;
        cur = head;
        while (cur != NULL) {
            n++;
            cur = cur->next;
        }
        
        if (n > 0) k = k % n;
        if (k == 0) return head;

        int i = 0;
        cur = head;
        while (cur != NULL) {
            if (n-k == i) {
                target = cur;
                last->next = NULL;
            }

            i++;
            last = cur;
            cur = cur->next;
        }

        if (last != NULL)
            last->next = head;

        return target;
    }
};
// @lc code=end

