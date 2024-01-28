#
# @lc app=leetcode id=1074 lang=python3
#
# [1074] Number of Submatrices That Sum to Target
#
# @lc code=start
class Solution:
    def numSubmatrixSumTarget(self, matrix: List[List[int]], target: int) -> int:
        width = len(matrix[0])
        height = len(matrix)

        sumMatrix = [[0 for _ in range(width + 1)] for _ in range(height + 1)]

        for i in range(height):
            for j in range(width):
                sumMatrix[i + 1][j + 1] = matrix[i][j] + sumMatrix[i][j + 1] + sumMatrix[i + 1][j] - \
                                  sumMatrix[i][j]

        curAns = 0
        for y2 in range(height):
            for x2 in range(width):
                for y1 in range(y2 + 1):
                    for x1 in range(x2 + 1):
                        curSum = sumMatrix[y2 + 1][x2 + 1] - sumMatrix[y1][x2 + 1] - sumMatrix[y2 + 1][x1] + \
                                 sumMatrix[y1][x1]
                        if curSum == target:
                            curAns += 1

        return curAns

# @lc code=end
