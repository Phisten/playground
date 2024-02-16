#
# @lc app=leetcode id=561 lang=python3
#
# [561] Array Partition
#

# @lc code=start
class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        nums = sorted(nums)

        ans = 0
        for i in range(len(nums)):
            if i % 2 == 0:
                ans += nums[i]

        return ans

    # @lc code=end

