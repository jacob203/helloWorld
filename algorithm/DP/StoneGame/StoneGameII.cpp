There is a stone game.At the beginning of the game the player picks n piles of stones in a circle.

The goal is to merge the stones in one pile observing the following rules:

At each step of the game,the player can merge two adjacent piles to a new pile.
The score is the number of stones in the new pile.
You are to determine the minimum of the total score.
Example
For [1, 4, 4, 1], in the best solution, the total score is 18:

1. Merge second and third piles => [2, 4, 4], score +2
2. Merge the first two piles => [6, 4]，score +6
3. Merge the last two piles => [10], score +10
Other two examples:
[1, 1, 1, 1] return 8
[4, 4, 5, 9] return 43

A[0],A[1:n-2],A[n-1]
F[i,j]=sum[i:j]+F[i,k]+F[k+1,j]
the final result is not equal to {A[0], A[1:n-2], A[n-1]}
we need to iterate i in the array, use each position as the start point, and find the min one.


class Solution {
public:
    /**
     * @param A an integer array
     * @return an integer
     */
    int stoneGame2(vector<int>& A) {
        // Write your code here
        //A[i,j]=min{A[i,k]+A[k+1,j] i<=k<=j}
        // add the array to the original A
        //A[0,n-1]=max{A[i, i+n-1], 0<=i<=n-1}
        if(A.size() < 2){
            return 0;
        }

        vector<int> newVector(2*A.size());
        for(int i = 0; i < A.size(); i++){
            newVector[i] = A[i];
            newVector[i+A.size()] = A[i];
        }
        vector<int> preSum(newVector.size()+1);
        for(int i = 1; i < preSum.size(); i++){
            preSum[i] += preSum[i-1] + newVector[i-1];
        }

        vector<vector<int>> score(newVector.size(), vector<int>(newVector.size()));
        for(int diff = 1; diff < score.size(); diff++){
            for(int i = 0; i+diff < score.size(); i++){
                auto minScore = numeric_limits<int>::max();
                for(int j = i; j < i+diff; j++){
                    minScore = min(minScore, score[i][j]+score[j+1][i+diff]+preSum[i+diff+1]-preSum[i]);
                }
                score[i][i+diff] = minScore;
            }
        }

        auto res = numeric_limits<int>::max();
        for(int i = 0; i < A.size(); i++){
            res = min(res, score[i][i+A.size()-1]);
        }

        return res;
    }

};