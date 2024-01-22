#
# @lc app=leetcode id=645 lang=python3
#
# [645] Set Mismatch
#
from typing import List

# @lc code=start
class Solution:
    def findErrorNums(self, nums: list[int]) -> list[int]:
        l, total = len(nums), 0

        for i in range(l):
            val = abs(nums[i])
            j = val - 1
            if nums[j] < 0:
                errVal = val
            total += val
            nums[j] = -nums[j]

        diff = total - (1 + l) * l // 2
        missVal = errVal - diff

        return [errVal, missVal]

# @lc code=end

