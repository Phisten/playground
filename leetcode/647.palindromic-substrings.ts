/*
 * @lc app=leetcode id=647 lang=typescript
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
function countSubstrings(s: string): number {
  // Expand substring width from central indices and calculate the count of palindromes.
  const widthSearch = (l, r): number => {
    let res = 0;
    while (l >= 0 && r < s.length) {
      if (s[l] !== s[r])
        break;

      res++; l--; r++;
    }
    return res;
  }

  let ans = widthSearch(0, 0);
  // Enumerate all central indices of substrings.
  for (let i = 1; i < s.length; i++) {
    ans += widthSearch(i, i);
    ans += widthSearch(i - 1, i);
  }

  return ans;
};
// @lc code=end

function countSubstrings_bf(s: string): number {
  const length = s.length;

  // Check if the substring is a palindrome.
  const isPalindromic = (start, end): boolean => {
    for (let i = 0; i <= end - start; i++)
      if (s[start + i] !== s[end - i])
        return false

    return true
  }

  let ans = 0;
  // Enumerate all substrings of varying widths and their starting indices.
  for (let width = 1; width <= length; width++) {
    for (let i = 0; i <= length - width; i++) {
      // Increment the counter for palindrome substrings.
      if (isPalindromic(i, i + width - 1)) {
        ans++;
      }
    }
  }

  return ans;
};
