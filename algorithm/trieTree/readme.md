# build trie tricks
generally the trie node is like
```
 class TrieNode {
 public:
     TrieNode() {}
     map<char, TrieNode*> children;
     vector<int> top10;
 };
```
remember that the current node is the result top10 of previous character