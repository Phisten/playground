#
# @lc app=leetcode id=33 lang=python3
#
# [33] Search in Rotated Sorted Array
#
from typing import List


# @lc code=start
def bs(nums: List[int], target: int, left: int, right: int) -> int:
    while left <= right:
        mid = (left + right) // 2
        cur = nums[mid]
        if cur == target:
            return mid
        elif cur > target:
            right = mid - 1
        elif cur < target:
            left = mid + 1

    return -1


def bsRotated(nums: List[int], left: int, right: int) -> int:
    endNum = nums[right]

    while left <= right:
        mid = (left + right) // 2
        cur = nums[mid]

        if mid == len(nums)-1:
            break
        elif cur > endNum and cur > nums[mid + 1]:
            return mid
        elif cur < endNum:
            right = mid - 1
        elif cur < nums[mid + 1]:
            left = mid + 1

    return -1


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        rotated_idx: int = bsRotated(nums, 0, len(nums) - 1)

        if rotated_idx > -1:
            if target >= nums[0]:
                return bs(nums, target, 0, rotated_idx)
            else:
                return bs(nums, target, rotated_idx + 1, len(nums) - 1)
        else:
            return bs(nums, target, 0, len(nums) - 1)


# @lc code=end
