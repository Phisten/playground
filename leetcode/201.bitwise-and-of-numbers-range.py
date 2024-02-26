#
# @lc app=leetcode id=201 lang=python3
#
# [201] Bitwise AND of Numbers Range
#


# @lc code=start
class Solution:
    def rangeBitwiseAnd(self, left: int, right: int) -> int:
        i, diff = 1, right - left

        while i < diff:
            i = i << 1

        return (left - left % i) & right


# @lc code=end


class Solution_offset:
    def rangeBitwiseAnd(self, left: int, right: int) -> int:
        i, diff = 0, right - left

        while diff > 0:
            diff = diff >> 1
            i += 1

        return (left >> i << i) & right
