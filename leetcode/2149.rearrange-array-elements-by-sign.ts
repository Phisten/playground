/*
 * @lc app=leetcode id=2149 lang=typescript
 *
 * [2149] Rearrange Array Elements by Sign
 */

// @lc code=start
function rearrangeArray(nums: number[]): number[] {
  const ans: number[] = new Array(nums.length);

  let p = -2, n = -1;
  nums.forEach((v) => {
    ans[v > 0 ? p += 2 : n += 2] = v
  });

  return ans;
};
// @lc code=end

function rearrangeArray_first(nums: number[]): number[] {
  let positive = 0, negative = 1;
  let idx = 0;

  const ans: number[] = Array.from({ length: nums.length });

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] > 0) {
      ans[positive] = nums[i];
      positive += 2;
    } else {
      ans[negative] = nums[i];
      negative += 2;
    }
  }

  return ans;
};
