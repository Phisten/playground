/*
 * @lc app=leetcode id=328 lang=cpp
 *
 * [328] Odd Even Linked List
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
    ListNode oddNodeHead = ListNode();
    ListNode* oddNodeEnd = &oddNodeHead;
    ListNode evenNodeHead = ListNode();
    ListNode* evenNodeEnd = &evenNodeHead;
    int curNum = 0;

    ListNode* oddEvenList(ListNode* head) {
        if (head == NULL) {
            oddNodeEnd->next = evenNodeHead.next;
            evenNodeEnd->next = NULL;
            return NULL;
        }

        curNum++;
        if (curNum % 2 == 1) {
            oddNodeEnd->next = head;
            oddNodeEnd = head;
        } else {
            evenNodeEnd->next = head;
            evenNodeEnd = head;
        }

        oddEvenList(head->next);

        return oddNodeHead.next;
    }
};
// @lc code=end

