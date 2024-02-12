/*
 * @lc app=leetcode id=169 lang=typescript
 *
 * [169] Majority Element
 */

// @lc code=start
function majorityElement(nums: number[]): number {
  const ht: { [key in number]: number } = {}

  let ans = -1100000000;
  nums.forEach((v) => {
    if (ans > -1100000000) return;
    if (ht[v]) {
      ht[v]++;
    } else
      ht[v] = 1;

    if (ht[v] >= nums.length / 2)
      ans = v;
  });

  return ans;
};
// @lc code=end
