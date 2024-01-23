#
# @lc app=leetcode id=907 lang=python3
#
# [907] Sum of Subarray Minimums
#

# @lc code=start
class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        left, right = 0, len(arr) - 1
        win = dict()

        for i in range(0, len(arr)):

            cur = arr[left]
            win[cur]++

            while ():







# @lc code=end

