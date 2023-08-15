# https://leetcode.com/problems/merge-sorted-array/

from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        curStart = 0
        curLen = m
        for n2cur in nums2:
            nums1[curLen] = n2cur
            for idx in range(curStart, curLen):
                if nums1[idx] > n2cur:
                    for j in range(curLen, idx, -1):
                        nums1[j] = nums1[j - 1]

                    nums1[idx] = n2cur
                    curStart = idx + 1
                    break
            curLen += 1

        print(nums1)


print(Solution().merge([4, 5, 6, 0, 0, 0], 3, [1, 2, 3], 3))
