/*
 * @lc app=leetcode id=41 lang=cpp
 *
 * [41] First Missing Positive
 */

// @lc code=start
class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        unordered_map<int, bool> ht;

        for (int i = 0; i < nums.size(); i++) {
            int v = nums[i];
            ht[v] = true;
        }

        for (int i = 1; i <= nums.size(); i++) {
            if (!ht[i]) {
                return i;
            }
        }

        return nums.size() + 1;
    }
};
// @lc code=end
