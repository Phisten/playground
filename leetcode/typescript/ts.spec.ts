import Func from '../518.coin-change-ii'


describe('playground ts test', () => {

  it('1', () => expect(Func(5, [1, 2, 5])).toEqual(4))
  it('2', () => expect(Func(3, [2])).toEqual(0))
  it('3', () => expect(Func(0, [2])).toEqual(1))
  it('test', () => expect(Func(5, [1, 5])).toEqual(2))
})
