#
# @lc app=leetcode id=322 lang=python3
#
# [322] Coin Change
#


# @lc code=start
from typing import List


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp: dict = {}

        def dfs(last_amount: int) -> int:
            res = float("inf")
            if last_amount == 0:
                return 0
            elif last_amount < 0:
                None
            elif last_amount in dp:
                return dp[last_amount]
            else:
                for v in coins:
                    res = min(res, dfs(last_amount - v) + 1)
                dp[last_amount] = res

            return res

        ans = dfs(amount)
        return ans if ans != float("inf") else -1


# @lc code=end
