/*
 * @lc app=leetcode id=998 lang=cpp
 *
 * [998] Maximum Binary Tree II
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
class Solution {
public:
    string smallestFromLeaf(TreeNode* root) {
        if (root == nullptr) return "";
        string minStr = "";

        function<bool(TreeNode*, string)> dfs = [&](TreeNode* node, string lastStr) -> bool {
            if (node == nullptr) return true;

            string s = "a";
            s[0] += node->val;
            s = s + lastStr;

            bool l = dfs(node->left, s);
            bool r = dfs(node->right, s);
            if (l and r) {
                if (minStr.length() == 0)
                    minStr = s;
                else
                    minStr = min(minStr, s);
            }
            return false;
        };
        dfs(root, "");

        return minStr;
    }
};
// @lc code=end
