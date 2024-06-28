//  nextJsDynamicPath("/users/[id]", "123") -> "/users/123"。
export function nextJsDynamicPath<T extends string>(nextPath: T, param: string) {
  return nextPath.replace(/\[\S+\]/, param);
}

// getNextJsDynamicPathParamName("/users/[id]") -> "id"。
export function getNextJsDynamicPathParamName<T extends string>(nextPath: T) {
  return nextPath.match(/\[(.+?)\]$/)?.[1];
}
