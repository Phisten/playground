/*
 * @lc app=leetcode id=268 lang=typescript
 *
 * [268] Missing Number
 */

// @lc code=start
function missingNumber(nums: number[]): number {
  const ht = {}
  for (let i = 0; i < nums.length; i++)
    ht[nums[i]] = true;

  for (let i = 0; i <= nums.length; i++)
    if (ht[i] !== true) return i;
};
// @lc code=end
