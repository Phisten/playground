/*
 * @lc app=leetcode id=2441 lang=cpp
 *
 * [2441] Largest Positive Integer That Exists With Its Negative
 */

// @lc code=start
class Solution {
public:
    int findMaxK(vector<int>& nums) {
        int n = nums.size();

        unordered_map<int, bool> ht;
        int ans = -1;
        for (int i = 0; i < n; i++) {
            ht[nums[i]] = true;
            if (ht[nums[i] * -1])
                ans = max(ans, abs(nums[i]));
        }

        return ans;
    }
};// @lc code=end

