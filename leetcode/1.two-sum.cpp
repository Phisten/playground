/*
 * @lc app=leetcode id=1 lang=cpp
 *
 * [1] Two Sum
 */

// @lc code=start
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        
        int n = nums.size();
        unordered_map<int, int> ht;

        for (int i = 0; i < n; i++) {
            ht[nums[i]] = i;
        }
        for (int i = 0; i < n; i++) {
            int tryFind = target - nums[i];

            if (ht.count(tryFind) > 0 and i != ht[tryFind]) {
                vector<int> ans({ i, ht[tryFind] });
                return ans;
            }
        }

        vector<int> ans;
        return ans;
    }
};
// @lc code=end

