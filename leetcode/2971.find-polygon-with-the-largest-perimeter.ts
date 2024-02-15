/*
 * @lc app=leetcode id=2971 lang=typescript
 *
 * [2971] Find Polygon With the Largest Perimeter
 */

// @lc code=start
function largestPerimeter(nums: number[]): number {
  nums.sort((a, b) => a - b);

  let ans = -1, sum = nums[0];
  for (let i = 2; i < nums.length; i++) {
    sum += nums[i - 1];

    if (sum > 0 && nums[i] < sum && sum + nums[i] > ans)
      ans = sum + nums[i];
  }

  return ans;
};
// @lc code=end
