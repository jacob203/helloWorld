/**
 * Definition of TrieNode:
 * class TrieNode {
 * public:
 *     TrieNode() {}
 *     map<char, TrieNode*> children;
 *     vector<int> top10;
 * };
 */
class TrieService {
private:
    TrieNode* root;

public:
    TrieService() {
        root = new TrieNode();
    }

    TrieNode* getRoot() {
        // Return root of trie root, and
        // lintcode will print the tree struct.
        return root;
    }

    void insert(string& word, int frequency) {
        auto pCur = root;
        for(int i = 0; i < word.size(); i++){
            insert(pCur, word[i], frequency);
            pCur = pCur->children[word[i]];
        }
    }

    void insert(TrieNode* pCur, char c, int frequency){
        if (!pCur->children[c]){
            pCur->children[c] = new TrieNode;
        }
        insertNum(pCur->children[c], frequency);
    }

    void insertNum(vector<int>& top10, int frequency){
        if(top10 >= 10 && top10.back() >= frequency){
            return;
        }

        if(top10.size() < 10) {
            top10.push_back(frequency);
        } else if(top10.back() < frequency){
            top10[9] = frequency;
        }
        int pos = top10.size() - 1;
        for(; pos > 0; pos--){
            if(top10[pos-1] < frequency){
                top10[pos] = top10[pos-1];
            } else {
                break;
            }
        }
        top10[pos] = frequency;
    }
};

