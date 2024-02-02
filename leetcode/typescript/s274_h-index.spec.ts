
describe('274. H-Index', () => {

  it('1', () => expect(hIndex([3, 0, 6, 1, 5])).toEqual(3))
  // 0 1 3 5 6
  it('2', () => expect(hIndex([1, 3, 1])).toEqual(1))
  it('3', () => expect(hIndex([100])).toEqual(1))
  it('4', () => expect(hIndex([0])).toEqual(0))
  // 3 3 4 4 6 6 8 10
  // 8 7 6 5 4 3 2 1
  it('5', () => expect(hIndex([6, 6, 4, 8, 4, 3, 3, 10])).toEqual(4))
})


function hIndex(citations: number[]): number {
  const len = citations.length;

  citations.sort((a, b) => a - b);

  for (let i = 0; i < len; i++) {
    const target = len - i;
    const v = citations[i];
    if (v >= target) {
      return target;
    }
  }

  return 0;
};
