/*
 * @lc app=leetcode id=153 lang=cpp
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
class Solution {
public:
    int findMin(vector<int>& nums) {
        int n = nums.size();
        int l = 0 , r = n - 1;

        while (l < r) {
            int m = (l + r) / 2;
            int left = nums[l];
            int right = nums[r];
            int mid = nums[m];

            if (left < right) {
                r = l;
            } else {
                if (mid > right) {
                    l = m + 1;
                } else {
                    r = m;
                }
            }
        }

        return nums[l];
    }
};
// @lc code=end

