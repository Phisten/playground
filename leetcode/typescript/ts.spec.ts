import Func from '../518.coin-change-ii'


describe('playground ts test', () => {

  it('1', () => expect(Func(5, [1, 2, 5])).toEqual(4))
  it('2', () => expect(Func(3, [2])).toEqual(0))
  it('3', () => expect(Func(0, [2])).toEqual(1))
  it('test', () => expect(Func(5, [1, 5])).toEqual(2))
})



describe('playground ts test isPalindrome', () => {

  const words = ["abc", "aaa", "aa", "aea"];
  const isPalindrome = (idx: number) => {
    const len = words[idx].length;
    for (let i = 0; i < len / 2; i++) {
      if (words[idx][i] !== words[idx][len - i - 1])
        return false
    }

    return true;
  }

  it('1', () => expect(isPalindrome(0)).toEqual(false))
  it('2', () => expect(isPalindrome(1)).toEqual(true))
  it('3', () => expect(isPalindrome(2)).toEqual(true))
  it('4', () => expect(isPalindrome(3)).toEqual(true))
})
