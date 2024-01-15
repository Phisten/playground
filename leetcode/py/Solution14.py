# Longest Common Prefix
# https://leetcode.com/problems/longest-common-prefix/

from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        minStr = min(strs, key=len)
        outStr = ""

        for index, char in enumerate(minStr):
            shared = True
            for peek in strs:
                if peek[index] != char:
                    shared = False

            if shared:
                outStr += char
            else:
                break

        return outStr


solution = Solution()

str_list = ["apple", "apricot", "appetite"]
result = solution.longestCommonPrefix(str_list)


print(result)
