/*
 * @lc app=leetcode id=146 lang=typescript
 *
 * [146] LRU Cache
 */

class ListNode_ {
    key: number;
    val: number;
    prev: ListNode_;
    next: ListNode_;
}

// @lc code=start
class LRUCache {
    head: ListNode_;
    tail: ListNode_;
    nodes: ListNode_[];
    capacity: number;
    len: number;
    keyToNode: Map<number, ListNode_>;

    constructor(capacity: number) {
        this.capacity = capacity;
        this.nodes = new Array<ListNode_>(capacity);
        this.head = new ListNode_();
        this.tail = new ListNode_();
        this.head.next = this.tail;
        this.tail.prev = this.head;
        this.len = 0;
        this.keyToNode = new Map();
    }

    private removeNode(cur: ListNode_) {
        cur.prev.next = cur.next;
        cur.next.prev = cur.prev;
        this.keyToNode.delete(cur.key);
        this.len -= 1;
    }
    private addToHead(cur: ListNode_) {
        this.head.next.prev = cur;
        cur.next = this.head.next;
        cur.prev = this.head;
        this.head.next = cur;
        this.keyToNode.set(cur.key, cur);
        this.len += 1;
    }

    get(key: number): number {
        const find = this.keyToNode.get(key);

        if (find) {
            this.removeNode(find);
            this.addToHead(find);

            return find.val;
        }

        return -1;
    }

    put(key: number, value: number): void {
        const find = this.keyToNode.get(key);
        if (find) {
            find.val = value;
            this.removeNode(find);
            this.addToHead(find);
        } else {
            const newNode = new ListNode_();
            newNode.key = key;
            newNode.val = value;
            this.addToHead(newNode);

            if (this.len > this.capacity) {
                this.removeNode(this.tail.prev)
            }
        }
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */
// @lc code=end
