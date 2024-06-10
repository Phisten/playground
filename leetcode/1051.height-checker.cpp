/*
 * @lc app=leetcode id=1051 lang=cpp
 *
 * [1051] Height Checker
 */

// @lc code=start
class Solution {
public:
    int heightChecker(vector<int>& heights) {
        vector<int> src(heights);

        sort(heights.begin(), heights.end());

        int n = heights.size();
        int count = 0;
        for (int i = 0; i < n; i++) {
            if (src[i] != heights[i]) count++;
        }
        return count;
    }
};
// @lc code=end

