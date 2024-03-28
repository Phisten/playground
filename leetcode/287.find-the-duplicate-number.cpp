/*
 * @lc app=leetcode id=287 lang=cpp
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
#include <vector>
#include <unordered_map>
#include <iostream>
class Solution
{
public:
    int findDuplicate(std::vector<int> &nums)
    {
        std::unordered_map<int, int> ht;
        int n = nums.size();
        for (int i = 0; i < n; i++)
        {
            if (ht[nums[i]] == 1)
                return nums[i];
            ht[nums[i]] += 1;
        }

        return 0;
    }
};
// @lc code=end
