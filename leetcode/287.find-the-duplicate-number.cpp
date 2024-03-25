/*
 * @lc app=leetcode id=287 lang=cpp
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
class Solution
{
public:
    int findDuplicate(vector<int> &nums)
    {
        unordered_map<int, int> ht;
        for (int i = 0; i < nums.size(); i++)
        {
            if (ht[nums[i]] == 1)
                return nums[i];
            ht[nums[i]] += 1;
        }
        return 0;
    }
};
// @lc code=end
