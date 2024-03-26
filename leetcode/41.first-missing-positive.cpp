/*
 * @lc app=leetcode id=41 lang=cpp
 *
 * [41] First Missing Positive
 */

// @lc code=start
class Solution
{
public:
    int firstMissingPositive(vector<int> &nums)
    {
        int n = nums.size();
        nums.push_back(0);
        int cap = nums.size();

        // fixed range in 0 to 10^5
        for (int i = 0; i < n; i++)
            if (nums[i] < 0 || nums[i] > n)
                nums[i] = 0;

        // counting
        for (int i = 0; i < n; i++)
            if (nums[i] != 0)
                nums[(nums[i] % cap)] += cap;

        for (int i = 1; i <= n; i++)
            if (nums[i] / cap == 0)
                return i;

        return cap;
    }
};
// @lc code=end

class Solution_nPlusTime_1Space
{
public:
    int firstMissingPositive(vector<int> &nums)
    {
        int n = nums.size();

        for (int i = 0; i < n; i++)
        {
            while (
                nums[i] > 0 && nums[i] <= n &&
                nums[i] != nums[nums[i] - 1])
            {
                int v = nums[i];
                nums[i] = nums[v - 1];
                nums[v - 1] = v;
            }
        }

        for (int i = 0; i < n; i++)
            if (nums[i] != i + 1)
                return i + 1;

        return n + 1;
    }
};

class Solution_ht_nSpace
{
public:
    int firstMissingPositive(vector<int> &nums)
    {
        unordered_map<int, bool> ht;

        for (int i = 0; i < nums.size(); i++)
        {
            int v = nums[i];
            ht[v] = true;
        }

        for (int i = 1; i <= nums.size(); i++)
        {
            if (!ht[i])
            {
                return i;
            }
        }

        return nums.size() + 1;
    }
};
