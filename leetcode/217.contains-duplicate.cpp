/*
 * @lc app=leetcode id=217 lang=cpp
 *
 * [217] Contains Duplicate
 */

// @lc code=start
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        int n = nums.size();

        unordered_map<int, int> ht;
        for (int i = 0; i < n; i++) {
            ht[nums[i]] = ht[nums[i]] + 1;
            if (ht[nums[i]] >= 2)
                return true;
        }

        return false;
    }
};
// @lc code=end

