#
# @lc app=leetcode id=2402 lang=python3
#
# [2402] Meeting Rooms III
#

from typing import List
import unittest


# @lc code=start
class Solution:
    def mostBooked(self, n: int, meetings: List[List[int]]) -> int:
        room_using: List[List[int]] = [0] * n
        room_time: List[int] = [0] * n

        # equivalent to meetings.sort(key=lambda x: x[0])
        meetings.sort()
        print(meetings)

        time_shift = 0
        for start, end in meetings:
            # diff = start - cur_time

            min_ru_idx = 0
            min_ru = room_using[0]
            for i, alive_time in enumerate(room_using):
                if start >= alive_time:
                    room_using[i] = end
                    room_time[i] += 1
                    min_ru_idx = -1
                    break

                if alive_time < min_ru:
                    min_ru = alive_time
                    min_ru_idx = i

            if min_ru_idx > -1:
                time_shift = room_using[min_ru_idx] - start
                room_using[min_ru_idx] = time_shift + end
                room_time[min_ru_idx] += 1

        print(room_time)
        return room_time.index(max(room_time))


# @lc code=end


# 測試類別，繼承自 unittest.TestCase
class TestFunction(unittest.TestCase):
    # 測試函數要以 "test_" 開頭
    def test_s2402(self):
        s = Solution()

        self.assertEqual(
            s.mostBooked(
                n=2,
                meetings=[
                    [43, 44],
                    [34, 36],
                    [11, 47],
                    [1, 8],
                    [30, 33],
                    [45, 48],
                    [23, 41],
                    [29, 30],
                ],
            ),
            1,
        )
        self.assertEqual(
            s.mostBooked(n=2, meetings=[[4, 11], [1, 13], [8, 15], [9, 18], [0, 17]]), 1
        )
        # self.assertEqual(result, 8)  # 斷言結果是否等於預期值


if __name__ == "__main__":
    unittest.main()
