/*
 * @lc app=leetcode id=322 lang=cpp
 *
 * [322] Coin Change
 */

// @lc code=start
#include <unordered_map>
#include <vector>
#include <functional>
using namespace std;

class Solution
{
public:
    int coinChange(vector<int> &coins, int amount)
    {
        unordered_map<int, int> dp;

        function<long long(int)> dfs = [&](int last_amount) -> long long
        {
            long long res = numeric_limits<int>::max();
            if (last_amount == 0)
                return 0;
            else if (last_amount < 0)
            {
            }
            else if (dp.count(last_amount) > 0)
                return dp[last_amount];
            else
            {
                for (auto &&v : coins)
                    res = min(res, dfs(last_amount - v) + 1);
                dp[last_amount] = res;
            }
            return res;
        };

        long long ans = dfs(amount);

        return ans != numeric_limits<int>::max() ? ans : -1;
    }
};
// @lc code=end
