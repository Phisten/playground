/*
 * @lc app=leetcode id=647 lang=typescript
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
function countSubstrings(s: string): number {
  const length = s.length;

  const isPalindromic = (start, end): boolean => {
    for (let i = 0; i <= end - start; i++)
      if (s[start + i] !== s[end - i])
        return false

    return true
  }

  let ans = 0;
  for (let width = 1; width <= length; width++) {
    for (let i = 0; i <= length - width; i++) {
      if (isPalindromic(i, i + width - 1)) {
        ans++;
      }
    }
  }

  return ans;
};
// @lc code=end
