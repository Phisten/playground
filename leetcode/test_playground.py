import timeit
import unittest


# 測試類別，繼承自 unittest.TestCase
class TestFunction(unittest.TestCase):
    def test_playground(self):

        # 測試函數要以 "test_" 開頭

        # 位元位移
        shift_time = timeit.timeit(
            "(a >> 7 & a >> 7) << 7", setup="a = 1000", number=1000000
        )

        # 模运算
        mod_time = timeit.timeit(
            "(a - (a % 10002340)) & (a - (a % 10002340))",
            setup="a = 1000",
            number=1000000,
        )

        print(f"位元位移时间：{shift_time}")
        print(f"模运算时间：{mod_time}")
        self.assertEqual(1, 1)


if __name__ == "__main__":
    unittest.main()
