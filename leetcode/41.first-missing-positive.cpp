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
        int hasMissing = false;
        int n = nums.size();

        for (int i = 0; i < n; i++)
        {
            while (nums[i] != i + 1 && nums[i] > 0 && nums[i] <= n)
            {
                int v = nums[i];
                if (v == nums[v - 1])
                    break;
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
// @lc code=end

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
