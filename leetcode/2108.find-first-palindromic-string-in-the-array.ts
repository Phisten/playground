/*
 * @lc app=leetcode id=2108 lang=typescript
 *
 * [2108] Find First Palindromic String in the Array
 */

// @lc code=start
function firstPalindrome(words: string[]): string {
  let ans = ""

  const isPalindrome = (str: string) => {
    for (let i = 0; i < str.length / 2; i++)
      if (str[i] !== str[str.length - i - 1])
        return false

    ans = str;
    return true;
  }

  words.some((v) => isPalindrome(v))

  return ans;
};
// @lc code=end
