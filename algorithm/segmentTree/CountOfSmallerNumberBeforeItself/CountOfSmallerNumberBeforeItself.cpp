/*
Give you an integer array (index from 0 to n-1, where n is the size of this array, data value from 0 to 10000) . For each element Ai in the array, count the number of element before this element Ai is smaller than it and return count number array.
For array [1,2,7,8,5], return [0,1,2,3,2]
specially deal with duplicated nums
[1,2,3,2,4,1] return [0, 1, 2, 1, 4, 0]
*/
class CNode {
    public:
    int low;
    int high;
    int count;
    CNode* pLeft;
    CNode* pRight;
    CNode(int _low, int _high, int _count):low(_low), high(_high), count(_count), pLeft(nullptr), pRight(nullptr){
    }
};
class Solution {
public:
   /**
     * @param A: An integer array
     * @return: Count the number of element before this element 'ai' is
     *          smaller than it and return count number array
     */
    vector<int> countOfSmallerNumberII(vector<int> &A) {
        vector<int> res;
        if(A.empty()){
            return res;
        }

        auto pRoot = buildSegmentTree(0, 10000);
        for(auto value:A){
            if(value < pRoot->low || value > pRoot->high){
                res.push_back(0);
            } else{
                res.push_back(lowerCount(pRoot, value));
                updateSegmentTree(pRoot, value);
            }
        }

        return res;
    }

    CNode* buildSegmentTree(int low, int high){
        auto pCur = new CNode(low, high, 0);
        if(low == high){
            return pCur;
        }

        auto mid = low + ((high-low)>>1);
        pCur->pLeft = buildSegmentTree(low, mid);
        pCur->pRight = buildSegmentTree(mid+1, high);

        return pCur;
    }

    int updateSegmentTree(CNode* pCur, int value) {
        if(!pCur || value < pCur->low || value > pCur->high){
            return 0;
        }

        if(pCur->low == pCur->high){
            if(pCur->low == value){
                pCur->count++;
                return 1;
            } else {
                return 0;
            }
        } else{
            auto mid = pCur->low + ((pCur->high-pCur->low)>>1);
            int diff = 0;
            if(value <= mid){
                diff = updateSegmentTree(pCur->pLeft, value);
            } else {
                diff = updateSegmentTree(pCur->pRight, value);
            }

            pCur->count += diff;
            return diff;
        }
    }

    int lowerCount(CNode* pCur, int value){
        if(!pCur
           || pCur->low > value
           || pCur->count == 0){
            return 0;
        }

        if(pCur->high < value){
            return pCur->count;
        }

        auto mid = pCur->low + ((pCur->high - pCur->low)>>1);
        if(value <= mid){
            return lowerCount(pCur->pLeft, value);
        } else{
            return lowerCount(pCur->pLeft, value) + lowerCount(pCur->pRight, value);
        }
    }

};