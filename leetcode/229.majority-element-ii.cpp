/*
 * @lc app=leetcode id=229 lang=cpp
 *
 * [229] Majority Element II
 */

// @lc code=start
class Solution {
public:
    vector<int> majorityElement(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, int> ht;
        vector<int> ans;

        for (int i = 0; i < n ; i++) {
            ht[nums[i]]++;
        }

        int limit = n / 3;
        for (int i = 0; i < n ; i++) {
            int cur = nums[i];
            if (ht[cur] > limit) {
                ans.push_back(cur);
                ht[cur] = 0;
            }
        }

        return ans;
    }
};
// @lc code=end

