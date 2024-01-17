/*
 * @lc app=leetcode id=116 lang=cpp
 *
 * [116] Populating Next Right Pointers in Each Node
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

#include "cpp/Node.cpp"
#include <vector>

using namespace std;

class Solution
{
    vector<Node *> vt;

    void preorder(Node *curNode, int lv)
    {
        if (curNode == NULL)
            return;
        if (vt.size() == lv)
            vt.push_back(curNode);
        else if (vt[lv] != NULL)
            vt[lv]->next = curNode;
        vt[lv] = curNode;

        preorder(curNode->left, lv + 1);
        preorder(curNode->right, lv + 1);
    }

public:
    Node *connect(Node *root)
    {
        preorder(root, 0);
        return root;
    }
};
// @lc code=end
