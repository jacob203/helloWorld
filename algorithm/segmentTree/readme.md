segment tree is used to fix the problems of range sum.  
for example we want to calculate the range sum, firstly we can calculate the sum from the range begin to the range end.  
but what if there are a lot of ranges you need to calculate? it can't be fixed in the same way, you need to use presum data structure, in the presum data structure,
for example:
```
  [1,2,     3,     4,      5,       6,       7,       8,       9]  
[0,1,3(1+2),6(3+3),10(6+4),15(10+5),21(15+6),28(21+7),36(28+8),45(36+9)]
```
preSum[i]=sum{[0:i)}  
so range[begin,end]=preSum[end+1] - preSum[begin]
all ranges sum can be fixed in O(1)

but what if the original data array can be modified?  
the preSum algorithm can not be used, as preSum is not accurate any more.  
segment tree is used to fix this type of issues. it calculates the range sum, when a data is changed, only the ranges are impacted, we can divide [0,3] as the following chart.
```
               [0,  3]
             /        \
      [0,  1]           [2, 3]
      /     \           /     \
   [0, 0]  [1, 1]     [2, 2]  [3, 3]
```
the time complexity of updating ranges when a data is changed is lgn, the time complexity of building segment tree is n。