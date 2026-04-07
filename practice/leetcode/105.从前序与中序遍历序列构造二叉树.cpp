/*
 * @lc app=leetcode.cn id=105 lang=cpp
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        for (int i = 0; i < inorder.size(); ++i) {
            inPos[inorder[i]] = i;
        }
        return build(preorder, 0, preorder.size() - 1, inorder, 0, inorder.size() - 1);
    }

private:
    unordered_map<int, int> inPos;

    TreeNode* build(const vector<int>& preorder, int preStart, int preEnd, const vector<int>& inorder, int inStart,
                    int inEnd) {
        if (preStart > preEnd) {
            return nullptr;
        }
        int rootVal = preorder[preStart];
        TreeNode* root = new TreeNode(rootVal);
        int idx = inPos[rootVal];
        int leftSz = idx - inStart;
        root->left = build(preorder, preStart + 1, preStart + leftSz, inorder, inStart, idx - 1);
        root->right = build(preorder, preStart + leftSz + 1, preEnd, inorder, idx + 1, inEnd);
        return root;
    }
};
// @lc code=end
