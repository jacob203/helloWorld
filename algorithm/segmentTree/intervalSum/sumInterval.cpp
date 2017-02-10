/**
 * Definition of Interval:
 * classs Interval {
 *     int start, end;
 *     Interval(int start, int end) {
 *         this->start = start;
 *         this->end = end;
 *     }
 */
class CNode{
    public:
        int start;
        int end;
        long long sum;
        CNode* pLeft;
        CNode* pRight;
        CNode(int start, int end, int sum){
            this->start = start;
            this->end = end;
            this->sum = sum;
        }
};
class Solution {
public:
    /**
     *@param A, queries: Given an integer array and an query list
     *@return: The result list
     */
    vector<long long> intervalSum(vector<int> &A, vector<Interval> &queries) {
        vector<long long> res;
        if(A.empty() || queries.empty()){
            return res;
        }

        auto pRoot = BuildSegmentTree(A, 0, A.size()-1);
        for(auto& interval:queries){
            if(interval.start < 0
               || interval.end >= A.size()
               || interval.start > interval.end) {
                   res.push_back(0);
            } else{
                res.push_back(getSum(pRoot, interval.start, interval.end));
            }
        }

        return res;
    }

    CNode* BuildSegmentTree(vector<int> &A, int start, int end) {
        if(start == end){
            return new CNode(start, end, A[start]);
        }

        auto mid = start + ((end-start)>>1);
        auto pCur = new CNode(start, end, 0);
        pCur->pLeft = BuildSegmentTree(A, start, mid);
        pCur->pRight = BuildSegmentTree(A, mid+1, end);
        pCur->sum = pCur->pLeft->sum + pCur->pRight->sum;

        return pCur;
    }

    long long getSum(CNode* pCur, int start, int end) {
        if ((pCur->start == start) && (pCur->end == end)){
            return pCur->sum;
        }

        long long res = 0;
        auto mid = pCur->pLeft->end;
        if(end <= mid){
            return getSum(pCur->pLeft, start, end);
        } else if(start > mid) {
            return getSum(pCur->pRight, start, end);
        } else {
            return getSum(pCur->pLeft, start, mid) + getSum(pCur->pRight, mid+1, end);
        }
    }
};

=============================================================
/**
 * Definition of Interval:
 * classs Interval {
 *     int start, end;
 *     Interval(int start, int end) {
 *         this->start = start;
 *         this->end = end;
 *     }
 */
class Solution {
public:
    /**
     *@param A, queries: Given an integer array and an query list
     *@return: The result list
     */
    vector<long long> intervalSum(vector<int> &A, vector<Interval> &queries) {
        vector<long long> res;
        if(A.empty() || queries.empty()){
            return res;
        }

        vector<long long> preSum(A.size()+1);
        for(int i = 0; i < A.size(); i++){
            preSum[i+1] = preSum[i] + A[i];
        }

        for(auto& interval : queries){
            if(interval.start < 0
               || interval.end >= A.size()
               || interval.start > interval.end) {
                   res.push_back(0);
                   continue;
            }

            res.push_back(preSum[interval.end+1] - preSum[interval.start]);
        }

        return res;
    }
};