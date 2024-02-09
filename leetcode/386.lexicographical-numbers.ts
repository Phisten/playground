/*
 * @lc app=leetcode id=386 lang=typescript
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
function largestDivisibleSubset(nums: number[]): number[] {
  const length = nums.length
  nums.sort((a, b) => a - b);

  // store the length of the largest divisible subset at each index
  const dp: number[] = Array.from<number>({ length }).fill(1);
  // store the index of the previous element in the largest divisible subset
  const linked: number[] = Array.from<number>({ length });

  // Variables to track the maximum subset length and the index of its head element
  let maxSubset = 1;
  let headIdx = 0;

  // Dynamic Programming: Iterate through the array to find the largest divisible subset
  for (let i = 1; i < length; i++) {
    for (let j = 0; j < i; j++) {
      if (nums[i] % nums[j] == 0) {
        // If nums[i] is divisible by nums[j], update the subset length and linked index
        if (dp[j] + 1 > dp[i]) {
          dp[i] = dp[j] + 1;
          linked[i] = j;

          // Update the maximum subset length and linkedList head index
          if (dp[i] > maxSubset) {
            headIdx = i;
            maxSubset = dp[i];
          }
        }
      }
    }
  }

  // Reconstruct the largest divisible subset based on linked indices
  const ans: number[] = [];
  while (nums[headIdx]) {
    ans.push(nums[headIdx])
    headIdx = linked[headIdx];
  }

  return ans;
};
// @lc code=end
