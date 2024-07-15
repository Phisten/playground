/*
 * @lc app=leetcode id=2196 lang=cpp
 *
 * [2196] Create Binary Tree From Descriptions
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
class Solution {
public:
    unordered_map<int, TreeNode*> nodes;
    TreeNode* createBinaryTree(vector<vector<int>>& descriptions) {
        unordered_map<int, int> parentIs;
        
        // build nodes
        int n = descriptions.size();
        for (int i = 0; i < n; i++) {
            int parent = descriptions[i][0];
            int child = descriptions[i][1];
            parentIs[child] = parent;
            
            if (nodes.count(parent) == 0)
                nodes[parent] = new TreeNode(parent);
            if (nodes.count(child) == 0) 
                nodes[child] = new TreeNode(child);
            
            if (descriptions[i][2])
                nodes[parent]->left = nodes[child];
            else
                nodes[parent]->right = nodes[child];
        }

        // build tree
        TreeNode* root = nullptr;
        for (int i = 0; i < n; i++) {
            int cur = descriptions[i][0];

            if (parentIs.count(cur) == 0) {
                root = nodes[cur];
            }
        }

        return root;
    }
};
// @lc code=end
