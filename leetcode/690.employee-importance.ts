/*
 * @lc app=leetcode id=690 lang=typescript
 *
 * [690] Employee Importance
 */

// @lc code=start
/**
 * Definition for Employee.
 * class Employee {
 *     id: number
 *     importance: number
 *     subordinates: number[]
 *     constructor(id: number, importance: number, subordinates: number[]) {
 *         this.id = (id === undefined) ? 0 : id;
 *         this.importance = (importance === undefined) ? 0 : importance;
 *         this.subordinates = (subordinates === undefined) ? [] : subordinates;
 *     }
 * }
 */
/**
 * Definition for Employee.
 * class Employee {
 *     id: number
 *     importance: number
 *     subordinates: number[]
 *     constructor(id: number, importance: number, subordinates: number[]) {
 *         this.id = (id === undefined) ? 0 : id;
 *         this.importance = (importance === undefined) ? 0 : importance;
 *         this.subordinates = (subordinates === undefined) ? [] : subordinates;
 *     }
 * }
 */

function getImportance(employees: Employee[], id: number): number {
  const ht = {};

  employees.forEach((v) => {
      ht[v.id] = v;
  });

  let ans = 0;

  const dfs = (cur: number) => {
      const curNode = ht[cur];
      ans += curNode.importance;

      curNode.subordinates.forEach((v) => {
          dfs(v);
      });
  };

  dfs(id);

  return ans;
};
// @lc code=end

