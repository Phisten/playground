/*
 * @lc app=leetcode id=53 lang=cpp
 *
 * [53] Maximum Subarray
 */

// @lc code=start
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int ans = nums[0];
        int n = nums.size();

        for (int i = 1; i < n; i++) {
            nums[i] = max(nums[i], nums[i] + nums[i-1]);
            ans = max(ans, nums[i]);
        }

        return ans;
    }
};
// @lc code=end
