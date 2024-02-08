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

function trap_s2(height: number[]): number {
  let left = 0, right = height.length - 1;
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

// time:O(n) space:O(1)
function trap(height: number[]): number {
  let left = 0, right = height.length - 1, ans = 0;
  let minWall = Math.min(height[left], height[right]);

  while (left < right) {
    minWall = Math.max(Math.min(height[left], height[right]), minWall)

    if (height[left] < height[right]) {
      if (minWall > height[left])
        ans += minWall - height[left];
      left++;
    } else {
      if (minWall > height[right])
        ans += minWall - height[right];
      right--;
    }
  }

  return ans;
};
// @lc code=end
