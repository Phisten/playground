#
# @lc app=leetcode id=1642 lang=python3
#
# [1642] Furthest Building You Can Reach
#

# 順序尋訪找出 i-1<i 時兩樓層的差, 用來扣磚塊, 磚塊足夠墊高才能扣
# 磚塊用完拿梯子補
# 磚塊梯子都沒了就輸出最後idx

# *前述貪婪法無法判斷梯子的最佳使用時機(磚塊沒了才用會浪費磚塊)
# 過程diff需儲存
# 排序後扣除梯子數量的最大diff元素
# 剩餘的磚塊可排除即是可達到的idx

# 維護len(ladders)長度的list, 儲存用梯子的diff
# list塞滿以後, 新值進入前先排序, 拿最小值跟新值比, 小的去扣磚塊
# 磚塊不夠扣時 輸出為idx - 1
# opt:若梯子大於剩餘樓層數, return lastIdx

from heapq import heappush, heappop
from typing import List


# @lc code=start
class Solution:
    def furthestBuilding(self, heights: List[int], bricks: int, ladders: int) -> int:
        n = len(heights)

        ladder_options: List[int] = []

        for i in range(1, n):
            v = heights[i]
            lastLadders = ladders - len(ladder_options)
            if lastLadders >= n - i:
                return n - 1

            diff = v - heights[i - 1]
            if diff > 0:
                heappush(ladder_options, diff)
                if lastLadders == 0:
                    min_value = heappop(ladder_options)
                    bricks -= min_value
                    if bricks < 0:
                        return i - 1

        return n - 1

# @lc code=end
