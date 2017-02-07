/**
 * Definition of SegmentTreeNode:
 * class SegmentTreeNode {
 * public:
 *     int start, end, max;
 *     SegmentTreeNode *left, *right;
 *     SegmentTreeNode(int start, int end, int max) {
 *         this->start = start;
 *         this->end = end;
 *         this->max = max;
 *         this->left = this->right = NULL;
 *     }
 * }
 */
class Solution {
public:
    /**
     *@param root, index, value: The root of segment tree and 
     *@ change the node's value with [index, index] to the new given value
     *@return: void
     */
    void modify(SegmentTreeNode *root, int index, int value) {
        updateMax(root, index, value);
    }
    
    int updateMax(SegmentTreeNode *root, int index, int value){
        if(!root || index < root->start || index > root->end){
            return -1;//error
        }
        
        if(index == root->start && index == root->end){
            root->max = value;
            
            return value;
        } else {
            auto mid = root->left->end;
            if(index <= mid){
                auto leftMax = updateMax(root->left, index, value);
                root->max = max(leftMax, root->right->max);
            }else{
                auto rightMax = updateMax(root->right, index, value);
                root->max = max(root->left->max, rightMax);
            }
            
            return root->max;
        }        
    }
};