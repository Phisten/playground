/*
 * @lc app=leetcode id=2962 lang=cpp
 *
 * [2962] Count Subarrays Where Max Element Appears at Least K Times
 */

// @lc code=start
class Solution {
public:
    long long countSubarrays(vector<int>& nums, int k) {
        int n = nums.size();
        int maxNum = 0;
        for (int i = 0; i < n; i++) 
            maxNum = max(maxNum, nums[i]);

        long long ans = 0;
        int valid = 0;
        for (int r = 0, l = 0; r < n; r++) {
            int rVal = nums[r];
            if (rVal == maxNum) {
                valid++;
            }
            if (valid >= k) {
                ans += n - r;
            }

            while (l <= r && valid >= k) {
                int lVal = nums[l];
                if (lVal == maxNum) {
                    valid--;
                }
                if (valid >= k) {
                    ans += n - r;
                }
                l++;
            }
        }

        return ans;
    }
};
// @lc code=end
