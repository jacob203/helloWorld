//Definition of SegmentTreeNode:
class SegmentTreeNode {
public:
    int start, end;
    SegmentTreeNode *left, *right;
    SegmentTreeNode(int start, int end) {
        this->start = start, this->end = end;
        this->left = this->right = NULL;
    }
}

class Solution {
public:
    /**
     *@param start, end: Denote an segment / interval
     *@return: The root of Segment Tree
     */
    SegmentTreeNode * build(int start, int end) {
        if(start > end){
            return nullptr;
        }

        if(start == end){
            return new SegmentTreeNode(start, end);
        }

        auto mid = start + ((end - start)>>1);
        auto pNode = new SegmentTreeNode(start, end);
        pNode->left = build(start, mid);
        pNode->right = build(mid+1, end);
        return pNode;
    }
};