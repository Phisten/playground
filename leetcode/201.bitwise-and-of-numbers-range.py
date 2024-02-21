#
# @lc app=leetcode id=201 lang=python3
#
# [201] Bitwise AND of Numbers Range
#


# @lc code=startclass Solution:
def rangeBitwiseAnd(self, left: int, right: int) -> int:
    i = 1
    l = left
    r = right

    while l > 0 and r > 0:
        l = l >> 1
        r = r >> 1
        i = i << 1

    if l == 0 and r == 0:
        diff = abs(left - right)
        i = 1
        while diff > 0:
            diff = diff >> 1
            i = i << 1

        if i > 0:
            return (left - (left % i)) & (right - (right % i))
        else:
            return left
    else:
        return 0


# @lc code=end
