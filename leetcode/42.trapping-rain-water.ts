/*
 * @lc app=leetcode id=42 lang=typescript
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
function trap_s1(height: number[]): number {
  let left = 0;
  let right = height.length - 1;

  const water: number[] = new Array(height.length).fill(0);

  while (left < right) {
    const minWall = Math.min(height[left], height[right])
    for (let i = left + 1; i < right; i++) {
      water[i] = minWall;
    }

    while (left < right && (height[left] <= minWall || height[right] <= minWall)) {
      if (height[left] < height[right])
        left++
      else
        right--
    }
  }

  return water.reduce((pre, cur, i) => pre + Math.max(0, cur - height[i]));
};

function trap(height: number[]): number {
  let left = 0;
  let right = height.length - 1;

  const water: number[] = new Array(height.length).fill(0);

  while (left < right) {
    const minWall = Math.min(height[left], height[right]);

    if (height[left] == minWall) {
      do {
        water[left] = minWall;
        left++;
      } while (left < right && height[left] <= minWall);
    }
    else {
      do {
        water[right] = minWall;
        right--;
      } while (left < right && height[right] <= minWall);
    }
  }

  return water.reduce((sum, water, i) => sum + Math.max(0, water - height[i]), 0);
};
// @lc code=end
