/*
 * @lc app=leetcode id=42 lang=typescript
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
function trap(height: number[]): number {
  let left = 0;
  let right = height.length - 1;

  const water: number[] = new Array(height.length).fill(0);

  while (left < right) {
    const minWall = Math.min(height[left], height[right])
    for (let i = left; i <= right; i++) {
      water[i] = Math.max(water[i], minWall - height[i])
    }

    if (height[left] < height[right]) {
      left++
    } else
      right--
  }

  let ans = 0;
  water.forEach(element => {
    ans += element;
  });

  return ans;
};
// @lc code=end
