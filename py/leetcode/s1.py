# https://leetcode.com/problems/two-sum/

from typing import List


class Solution:
    # hashmap
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        sumDict = {target - item: [index, item] for index, item in enumerate(nums)}

        for index, item in enumerate(nums):
            if (
                item in sumDict
                and sumDict[item][1] is not None
                and sumDict[item][0] != index
            ):
                return [index, sumDict[item][0]]

    def twoSum_for(self, nums: List[int], target: int) -> List[int]:
        numsLen = len(nums)
        for i, item in enumerate(nums):
            curTarget = target - item
            for j in range(i + 1, numsLen):
                if curTarget == nums[j]:
                    return [i, j]


# print(Solution().twoSum([2, 7, 11, 15], 9))
# print(Solution().twoSum_for([2, 5, 5, 11], 10))
print(Solution().twoSum_for([2222222, 2222222], 4444444))
