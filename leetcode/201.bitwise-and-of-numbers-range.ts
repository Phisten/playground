/*
 * @lc app=leetcode id=201 lang=typescript
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
function rangeBitwiseAnd(left: number, right: number): number {
  let i = 1;
  let diff = right - left;
  let j = 1

  while (i < diff && i > 0) {
    i = i << 1;
  }

  return (left - left % i) & right;
};
// @lc code=end
