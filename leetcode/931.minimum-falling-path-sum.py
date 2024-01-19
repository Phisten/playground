#
# @lc app=leetcode id=931 lang=python3
#
# [931] Minimum Falling Path Sum
#
from typing import List


# @lc code=start
class Solution:
    def minFallingPathSum(self, matrix: List[List[int]]) -> int:
        m_len = len(matrix)
        for i in range(1, m_len):
            for j in range(m_len):
                minPath = matrix[i-1][j]
                if j-1 >= 0:
                    minPath = min(minPath, matrix[i - 1][j-1])
                if j+1 < m_len:
                    minPath = min(minPath, matrix[i - 1][j+1])
                matrix[i][j] += minPath

        min_val = float('inf')
        for i in range(m_len):
            min_val = min(min_val, matrix[m_len-1][i])

        return min_val
# @lc code=end

