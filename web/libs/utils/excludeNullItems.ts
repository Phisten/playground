// fix for "typescript": "~5.3.3"
export const excludeNullItems = <T>(arr: (T | null)[]) =>
  arr.filter((v: T | null) => v !== null) as Exclude<T, null>[];
