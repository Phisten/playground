/*
 * @lc app=leetcode id=404 lang=cpp
 *
 * [404] Sum of Left Leaves
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution
{
public:
    int dsf(TreeNode *node, bool isLeft)
    {
        if (node == nullptr)
            return 0;

        if (node->left == nullptr && node->right == nullptr && isLeft)
        {
            return node->val;
        }
        int l = dsf(node->left, true);
        int r = dsf(node->right, false);

        return l + r;
    }
    int sumOfLeftLeaves(TreeNode *root)
    {
        return dsf(root, false);
    }
};
// @lc code=end
