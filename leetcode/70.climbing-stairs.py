#
# @lc app=leetcode id=70 lang=python3
#
# [70] Climbing Stairs
#

# @lc code=start
class Solution:
    def climbStairs(self, n: int) -> int:
        cur = 1
        prev = 1
        for i in range(2, n+1):
            tmp = cur + prev
            prev = cur
            cur = tmp
        return cur
        
# @lc code=end

