/*
 * @lc app=leetcode id=509 lang=cpp
 *
 * [509] Fibonacci Number
 */

// @lc code=start
#include <vector>
using namespace std;

class Solution
{
public:
    int fib(int n)
    {
        if (n <= 0)
            return n;

        vector<int> dp(n + 1, 0);

        dp[0] = 0;
        dp[1] = 1;
        for (int i = 2; i <= n; i++)
        {
            dp[i] = dp[i - 1] + dp[i - 2];
        }

        return dp[n];
    }
};
// @lc code=end
