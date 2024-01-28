/*
 * @lc app=leetcode id=35 lang=typescript
 *
 * [35] Search Insert Position
 */

// @lc code=start
function searchInsert(nums: number[], target: number): number {

  let l = 0, r = nums.length - 1;

  while (l < r) {
    const mid = Math.floor((l + r) / 2);
    const cur = nums[mid]
    if (cur == target)
      return mid;
    else if (cur < target)
      l = mid + 1;
    else if (target < cur)
      r = mid - 1;
  }

  if (nums[l] >= target)
    return l;
  else if (nums[l] < target)
    return l + 1;
};
// @lc code=end
