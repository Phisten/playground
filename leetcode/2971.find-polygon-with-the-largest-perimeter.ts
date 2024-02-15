/*
 * @lc app=leetcode id=2971 lang=typescript
 *
 * [2971] Find Polygon With the Largest Perimeter
 */

// @lc code=start
function largestPerimeter(nums: number[]): number {
  let ans = -1;

  nums.sort((a, b) => a - b);

  nums.forEach((v, i) => {
    let sum = 0;
    for (let j = 0; j < i; j++)
      sum += nums[j];

    if (sum > 0 && v < sum && sum + v > ans)
      ans = sum + v;
  });

  return ans;
};
// @lc code=end
