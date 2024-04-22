class Solution {
public:
    int numberOfSpecialChars(string word) {
        unordered_map<char, int> ht;
        char shift = 'a' - 'A';
        
        int n = word.length();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            char c = word[i];
            if (ht[c] == 0) {
                if (ht[c-shift] > 0) ans++; 
                if (ht[c+shift] > 0) ans++; 
                ht[c]++;
            }
        }
        
        return ans;
    }
};