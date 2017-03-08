* dp is used to avoid repeated calculation issues.
* max/min/how many issues might be dp issues
* find all solutions can't be dp issues
* the big issue with dp is how to present questions
* if you find how to present questions, use memory search to fix it, if it is hard for you to find the formulation

# stone game
```
There is a stone game.At the beginning of the game the player picks n piles of stones in a line.

The goal is to merge the stones in one pile observing the following rules:

At each step of the game,the player can merge two adjacent piles to a new pile.
The score is the number of stones in the new pile.
You are to determine the minimum of the total score.

For [4, 1, 1, 4], in the best solution, the total score is 18:

1. Merge second and third piles => [4, 2, 4], score +2
2. Merge the first two piles => [6, 4]，score +6
3. Merge the last two piles => [10], score +10
Other two examples:
[1, 1, 1, 1] return 8
[4, 4, 5, 9] return 43
```
use one dimension array:
F[i] means to array[i], min score, but you will find that F[i] can not be calculated from F[0~i-1]
so use two dimension array:
F[i][j] means from i to j min score, F[i][j] can be deduced by min{F[i][k]+F[k+1][j], i<=k<=j}
but soon you will find, it is not also enough, F[i][j] should has two elements, one is cur value, and the other is cur score
F[i][j].first means from i to j cur value, and F[i][j].second means the cur score
F[i][j] = min{F[i][k].first+F[i][k].second+F[k+1][j].first+F[k+1][j].second, i<=k<=j}
res = F[0][n-1].second

init:F[i][i] = (A[i],0)
if not using memory search,
for this kind of init, generally it expands from the mid, begin from [i][i], expand to the right.

another algorithm is :
F[i][j] = min{F[i][k] + F[k+1][j] + sum[i:j]}
actually it is how we find a good cut point.


stone game II
it can merge the piles by circular, it means it can merge the end pile and the begin pile.
generally this kind of issue can be divided into two types:
* double the original array, so the array becomes [original, original], then calc the F[0][newArraySize-1], all F[i][j] are calculated.
then the result is min{F[i][i+originalSize-1] 0<=i<originalSize}
* check if the issue can be divided into [begin, middle, end], middle means F[end, [0][n-2]], F[begin][end], but it isn't.




