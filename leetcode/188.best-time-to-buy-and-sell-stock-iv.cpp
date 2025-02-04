/*
 * @lc app=leetcode id=188 lang=cpp
 *
 * [188] Best Time to Buy and Sell Stock IV
 */

// @lc code=start
class Solution
{
public:
    int maxProfit(int k, vector<int> &prices)
    {
        int n = prices.size();
        if (n < 2)
            return 0;

        vector<vector<int>> profit(k + 1, vector<int>(n));

        int minBuy = prices[0];
        for (int j = 1; j < n; j++)
        {
            profit[1][j] = max(profit[1][j - 1], prices[j] - minBuy);
            minBuy = min(minBuy, prices[j]);
        }

        int ans = profit[1][n - 1];
        for (int i = 2; i <= k; i++)
        {
            vector<int> cash(n, -1000);

            int start = (i - 1) * 2;
            if (start >= n)
                break;

            for (int j = start; j < n; j++)
            {
                cash[j] = max(cash[j - 1], profit[i - 1][j - 1] - prices[j]);
            }

            profit[i][start] = profit[i - 1][start];

            for (int j = start + 1; j < n; j++)
            {
                profit[i][j] = max(max(profit[i - 1][j], profit[i][j - 1]), cash[j - 1] + prices[j]);
            }

            ans = max(ans, profit[i][n - 1]);
        }

        return ans;
    }
};
// @lc code=end
