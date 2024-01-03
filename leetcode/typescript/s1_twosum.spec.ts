
describe('s1 two sum', () => {

  it('1', () => expect(twoSum([2, 7, 11, 15], 9)).toEqual([0, 1]))
  it('2', () => expect(twoSum([3, 2, 4], 6)).toEqual([1, 2]))
})

function twoSum(nums: number[], target: number): number[] {
  const hash = {};

  nums.forEach((v, i) => {
    hash[v] = i
  })

  for (let i = 0; i < nums.length; i++) {
    const v = nums[i];

    if (hash[target - v] !== i && hash[target - v] !== undefined) {
      return [i, hash[target - v]]
    }
  }
  return []
};
