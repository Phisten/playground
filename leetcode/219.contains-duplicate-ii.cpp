/*
 * @lc app=leetcode id=219 lang=cpp
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        int n = nums.size();

        unordered_map<int, int> ht;
        for (int i = 0; i < n; i++) {
            if (ht.count(nums[i]) > 0) {
                if (k >= i - ht[ nums[i] ]) {
                    return true;
                }
            }

            ht[ nums[i] ] = i;
        }

        return false;
    }
};
// @lc code=end

