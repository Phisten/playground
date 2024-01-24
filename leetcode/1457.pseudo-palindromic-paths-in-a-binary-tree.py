#
# @lc app=leetcode id=1457 lang=python3
#
# [1457] Pseudo-Palindromic Paths in a Binary Tree
#
from collections import defaultdict


# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pseudoPalindromicPaths(self, root: Optional[TreeNode]) -> int:
        ht = defaultdict(int)
        oddCount = 0
        ans = 0

        def dsf(cur, depth) -> bool:
            nonlocal oddCount, ans

            if not cur:
                return True

            ht[cur.val] += 1
            oddCount += 1 if ht[cur.val] % 2 == 1 else -1

            leftIsNull = dsf(cur.left, depth + 1)
            rightIsNull = dsf(cur.right, depth + 1)

            if leftIsNull and rightIsNull:
                if depth % 2 == oddCount or depth % 2 == oddCount:
                    ans += 1

            # rollback
            ht[cur.val] -= 1
            oddCount += 1 if ht[cur.val] % 2 == 1 else -1

            return False

        dsf(root, 1)

        return ans












# @lc code=end
