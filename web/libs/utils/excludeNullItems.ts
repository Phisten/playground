// fix for typescript <5.5
export const excludeNullItems = <T>(arr: (T | null)[]) =>
  arr.filter((v: T | null) => v !== null) as Exclude<T, null>[];
