class Solution {
public:
    int numberOfSpecialChars(string word) {
        unordered_map<char, int> ht;
        unordered_map<char, int> ht_used;
        unordered_map<char, int> ht_not_ans;
        char shift = 'a' - 'A';
        
        int n = word.length();
        for (int i = 0; i < n; i++) {
            char c = word[i];
            ht[c]++;
            
            // cur lowcase, find uppercase
            if (ht[c-shift] > 0) {
                ht_not_ans[c]++;
                ht_not_ans[c-shift]++;
            }
        }
        
        int ans = 0;
        for (int i = 0; i < n; i++) {
            char c = word[i];
            
            if (ht[c+shift] > 0 && ht_not_ans[c] == 0 && ht_used[c] == 0) {
                ht_used[c] = 1;
                ans++;
            }
            
        }
        
        return ans;
    }
};