/*
 * @lc app=leetcode id=442 lang=cpp
 *
 * [442] Find All Duplicates in an Array
 */

// @lc code=start
class Solution
{
public:
    vector<int> findDuplicates(vector<int> &nums)
    {
        vector<int> ans;
        unordered_map<int, int> ht;
        for (int i = 0; i < nums.size(); i++)
        {
            int v = nums[i];
            if (ht[v] == 1)
                ans.push_back(v);
            ht[v]++;
        }

        return ans;
    }
};
// @lc code=end
