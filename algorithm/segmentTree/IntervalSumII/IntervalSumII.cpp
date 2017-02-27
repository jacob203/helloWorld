class CNode{
    public:
    int start;
    int end;
    long long sum;
    CNode* pLeft;
    CNode* pRight;
    CNode(int _start, int _end, int _sum):start(_start), end(_end), sum(_sum), pLeft(nullptr), pRight(nullptr){

    }
};
class Solution {
public:
    /* you may need to use some attributes here */


    /**
     * @param A: An integer vector
     */
    Solution(vector<int> A) {
        if(A.empty()){
            m_pRoot = nullptr;
        } else{
            m_pRoot = build(A, 0, A.size()-1);
        }
    }

    /**
     * @param start, end: Indices
     * @return: The sum from start to end
     */
    long long query(int start, int end) {
        return querySum(m_pRoot, start, end);
    }

    /**
     * @param index, value: modify A[index] to value.
     */
    void modify(int index, int value) {
        updateValue(m_pRoot, index, value);
    }

    private:
        CNode* m_pRoot;
    private:
        CNode* build(vector<int>&A, int start, int end){
            if(start == end){
                return new CNode(start, end, A[start]);
            }

            auto mid = start + ((end - start)>>1);
            auto pCur = new CNode(start, end, 0);
            pCur->pLeft = build(A, start, mid);
            pCur->pRight = build(A, mid+1, end);
            pCur->sum = pCur->pLeft->sum + pCur->pRight->sum;

            return pCur;
        }

        long long querySum(CNode* pRoot, int start, int end){
            if(!pRoot
               || start > end
               || start > pRoot->end
               || end < pRoot->start) {
                   return 0;
            }

            if(start <= pRoot->start && end >= pRoot->end) {
                return pRoot->sum;
            }

            if(pRoot->start == pRoot->end){
                return 0;
            }

            auto mid = pRoot->pLeft->end;
            if(end <= mid){
                return querySum(pRoot->pLeft, start, end);
            } else if(start > mid){
                return querySum(pRoot->pRight, start, end);
            } else {
                return querySum(pRoot->pLeft, start, mid) + querySum(pRoot->pRight, mid+1, end);
            }
        }

        int updateValue(CNode* pRoot, int index, int value){
            if(!pRoot
               || index < pRoot->start
               || index > pRoot->end){
                   return 0;
               }

            if(pRoot->start == pRoot->end && pRoot->start == index){
                auto diff = value - pRoot->sum;
                pRoot->sum += diff;
                return diff;
            }

            auto mid = pRoot->pLeft->end;
            int diff = 0;
            if(index <= mid){
                diff = updateValue(pRoot->pLeft, index, value);
            } else {
                diff = updateValue(pRoot->pRight, index, value);
            }
            pRoot->sum += diff;

            return diff;
        }

};
