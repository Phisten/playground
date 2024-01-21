#
# @lc app=leetcode id=509 lang=python3
#
# [509] Fibonacci Number
#

# @lc code=start
class Solution:
    def fib(self, n: int) -> int:
        dp = [-1] * n+1

        def fibSelf(n: int) -> int:
            if n <= 1:
                return n

            if dp[n] != -1:
                return dp[n]

            return fibSelf(n - 1) + fibSelf(n - 2)

        return fibSelf(n)
# @lc code=end

