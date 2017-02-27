/**
 * Definition of SegmentTreeNode:
 * class SegmentTreeNode {
 * public:
 *     int start, end, count;
 *     SegmentTreeNode *left, *right;
 *     SegmentTreeNode(int start, int end, int count) {
 *         this->start = start;
 *         this->end = end;
 *         this->count = count;
 *         this->left = this->right = NULL;
 *     }
 * }
 */
class Solution {
public:
    /**
     *@param root, start, end: The root of segment tree and
     *                         an segment / interval
     *@return: The count number in the interval [start, end]
     */
    int query(SegmentTreeNode *root, int start, int end) {
        //deal with connor cases
        if (!root
            || start > end
            || root->start > end
            || root->end < start){
            return 0;
        }

        if(root->start >= start && root->end <= end){
            return root->count;
        }

        if(root->start == root->end){// if start == end, it means it is a leaf. make sure it will not crash by calling root->left->end
            return 0;
        }

        auto mid = root->left->end;
        if(end <= mid){
            return query(root->left, start, end);
        } else if(start > mid){
            return query(root->right, start, end);
        } else {
            return query(root->left, start, mid) + query(root->right, mid+1, end);
        }
    }
};
