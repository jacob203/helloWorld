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
### calc the interval sum
if there are no changes in the array, using the method presum to calc the interval sum is the fastest way.
but if the array datas can be changed, can't use the method preSum, segment tree is the better method
### Count of Smaller Number before itself
Give you an integer array (index from 0 to n-1, where n is the size of this array, data value from 0 to 10000) . For each element Ai in the array, count the number of element before this element Ai is smaller than it and return count number array.
```
For array [1,2,7,8,5], return [0,1,2,3,2]
[1,2,3,2,4,1] return [0, 1, 2, 1, 4, 0]
watch out duplicated nums
```
the question says the array data are in the range 0~10000, when it appears, it means that you can use an array whose size is 10000 to mark if the data exists or not.
and it wants the value counts smaller than the value, in another words, it want the value count in the range between 0 and value except value itself.
when it comes to range result, you't better think of segment tree to fix it.
so build the segment tree in 0~10000, whose count is the count of values between low and high, when querying a num, it means finding the 0~num counts, so it is translated into a segment question.
