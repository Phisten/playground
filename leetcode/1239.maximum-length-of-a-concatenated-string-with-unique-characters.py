#
# @lc app=leetcode id=1239 lang=python3
#
# [1239] Maximum Length of a Concatenated String with Unique Characters
#
from typing import List

import numpy as np


# @lc code=start
# dsf+dp?

class Solution:
    def maxLength(self, arr: List[str]) -> int:
        tmpSet = [set()]
        for i in arr:
            if len(set(i)) < len(i):
                continue

            i = set(i)

            for j in tmpSet:
                if i & j:
                    continue

                tmpSet.append(i | j)

        m = 0
        for i in tmpSet:
            if m < len(i):
                m = len(i)
        return m

# @lc code=end
