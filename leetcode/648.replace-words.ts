/*
 * @lc app=leetcode id=648 lang=typescript
 *
 * [648] Replace Words
 */

// @lc code=start
class MyTrie {
  ht: Map<string, MyTrie> = new Map<string, MyTrie>();
  isWord: boolean = false;

  constructor() {
    this.ht = new Map<string, MyTrie>();
    this.isWord = false;
  }

  insert(s: string) {
    const n = s.length;
    let curTrie: MyTrie = this;

    Array.from(s).forEach((e, i) => {
      let nextTrie = curTrie.ht.get(e);
      if (!nextTrie) {
        nextTrie = new MyTrie();
        curTrie.ht.set(e, nextTrie);
      }
      curTrie = nextTrie;

      if (i === n - 1) {
        curTrie.isWord = true;
      }
    });
  }

  matchShortWord(s: string): string | null {
    const n = s.length;
    let curTrie: MyTrie = this;

    let matchStr = "";
    for (let i = 0; i < n; i++) {
      const element = s[i];

      const nextTrie = curTrie.ht.get(element);
      if (!nextTrie) break;

      matchStr += element;
      curTrie = nextTrie;
      if (curTrie.isWord) return matchStr;
    }

    return null;
  }
}

function replaceWords(dictionary: string[], sentence: string): string {
  const trie: MyTrie = new MyTrie();

  dictionary.forEach((e) => {
    trie.insert(e);
  });

  const ansArr = sentence.split(" ").map((e, i) => {
    return trie.matchShortWord(e) ?? e;
  });

  return ansArr.join(" ");
}
// @lc code=end
