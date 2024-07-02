/*
 * @lc app=leetcode id=1791 lang=cpp
 *
 * [1791] Find Center of Star Graph
 */

// @lc code=start
class Solution {
public:
    int findCenter(vector<vector<int>>& edges) {
        unordered_map<int, int> ht;
        int n = edges.size();

        for (int i = 0; i < n; i++) {
            ht[edges[i][1]]++;
            ht[edges[i][0]]++;

            if (ht[edges[i][0]] > 1) return edges[i][0];
            if (ht[edges[i][1]] > 1) return edges[i][1];
        }

        return 1;
    }
};
// @lc code=end

