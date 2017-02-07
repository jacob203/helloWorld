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
the segment tree is used to calculate the range result when the array can be changed.

# usage
### find the range max number
the time complexity is lgn
For array [1, 4, 2, 3], the corresponding Segment Tree is:
```

                  [0, 3, max=4]
                 /             \
          [0,1,max=4]        [2,3,max=3]
          /         \        /         \
   [0,0,max=1] [1,1,max=4] [2,2,max=2], [3,3,max=3]
query(root, 1, 1), return 4

query(root, 1, 2), return 4

query(root, 2, 3), return 3

query(root, 0, 2), return 4
```
### update the range max number
the time complexity is lgn
```
For segment tree:

                      [1, 4, max=3]
                    /                \
        [1, 2, max=2]                [3, 4, max=3]
       /              \             /             \
[1, 1, max=2], [2, 2, max=1], [3, 3, max=0], [4, 4, max=3]
if call modify(root, 2, 4), we can get:

                      [1, 4, max=4]
                    /                \
        [1, 2, max=4]                [3, 4, max=3]
       /              \             /             \
[1, 1, max=2], [2, 2, max=4], [3, 3, max=0], [4, 4, max=3]
or call modify(root, 4, 0), we can get:

                      [1, 4, max=2]
                    /                \
        [1, 2, max=2]                [3, 4, max=0]
       /              \             /             \
[1, 1, max=2], [2, 2, max=1], [3, 3, max=0], [4, 4, max=0]
```