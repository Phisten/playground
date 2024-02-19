#
# @lc app=leetcode id=231 lang=python3
#
# [231] Power of Two
#


# @lc code=start
import unittest


class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        if n <= 0:
            return False
        if n == 1:
            return True
        while n > 1:
            if n % 2 == 1:
                return False
            n = n >> 1

        return True


# @lc code=end


# 測試類別，繼承自 unittest.TestCase
class TestFunction(unittest.TestCase):
    # 測試函數要以 "test_" 開頭
    def test_s231(self):
        s = Solution()

        self.assertEqual(s.isPowerOfTwo(n=6), False, "1")
        self.assertEqual(s.isPowerOfTwo(n=16), True, "2")
        self.assertEqual(s.isPowerOfTwo(n=3), False, "3")


if __name__ == "__main__":
    unittest.main()
