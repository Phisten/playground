#
# @lc app=leetcode id=1481 lang=python3
#
# [1481] Least Number of Unique Integers after K Removals
#


# @lc code=start
class Solution:
    def findLeastNumOfUniqueInts(self, arr: List[int], k: int) -> int:
        keys, ht = [], {}

        for item in arr:
            if item not in ht:
                ht[item] = 1
                keys.append(item)
            else:
                ht[item] += 1

        keys = sorted(keys, key=lambda v: ht[v])

        for i in range(len(arr)):
            k -= ht[keys[i]]
            if k == 0:
                return len(keys) - i - 1
            elif k < 0:
                return len(keys) - i

        return 0


# @lc code=end
