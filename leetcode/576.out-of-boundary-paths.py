#
# @lc app=leetcode id=576 lang=python3
#
# [576] Out of Boundary Paths
#

# @lc code=start
class Solution:
    mod = 1000000007

    def findPaths(self, m: int, n: int, maxMove: int, startRow: int, startColumn: int) -> int:

        dp = [[[None] * (maxMove+1) for _ in range(n)] for _ in range(m)]

        def dsf(remainingStep: int, x: int, y: int):
            if not (0 <= x < m and 0 <= y < n):
                return 1
            if remainingStep <= 0:
                return 0

            if dp[x][y][remainingStep] is not None:
                return dp[x][y][remainingStep]

            remainingStep -= 1

            res = dsf(remainingStep, x+1, y) % self.mod
            res += dsf(remainingStep, x-1, y) % self.mod
            res += dsf(remainingStep, x, y+1) % self.mod
            res += dsf(remainingStep, x, y-1) % self.mod

            dp[x][y][remainingStep] = res

            return res

        return dsf(maxMove, startRow, startColumn)

# @lc code=end

