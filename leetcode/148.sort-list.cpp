/*
 * @lc app=leetcode id=148 lang=cpp
 *
 * [148] Sort List
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
    ListNode* mergeSort(ListNode* cur, int n) {
        if (cur == NULL) return NULL;
        if (n == 1) {
            cur->next = NULL;
            return cur;
        }

        ListNode* r = cur;
        for (int i = 0; i < n / 2; i++) r = r->next;
        r = mergeSort(r, n - n / 2);
        
        ListNode* l = mergeSort(cur, n / 2);

        ListNode head;
        ListNode* curEnd = &head;
        while (l != NULL or r != NULL) {
            if (l == NULL || (r != NULL and r->val < l->val)) {
                curEnd->next = r;
                curEnd = r;
                r = r->next;
            }
            else {
                curEnd->next = l;
                curEnd = l;
                l = l->next;
            }
        }
        curEnd->next = NULL;

        return head.next;
    }
    ListNode* sortList(ListNode* head) {
        int n = 0;
        ListNode *cur = head;
        while (cur) {
            n++;
            cur = cur->next;
        }

        head = mergeSort(head, n);

        return head;
    }
};
// @lc code=end

