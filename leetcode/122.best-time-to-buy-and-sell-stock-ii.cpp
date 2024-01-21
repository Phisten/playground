/*
 * @lc app=leetcode id=122 lang=cpp
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
#include <vector>
using namespace std;

class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int len = prices.size();

        int money = 0;

        for (int i = 1; i < len; i++)
            // try sell
            if (prices[i] - prices[i - 1] > 0)
                money += prices[i] - prices[i - 1];

        return money;
    }
};
// @lc code=end
