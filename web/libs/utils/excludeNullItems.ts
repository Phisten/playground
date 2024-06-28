// fix for typescript <5.5
export const excludeNullItems = <T>(arr: T[]) =>
  arr.filter((v: T) => v !== null) as Exclude<T, null>[];
export const excludeUndefinedItems = <T>(arr: T[]) =>
  arr.filter((v: T) => v !== undefined) as Exclude<T, undefined>[];
export const excludeNullUndefinedItems = <T>(arr: T[]) =>
  arr.filter((v: T) => v !== null && v !== undefined) as Exclude<
    T,
    null | undefined
  >[];
