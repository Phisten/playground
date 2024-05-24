/*
 * @lc app=leetcode id=1 lang=typescript
 *
 * [1] Two Sum
 */

// @lc code=start
function twoSum(nums: number[], target: number): number[] {
  const ht = {};
  const ans = [];

  nums.forEach((v, i) => {
    if (ht[v] && ht[v] > 0) {
      ans.push(i);
      ans.push(ht[v] - 1);
      return;
    }
    ht[target - v] = i + 1;
  });

  ans.sort();
  return ans;
}
// @lc code=end
