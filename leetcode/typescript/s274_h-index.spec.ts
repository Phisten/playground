
describe('274. H-Index', () => {

  it('1', () => expect(hIndex([3, 0, 6, 1, 5])).toEqual(3))
  it('2', () => expect(hIndex([1, 3, 1])).toEqual(1))
})


// 0 1 3 5 6
function hIndex(citations: number[]): number {
  const len = citations.length;
  if (len === 1) {
    return citations[0];
  }

  citations.sort();
  for (let i = 1; i <= len; i++) {
    const ri = len - i;
    const v = citations[ri];
    if (v > ri) {
      return ri + 1
    }
  }

  return 0;
};
