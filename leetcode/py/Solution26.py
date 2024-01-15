# 26. Remove Duplicates from Sorted Array
# https://leetcode.com/problems/remove-duplicates-from-sorted-array/

from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 1
        for idx in range(1, len(nums)):
            if nums[idx] != nums[idx - 1]:
                nums[k] = nums[idx]
                k += 1
        return k


solution = Solution()
data_list = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
result = solution.removeDuplicates(data_list)
print(data_list)
print(result)
