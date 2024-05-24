/*
 * @lc app=leetcode id=146 lang=cpp
 *
 * [146] LRU Cache
 */

// @lc code=start
class DLink {
public:
    int key;
    int val;
    DLink *prev;
    DLink *next;
};
class LRUCache {
public:
    unordered_map<int, DLink*> ht;
    DLink head;
    DLink end;
    int n = 0;
    LRUCache(int capacity) {
        head.next = &end;
        head.key = -1;
        head.val = -1;
        end.prev = &head;
        end.key = -2;
        end.val = -2;
        n = capacity;
        ht.clear();
    }

    void removeFirst() {
        // cout << "remove: k:" << head.next->key << ", v:" << head.next->val << endl;
        ht.erase(head.next->key);
        head.next = head.next->next;
        head.next->prev = &head;
    }
    void removeByKey(int key) {
        ht[key]->next->prev = ht[key]->prev;
        ht[key]->prev->next = ht[key]->next;
        ht.erase(key);
    }
    void addToEnd(int key, int value) {
        DLink *curNode = new DLink(key, value, end.prev, &end);
        end.prev->next = curNode;
        end.prev = curNode;
        ht[key] = curNode;
    }
    
    int get(int key) {
        if (ht.count(key) > 0) {
            int value = ht[key]->val;
            removeByKey(key);
            addToEnd(key, value);
            return value;
        }
        return -1;
    }
    
    void put(int key, int value) {
        if (ht.count(key) > 0) {
            removeByKey(key);
            addToEnd(key, value);
        } else {
            if (ht.size() >= n) {
                removeFirst();
            }
            addToEnd(key, value);
        }

        
    }
};
/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
// @lc code=end

