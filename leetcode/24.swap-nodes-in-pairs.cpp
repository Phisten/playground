/*
 * @lc app=leetcode id=24 lang=cpp
 *
 * [24] Swap Nodes in Pairs
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
    bool shouldSwap = true;
    ListNode* swapPairs(ListNode* head) {
        if (head == NULL) return NULL;
        if (head->next != NULL) {
            if (shouldSwap) {
                ListNode* srcHead = head;
                ListNode* newHead = head->next;
                srcHead->next = newHead->next;
                newHead->next = srcHead;
                head = newHead;
                
                shouldSwap = !shouldSwap;
                swapPairs(srcHead);
            } else {
                if (head->next->next != NULL) {
                    ListNode* next = head->next;
                    head->next = next->next;
                    shouldSwap = !shouldSwap;
                    swapPairs(next);
                }
                
            }
        }

        return head;
    }
};
// @lc code=end

