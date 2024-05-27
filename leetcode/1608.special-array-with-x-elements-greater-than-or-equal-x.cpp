/*
 * @lc app=leetcode id=1608 lang=cpp
 *
 * [1608] Special Array With X Elements Greater Than or Equal X
 */

// @lc code=start
class Solution {
public:
    int specialArray(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int n = nums.size();

        for (int i = 0; i < n; i++)
            if (nums[i] >= n - i)
                if (i == 0 || nums[i-1] < n - i)
                    return n - i;

        return -1;
    }
};
// @lc code=end

