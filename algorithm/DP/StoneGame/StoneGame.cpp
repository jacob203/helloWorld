f[i, j] min value to merge from i to j
f[i,i] = A[i]
f[i,j] = min{f[i,k]+f[k+1,j] i<=k<j}
res = f[0, n-1]

[3,4,3]

each element should has two subelements, one is cur value, and the other is cur score

class Solution {
public:
    /**
     * @param A an integer array
     * @return an integer
     */
    int stoneGame(vector<int>& A) {
        if(A.size() < 2){
            return 0;
        }

        vector<vector<pair<int,int>>> minMergeValues(A.size(), vector<pair<int,int>>(A.size()));
        for(int i = 0; i < A.size(); i++){
            minMergeValues[i][i].first = A[i];
        }

        for(int diff = 1; diff < A.size(); diff++){
            for(int i = 0;  i+diff<A.size(); i++){
                int minValue = numeric_limits<int>::max();
                int curScore = 0;
                for(int k = i; k+1 <= i+diff; k++){
                    auto curValue = minMergeValues[i][k].first
                                    + minMergeValues[i][k].second
                                    + minMergeValues[k+1][i+diff].first
                                    + minMergeValues[k+1][i+diff].second;

                    if (minValue > curValue){
                        minValue = curValue;
                        curScore = minMergeValues[i][k].first + minMergeValues[k+1][i+diff].first;
                    }
                }
                minMergeValues[i][i+diff].first = curScore;
                minMergeValues[i][i+diff].second = minValue;
            }
        }

        return minMergeValues[0][A.size()-1].second;
    }
};

class Solution {
public:
    /**
     * @param A an integer array
     * @return an integer
     */
    int stoneGame(vector<int>& A) {
        if(A.size() < 2){
            return 0;
        }

        vector<vector<pair<int,int>>> minMergeValues(A.size(), vector<pair<int,int>>(A.size(), make_pair(-1,0)));
        for(int i = 0; i < A.size(); i++){
            minMergeValues[i][i].first = A[i];
        }

        return memorySearch(0, A.size()-1, minMergeValues).second;
    }

    pair<int,int> memorySearch(int begin, int end, vector<vector<pair<int,int>>>& minMergeValues){
        if(-1 != minMergeValues[begin][end].first){
            return minMergeValues[begin][end];
        }

        auto minValue = make_pair(0, numeric_limits<int>::max())
        for(int i = begin; i < end; i++){
            auto left = memorySearch(begin, i, minMergeValues);
            auto right = memorySearch(i+1, end, minMergeValues);
            if(left.first + left.second + right.first + right.second  < minValue.second){
                minValue.first = left.first + right.first;
                minValue.second = minValue.first + left.second + right.second;
            }
        }
        minMergeValues[begin][end] = minValue;

        return minValue;
    }
};


class WrongSolution {
public:
    /**
     * @param A an integer array
     * @return an integer
     */
    int stoneGame(vector<int>& A) {
        if(A.empty()){
            return 0;
        }
        if(A.size() == 1){
            return A[0];
        }

        vector<vector<int>> minMergeValues(A.size(), vector<int>(A.size()));
        for(int i = 0; i < A.size(); i++){
            minMergeValues[i][i] = A[i];
        }

        for(int diff = 1; diff < A.size(); diff++){
            for(int i = 0;  i+diff<A.size(); i++){
                int minValue = numeric_limits<int>::max();
                for(int k = i; k+1 <= i+diff; k++){
                    minValue = min(minValue, minMergeValues[i][k] + minMergeValues[k+1][i+diff]);
                }
                minMergeValues[i][i+diff] = 2*minValue;
            }
        }

        return minMergeValues[0][A.size()-1]/2;
    }
};

